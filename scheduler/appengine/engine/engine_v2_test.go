// Copyright 2017 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package engine

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"

	"go.chromium.org/gae/filter/featureBreaker"
	"go.chromium.org/gae/service/datastore"
	"go.chromium.org/gae/service/taskqueue"

	"go.chromium.org/luci/appengine/tq/tqtesting"
	"go.chromium.org/luci/common/auth/identity"
	"go.chromium.org/luci/common/clock"
	"go.chromium.org/luci/common/clock/testclock"
	"go.chromium.org/luci/common/proto/google"
	"go.chromium.org/luci/common/retry/transient"
	"go.chromium.org/luci/server/auth"

	"go.chromium.org/luci/scheduler/appengine/catalog"
	"go.chromium.org/luci/scheduler/appengine/internal"
	"go.chromium.org/luci/scheduler/appengine/task"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEnqueueInvocations(t *testing.T) {
	t.Parallel()

	Convey("Works", t, func() {
		c := newTestContext(epoch)
		e, _ := newTestEngine()

		tq := tqtesting.GetTestable(c, e.cfg.Dispatcher)
		tq.CreateQueues()

		job := Job{JobID: "project/job-v2"}
		So(datastore.Put(c, &job), ShouldBeNil)

		var invs []*Invocation
		err := runTxn(c, func(c context.Context) error {
			var err error
			invs, err = e.enqueueInvocations(c, &job, []InvocationRequest{
				{TriggeredBy: "user:a@example.com"},
				{TriggeredBy: "user:b@example.com"},
			})
			datastore.Put(c, &job)
			return err
		})
		So(err, ShouldBeNil)

		// The order of new invocations is undefined (including IDs assigned to
		// them), so convert them to map and clear IDs.
		invsByTrigger := map[identity.Identity]Invocation{}
		invIDs := map[int64]bool{}
		for _, inv := range invs {
			invIDs[inv.ID] = true
			cpy := *inv
			So(cpy.ID, ShouldResemble, cpy.InvocationNonce)
			cpy.ID = 0
			cpy.InvocationNonce = 0
			invsByTrigger[inv.TriggeredBy] = cpy
		}
		So(invsByTrigger, ShouldResemble, map[identity.Identity]Invocation{
			"user:a@example.com": {
				JobID:       "project/job-v2",
				Started:     epoch,
				TriggeredBy: "user:a@example.com",
				Status:      task.StatusStarting,
				DebugLog: "[22:42:00.000] New invocation initialized\n" +
					"[22:42:00.000] Manually triggered by user:a@example.com\n",
			},
			"user:b@example.com": {
				JobID:       "project/job-v2",
				Started:     epoch,
				TriggeredBy: "user:b@example.com",
				Status:      task.StatusStarting,
				DebugLog: "[22:42:00.000] New invocation initialized\n" +
					"[22:42:00.000] Manually triggered by user:b@example.com\n",
			},
		})

		// Both invocations are in ActiveInvocations list of the job.
		So(len(job.ActiveInvocations), ShouldEqual, 2)
		for _, invID := range job.ActiveInvocations {
			So(invIDs[invID], ShouldBeTrue)
		}

		// And we've emitted the launch task.
		tasks := tq.GetScheduledTasks()
		So(tasks[0].Payload, ShouldHaveSameTypeAs, &internal.LaunchInvocationsBatchTask{})
		batch := tasks[0].Payload.(*internal.LaunchInvocationsBatchTask)
		So(len(batch.Tasks), ShouldEqual, 2)
		for _, subtask := range batch.Tasks {
			So(subtask.JobId, ShouldEqual, "project/job-v2")
			So(invIDs[subtask.InvId], ShouldBeTrue)
		}
	})
}

func TestTriageTaskDedup(t *testing.T) {
	t.Parallel()

	Convey("with fake env", t, func() {
		c := newTestContext(epoch)
		e, _ := newTestEngine()

		tq := tqtesting.GetTestable(c, e.cfg.Dispatcher)
		tq.CreateQueues()

		Convey("single task", func() {
			So(e.kickTriageJobStateTask(c, "fake/job"), ShouldBeNil)

			tasks := tq.GetScheduledTasks()
			So(len(tasks), ShouldEqual, 1)
			So(tasks[0].Task.ETA.Equal(epoch.Add(2*time.Second)), ShouldBeTrue)
			So(tasks[0].Payload, ShouldResemble, &internal.TriageJobStateTask{JobId: "fake/job"})
		})

		Convey("a bunch of tasks, deduplicated by hitting memcache", func() {
			So(e.kickTriageJobStateTask(c, "fake/job"), ShouldBeNil)

			clock.Get(c).(testclock.TestClock).Add(time.Second)
			So(e.kickTriageJobStateTask(c, "fake/job"), ShouldBeNil)

			clock.Get(c).(testclock.TestClock).Add(900 * time.Millisecond)
			So(e.kickTriageJobStateTask(c, "fake/job"), ShouldBeNil)

			tasks := tq.GetScheduledTasks()
			So(len(tasks), ShouldEqual, 1)
			So(tasks[0].Task.ETA.Equal(epoch.Add(2*time.Second)), ShouldBeTrue)
			So(tasks[0].Payload, ShouldResemble, &internal.TriageJobStateTask{JobId: "fake/job"})
		})

		Convey("a bunch of tasks, deduplicated by hitting task queue", func() {
			c, fb := featureBreaker.FilterMC(c, fmt.Errorf("omg, memcache error"))
			fb.BreakFeatures(nil, "GetMulti", "SetMulti")

			So(e.kickTriageJobStateTask(c, "fake/job"), ShouldBeNil)

			clock.Get(c).(testclock.TestClock).Add(time.Second)
			So(e.kickTriageJobStateTask(c, "fake/job"), ShouldBeNil)

			clock.Get(c).(testclock.TestClock).Add(900 * time.Millisecond)
			So(e.kickTriageJobStateTask(c, "fake/job"), ShouldBeNil)

			tasks := tq.GetScheduledTasks()
			So(len(tasks), ShouldEqual, 1)
			So(tasks[0].Task.ETA.Equal(epoch.Add(2*time.Second)), ShouldBeTrue)
			So(tasks[0].Payload, ShouldResemble, &internal.TriageJobStateTask{JobId: "fake/job"})
		})
	})
}

func TestLaunchInvocationTask(t *testing.T) {
	t.Parallel()

	Convey("with fake env", t, func() {
		c := newTestContext(epoch)
		e, mgr := newTestEngine()

		tq := tqtesting.GetTestable(c, e.cfg.Dispatcher)
		tq.CreateQueues()

		// Add the job.
		So(e.UpdateProjectJobs(c, "project", []catalog.Definition{
			{
				JobID:    "project/job-v2",
				Revision: "rev1",
				Schedule: "*/5 * * * * * *",
				Task:     noopTaskBytes(),
				Acls:     aclOne,
			},
		}), ShouldBeNil)

		// Prepare Invocation in Starting state.
		job := Job{JobID: "project/job-v2"}
		So(datastore.Get(c, &job), ShouldBeNil)
		inv, err := e.allocateInvocation(c, &job, InvocationRequest{
			IncomingTriggers: []*internal.Trigger{{Id: "a"}},
		})
		So(err, ShouldBeNil)

		callLaunchInvocation := func(c context.Context, execCount int64) error {
			return tq.ExecuteTask(c, tqtesting.Task{
				Task: &taskqueue.Task{},
				Payload: &internal.LaunchInvocationTask{
					JobId: job.JobID,
					InvId: inv.ID,
				},
			}, &taskqueue.RequestHeaders{TaskExecutionCount: execCount})
		}

		fetchInvocation := func(c context.Context) *Invocation {
			toFetch := Invocation{ID: inv.ID}
			So(datastore.Get(c, &toFetch), ShouldBeNil)
			return &toFetch
		}

		Convey("happy path", func() {
			mgr.launchTask = func(ctx context.Context, ctl task.Controller, triggers []*internal.Trigger) error {
				So(ctl.InvocationID(), ShouldEqual, inv.ID)
				So(ctl.InvocationNonce(), ShouldEqual, inv.InvocationNonce)
				ctl.DebugLog("Succeeded!")
				ctl.State().Status = task.StatusSucceeded
				return nil
			}
			So(callLaunchInvocation(c, 0), ShouldBeNil)

			updated := fetchInvocation(c)
			triggers, err := updated.IncomingTriggers()
			updated.IncomingTriggersRaw = nil

			So(err, ShouldBeNil)
			So(triggers, ShouldResemble, []*internal.Trigger{{Id: "a"}})
			So(updated, ShouldResemble, &Invocation{
				ID:              inv.ID,
				InvocationNonce: inv.ID,
				JobID:           "project/job-v2",
				IndexedJobID:    "project/job-v2",
				Started:         epoch,
				Finished:        epoch,
				Revision:        job.Revision,
				Task:            job.Task,
				Status:          task.StatusSucceeded,
				MutationsCount:  2,
				DebugLog: "[22:42:00.000] New invocation initialized\n" +
					"[22:42:00.000] Starting the invocation (attempt 1)\n" +
					"[22:42:00.000] Succeeded!\n" +
					"[22:42:00.000] Invocation finished in 0s with status SUCCEEDED\n",
			})
		})

		Convey("already aborted", func() {
			inv.Status = task.StatusAborted
			So(datastore.Put(c, inv), ShouldBeNil)
			mgr.launchTask = func(ctx context.Context, ctl task.Controller, triggers []*internal.Trigger) error {
				return fmt.Errorf("must not be called")
			}
			So(callLaunchInvocation(c, 0), ShouldBeNil)
			So(fetchInvocation(c).Status, ShouldEqual, task.StatusAborted)
		})

		Convey("retying", func() {
			// Attempt #1.
			mgr.launchTask = func(ctx context.Context, ctl task.Controller, triggers []*internal.Trigger) error {
				return transient.Tag.Apply(fmt.Errorf("oops, failed to start"))
			}
			So(callLaunchInvocation(c, 0), ShouldEqual, errRetryingLaunch)
			So(fetchInvocation(c).Status, ShouldEqual, task.StatusRetrying)

			// Attempt #2.
			mgr.launchTask = func(ctx context.Context, ctl task.Controller, triggers []*internal.Trigger) error {
				ctl.DebugLog("Succeeded!")
				ctl.State().Status = task.StatusSucceeded
				return nil
			}
			So(callLaunchInvocation(c, 1), ShouldBeNil)

			updated := fetchInvocation(c)
			So(updated.Status, ShouldEqual, task.StatusSucceeded)
			So(updated.RetryCount, ShouldEqual, 1)
			So(updated.DebugLog, ShouldEqual, "[22:42:00.000] New invocation initialized\n"+
				"[22:42:00.000] Starting the invocation (attempt 1)\n"+
				"[22:42:00.000] The invocation will be retried\n"+
				"[22:42:00.000] Starting the invocation (attempt 2)\n"+
				"[22:42:00.000] Succeeded!\n"+
				"[22:42:00.000] Invocation finished in 0s with status SUCCEEDED\n")
		})
	})
}

func TestForceInvocationV2(t *testing.T) {
	t.Parallel()

	Convey("with fake env", t, func() {
		c := newTestContext(epoch)
		e, mgr := newTestEngine()

		tq := tqtesting.GetTestable(c, e.cfg.Dispatcher)
		tq.CreateQueues()

		So(e.UpdateProjectJobs(c, "project", []catalog.Definition{
			{
				JobID:    "project/job-v2",
				Revision: "rev1",
				Schedule: "*/5 * * * * * *",
				Task:     noopTaskBytes(),
				Acls:     aclOne,
			},
		}), ShouldBeNil)

		Convey("happy path", func() {
			const expectedInvID int64 = 9200093523825174512

			futureInv, err := e.ForceInvocation(auth.WithState(c, asUserOne), "project/job-v2")
			So(err, ShouldBeNil)

			// Invocation ID is resolved right away.
			invID, err := futureInv.InvocationID(c)
			So(err, ShouldBeNil)
			So(invID, ShouldEqual, expectedInvID)

			// It is marked as active in the job state.
			job, err := e.getJob(c, "project/job-v2")
			So(err, ShouldBeNil)
			So(job.ActiveInvocations, ShouldResemble, []int64{invID})

			// All its fields are good.
			inv, err := e.getInvocation(c, "project/job-v2", invID)
			So(err, ShouldBeNil)
			So(inv, ShouldResemble, &Invocation{
				ID:              expectedInvID,
				JobID:           "project/job-v2",
				InvocationNonce: expectedInvID,
				Started:         epoch,
				TriggeredBy:     "user:one@example.com",
				Revision:        "rev1",
				Task:            noopTaskBytes(),
				Status:          task.StatusStarting,
				DebugLog: "[22:42:00.000] New invocation initialized\n" +
					"[22:42:00.000] Manually triggered by user:one@example.com\n",
			})

			// Eventually it runs the task, which then cleans up job state.
			mgr.launchTask = func(ctx context.Context, ctl task.Controller, triggers []*internal.Trigger) error {
				ctl.DebugLog("Started!")
				ctl.State().Status = task.StatusSucceeded
				return nil
			}
			tasks, _, err := tq.RunSimulation(c, nil)
			So(err, ShouldBeNil)

			// The sequence of tasks we've just performed.
			So(tasks.Payloads(), ShouldResemble, []proto.Message{
				&internal.LaunchInvocationsBatchTask{
					Tasks: []*internal.LaunchInvocationTask{{JobId: "project/job-v2", InvId: expectedInvID}},
				},
				&internal.LaunchInvocationTask{
					JobId: "project/job-v2", InvId: expectedInvID,
				},
				&internal.InvocationFinishedTask{
					JobId: "project/job-v2", InvId: expectedInvID,
				},
				&internal.TriageJobStateTask{JobId: "project/job-v2"},
			})

			// The invocation is in finished state.
			inv, err = e.getInvocation(c, "project/job-v2", invID)
			So(err, ShouldBeNil)
			So(inv, ShouldResemble, &Invocation{
				ID:              expectedInvID,
				JobID:           "project/job-v2",
				IndexedJobID:    "project/job-v2", // set for finished tasks!
				InvocationNonce: expectedInvID,
				Started:         epoch,
				Finished:        epoch.Add(time.Second),
				TriggeredBy:     "user:one@example.com",
				Revision:        "rev1",
				Task:            noopTaskBytes(),
				Status:          task.StatusSucceeded,
				MutationsCount:  2,
				DebugLog: "[22:42:00.000] New invocation initialized\n" +
					"[22:42:00.000] Manually triggered by user:one@example.com\n" +
					"[22:42:01.000] Starting the invocation (attempt 1)\n" +
					"[22:42:01.000] Started!\n" +
					"[22:42:01.000] Invocation finished in 1s with status SUCCEEDED\n",
			})

			// The job state is updated (the invocation is no longer active).
			job, err = e.getJob(c, "project/job-v2")
			So(err, ShouldBeNil)
			So(job.ActiveInvocations, ShouldBeNil)

			// The invocation is now in the list of finish invocations.
			datastore.GetTestable(c).CatchupIndexes()
			invs, _, _ := e.ListVisibleInvocations(auth.WithState(c, asUserOne), "project/job-v2", 100, "")
			So(invs, ShouldResemble, []*Invocation{inv})
		})
	})
}

func TestOneJobTriggersAnother(t *testing.T) {
	t.Parallel()

	Convey("with fake env", t, func() {
		c := newTestContext(epoch)
		e, mgr := newTestEngine()

		tq := tqtesting.GetTestable(c, e.cfg.Dispatcher)
		tq.CreateQueues()

		triggeringJob := "project/triggering-job-v2"
		triggeredJob := "project/triggered-job-v2"

		So(e.UpdateProjectJobs(c, "project", []catalog.Definition{
			{
				JobID:           triggeringJob,
				TriggeredJobIDs: []string{triggeredJob},
				Revision:        "rev1",
				Schedule:        "triggered",
				Task:            noopTaskBytes(),
				Acls:            aclOne,
			},
			{
				JobID:    triggeredJob,
				Revision: "rev1",
				Schedule: "triggered",
				Task:     noopTaskBytes(),
				Acls:     aclOne,
			},
		}), ShouldBeNil)

		Convey("happy path", func() {
			const triggeringInvID int64 = 9200093523825174512
			const triggeredInvID int64 = 9200093521728243376

			// Force launch triggering job.
			_, err := e.ForceInvocation(auth.WithState(c, asUserOne), triggeringJob)
			So(err, ShouldBeNil)

			// Eventually it runs the task which emits a bunch of triggers, which
			// cause triggered job triage, which eventually results in a new
			// invocation launch. At this point we stop and examine what we see.
			mgr.launchTask = func(ctx context.Context, ctl task.Controller, triggers []*internal.Trigger) error {
				ctl.EmitTrigger(ctx, &internal.Trigger{Id: "t1"})
				So(ctl.Save(ctx), ShouldBeNil)
				ctl.EmitTrigger(ctx, &internal.Trigger{Id: "t2"})
				ctl.State().Status = task.StatusSucceeded
				return nil
			}
			tasks, _, err := tq.RunSimulation(c, &tqtesting.SimulationParams{
				ShouldStopBefore: func(t tqtesting.Task) bool {
					task, ok := t.Payload.(*internal.LaunchInvocationsBatchTask)
					return ok && task.Tasks[0].JobId == triggeredJob
				},
			})
			So(err, ShouldBeNil)

			// How these triggers are seen from outside the task.
			expectedTrigger1 := &internal.Trigger{
				Id:           "t1",
				JobId:        triggeringJob,
				InvocationId: triggeringInvID,
				Created:      google.NewTimestamp(epoch.Add(1 * time.Second)),
			}
			expectedTrigger2 := &internal.Trigger{
				Id:           "t2",
				JobId:        triggeringJob,
				InvocationId: triggeringInvID,
				Created:      google.NewTimestamp(epoch.Add(1 * time.Second)),
			}

			// All the tasks we've just executed.
			So(tasks.Payloads(), ShouldResemble, []proto.Message{
				// Triggering job begins execution.
				&internal.LaunchInvocationsBatchTask{
					Tasks: []*internal.LaunchInvocationTask{{JobId: triggeringJob, InvId: triggeringInvID}},
				},
				&internal.LaunchInvocationTask{
					JobId: triggeringJob, InvId: triggeringInvID,
				},

				// It emits a trigger in the middle.
				&internal.FanOutTriggersTask{
					JobIds:   []string{triggeredJob},
					Triggers: []*internal.Trigger{expectedTrigger1},
				},
				&internal.EnqueueTriggersTask{
					JobId:    triggeredJob,
					Triggers: []*internal.Trigger{expectedTrigger1},
				},

				// Triggering job finishes execution, emitting another trigger.
				&internal.InvocationFinishedTask{
					JobId: triggeringJob,
					InvId: triggeringInvID,
					Triggers: &internal.FanOutTriggersTask{
						JobIds:   []string{triggeredJob},
						Triggers: []*internal.Trigger{expectedTrigger2},
					},
				},
				&internal.EnqueueTriggersTask{
					JobId:    triggeredJob,
					Triggers: []*internal.Trigger{expectedTrigger2},
				},

				// Triggered job is getting triaged (because pending triggers).
				&internal.TriageJobStateTask{
					JobId: triggeredJob,
				},

				// Triggering job is getting triaged (because it has just finished).
				&internal.TriageJobStateTask{
					JobId: triggeringJob,
				},
			})

			// At this point triggered job is just about to start.

			// Triggering invocation has finished (with triggers recorded).
			triggeringInv, err := e.getInvocation(c, triggeringJob, triggeringInvID)
			So(err, ShouldBeNil)
			So(triggeringInv.Status, ShouldEqual, task.StatusSucceeded)
			outgoing, err := triggeringInv.OutgoingTriggers()
			So(err, ShouldBeNil)
			So(outgoing, ShouldResemble, []*internal.Trigger{expectedTrigger1, expectedTrigger2})

			// Now we resume the simulation. It will start the triggered invocation
			// and run it.
			var seen []*internal.Trigger
			mgr.launchTask = func(ctx context.Context, ctl task.Controller, triggers []*internal.Trigger) error {
				seen = triggers
				ctl.State().Status = task.StatusSucceeded
				return nil
			}
			tasks, _, err = tq.RunSimulation(c, nil)
			So(err, ShouldBeNil)

			// All the tasks we've just executed.
			So(tasks.Payloads(), ShouldResemble, []proto.Message{
				// The triggered job begins execution.
				&internal.LaunchInvocationsBatchTask{
					Tasks: []*internal.LaunchInvocationTask{{JobId: triggeredJob, InvId: triggeredInvID}},
				},
				&internal.LaunchInvocationTask{
					JobId: triggeredJob, InvId: triggeredInvID,
				},
				// ...and finishes. Note that the triage doesn't launch new invocation.
				&internal.InvocationFinishedTask{
					JobId: triggeredJob, InvId: triggeredInvID,
				},
				&internal.TriageJobStateTask{JobId: triggeredJob},
			})

			// Verify LaunchTask callback saw the triggers.
			So(seen, ShouldResemble, []*internal.Trigger{expectedTrigger1, expectedTrigger2})

			// And they are recoded in IncomingTriggers set.
			triggeredInv, err := e.getInvocation(c, triggeredJob, triggeredInvID)
			So(err, ShouldBeNil)
			So(triggeredInv.Status, ShouldEqual, task.StatusSucceeded)
			incoming, err := triggeredInv.IncomingTriggers()
			So(err, ShouldBeNil)
			So(incoming, ShouldResemble, []*internal.Trigger{expectedTrigger1, expectedTrigger2})
		})
	})
}

func TestInvocationTimers(t *testing.T) {
	t.Parallel()

	Convey("with fake env", t, func() {
		c := newTestContext(epoch)
		e, mgr := newTestEngine()

		tq := tqtesting.GetTestable(c, e.cfg.Dispatcher)
		tq.CreateQueues()

		const testJobID = "project/job-v2"
		So(e.UpdateProjectJobs(c, "project", []catalog.Definition{
			{
				JobID:    testJobID,
				Revision: "rev1",
				Schedule: "triggered",
				Task:     noopTaskBytes(),
				Acls:     aclOne,
			},
		}), ShouldBeNil)

		Convey("happy path", func() {
			const testInvID int64 = 9200093523825174512

			// Force launch the job.
			_, err := e.ForceInvocation(auth.WithState(c, asUserOne), testJobID)
			So(err, ShouldBeNil)

			// See handelTimer. Name of the timer => time since epoch.
			callTimes := map[string]time.Duration{}

			// Eventually it runs the task which emits a bunch of timers and then
			// some more, and then stops.
			mgr.launchTask = func(ctx context.Context, ctl task.Controller, triggers []*internal.Trigger) error {
				ctl.AddTimer(ctx, time.Minute, "1 min", []byte{1})
				ctl.AddTimer(ctx, 2*time.Minute, "2 min", []byte{2})
				ctl.State().Status = task.StatusRunning
				return nil
			}
			mgr.handleTimer = func(ctx context.Context, ctl task.Controller, name string, payload []byte) error {
				callTimes[name] = clock.Now(ctx).Sub(epoch)
				switch name {
				case "1 min": // ignore
				case "2 min":
					// Call us again later.
					ctl.AddTimer(ctx, time.Minute, "stop", []byte{3})
				case "stop":
					ctl.AddTimer(ctx, time.Minute, "ignored-timer", nil)
					ctl.State().Status = task.StatusSucceeded
				}
				return nil
			}
			tasks, _, err := tq.RunSimulation(c, nil)
			So(err, ShouldBeNil)

			timerMsg := func(idSuffix string, created, eta time.Duration, title string, payload []byte) *internal.Timer {
				return &internal.Timer{
					Id:      fmt.Sprintf("%s:%d:%s", testJobID, testInvID, idSuffix),
					Created: google.NewTimestamp(epoch.Add(created)),
					Eta:     google.NewTimestamp(epoch.Add(eta)),
					Title:   title,
					Payload: payload,
				}
			}

			// Individual timers emitted by the test. Note that 1 extra sec comes from
			// the delay added by kickLaunchInvocationsBatchTask.
			timer1 := timerMsg("1:0", time.Second, time.Second+time.Minute, "1 min", []byte{1})
			timer2 := timerMsg("1:1", time.Second, time.Second+2*time.Minute, "2 min", []byte{2})
			timer3 := timerMsg("3:0", time.Second+2*time.Minute, time.Second+3*time.Minute, "stop", []byte{3})

			// All 'handleTimer' ticks happened at expected moments in time.
			So(callTimes, ShouldResemble, map[string]time.Duration{
				"1 min": time.Second + time.Minute,
				"2 min": time.Second + 2*time.Minute,
				"stop":  time.Second + 3*time.Minute,
			})

			// All the tasks we've just executed.
			So(tasks.Payloads(), ShouldResemble, []proto.Message{
				// Triggering job begins execution.
				&internal.LaunchInvocationsBatchTask{
					Tasks: []*internal.LaunchInvocationTask{{JobId: testJobID, InvId: testInvID}},
				},
				&internal.LaunchInvocationTask{
					JobId: testJobID, InvId: testInvID,
				},

				// Request to schedule a bunch of timers.
				&internal.ScheduleTimersTask{
					JobId:  testJobID,
					InvId:  testInvID,
					Timers: []*internal.Timer{timer1, timer2},
				},

				// Actual individual timers.
				&internal.TimerTask{
					JobId: testJobID,
					InvId: testInvID,
					Timer: timer1,
				},
				&internal.TimerTask{
					JobId: testJobID,
					InvId: testInvID,
					Timer: timer2,
				},

				// One more, scheduled from handleTimer.
				&internal.TimerTask{
					JobId: testJobID,
					InvId: testInvID,
					Timer: timer3,
				},

				// End of the invocation.
				&internal.InvocationFinishedTask{
					JobId: testJobID, InvId: testInvID,
				},
				&internal.TriageJobStateTask{JobId: testJobID},
			})
		})
	})
}
