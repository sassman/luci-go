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
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"

	"go.chromium.org/gae/service/datastore"
	"go.chromium.org/gae/service/memcache"

	"go.chromium.org/luci/common/clock"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/common/retry/transient"

	"go.chromium.org/luci/scheduler/appengine/internal"
)

// assertInTransaction panics if the context is not transactional.
func assertInTransaction(c context.Context) {
	if datastore.CurrentTransaction(c) == nil {
		panic("expecting to be called from inside a transaction")
	}
}

// debugLog mutates a string by appending a line to it.
func debugLog(c context.Context, str *string, format string, args ...interface{}) {
	prefix := clock.Now(c).UTC().Format("[15:04:05.000] ")
	*str += prefix + fmt.Sprintf(format+"\n", args...)
}

// defaultTransactionOptions is used for all transactions.
//
// Almost all transactions done by the scheduler service happen in background
// task queues, it is fine to retry more there.
var defaultTransactionOptions = datastore.TransactionOptions{
	Attempts: 10,
}

// abortTransaction makes the error abort the transaction (even if it is marked
// as transient).
//
// See runTxn for more info. This is used primarily by errUpdateConflict.
var abortTransaction = errors.BoolTag{Key: errors.NewTagKey("this error aborts the transaction")}

// runTxn runs a datastore transaction retrying the body on transient errors or
// when encountering a commit conflict.
//
// It will NOT retry errors (even if transient) marked with abortTransaction
// tag. This is primarily used to tag errors that are transient at a level
// higher than the transaction: errors marked with both transient.Tag and
// abortTransaction are not retried by runTxn, but may be retried by something
// on top (like Task Queue).
func runTxn(c context.Context, cb func(context.Context) error) error {
	var attempt int
	var innerErr error

	err := datastore.RunInTransaction(c, func(c context.Context) error {
		attempt++
		if attempt != 1 {
			if innerErr != nil {
				logging.Warningf(c, "Retrying the transaction after the error: %s", innerErr)
			} else {
				logging.Warningf(c, "Retrying the transaction: failed to commit")
			}
		}
		innerErr = cb(c)
		if transient.Tag.In(innerErr) && !abortTransaction.In(innerErr) {
			return datastore.ErrConcurrentTransaction // causes a retry
		}
		return innerErr
	}, &defaultTransactionOptions)

	if err != nil {
		logging.WithError(err).Errorf(c, "Transaction failed")
		if innerErr != nil {
			return innerErr
		}
		// Here it can only be a commit error (i.e. produced by RunInTransaction
		// itself, not by its callback). We treat them as transient.
		return transient.Tag.Apply(err)
	}

	return nil
}

// runIsolatedTxn is like runTxn, except it executes the callback in a new
// isolated transaction (even if the original context is already transactional).
func runIsolatedTxn(c context.Context, cb func(context.Context) error) error {
	return runTxn(datastore.WithoutTransaction(c), cb)
}

// equalSortedLists returns true if lists contain the same sequence of strings.
func equalSortedLists(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, s := range a {
		if s != b[i] {
			return false
		}
	}
	return true
}

// equalInt64Lists returns true if two lists of int64 are equal.
//
// Order is important.
func equalInt64Lists(a, b []int64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, s := range a {
		if s != b[i] {
			return false
		}
	}
	return true
}

// marshalTriggersList serializes list of triggers.
//
// Panics on errors.
func marshalTriggersList(t []*internal.Trigger) []byte {
	if len(t) == 0 {
		return nil
	}
	blob, err := proto.Marshal(&internal.TriggerList{Triggers: t})
	if err != nil {
		panic(err)
	}
	return blob
}

// unmarshalTriggersList deserializes list of triggers.
func unmarshalTriggersList(blob []byte) ([]*internal.Trigger, error) {
	if len(blob) == 0 {
		return nil, nil
	}
	list := internal.TriggerList{}
	if err := proto.Unmarshal(blob, &list); err != nil {
		return nil, err
	}
	return list.Triggers, nil
}

// mutateTriggersList deserializes the list, calls a callback, which modifies
// the list and serializes it back.
func mutateTriggersList(blob *[]byte, cb func(*[]*internal.Trigger)) error {
	list, err := unmarshalTriggersList(*blob)
	if err != nil {
		return err
	}
	cb(&list)
	*blob = marshalTriggersList(list)
	return nil
}

// marshalTimersList serializes list of timers.
//
// Panics on errors.
func marshalTimersList(t []*internal.Timer) []byte {
	if len(t) == 0 {
		return nil
	}
	blob, err := proto.Marshal(&internal.TimerList{Timers: t})
	if err != nil {
		panic(err)
	}
	return blob
}

// unmarshalTimersList deserializes list of timers.
func unmarshalTimersList(blob []byte) ([]*internal.Timer, error) {
	if len(blob) == 0 {
		return nil, nil
	}
	list := internal.TimerList{}
	if err := proto.Unmarshal(blob, &list); err != nil {
		return nil, err
	}
	return list.Timers, nil
}

// mutateTimersList deserializes the list, calls a callback, which modifies
// the list and serializes it back.
func mutateTimersList(blob *[]byte, cb func(*[]*internal.Timer)) error {
	list, err := unmarshalTimersList(*blob)
	if err != nil {
		return err
	}
	cb(&list)
	*blob = marshalTimersList(list)
	return nil
}

// opsCache "remembers" recently executed operations, and skips executing them
// if they already were done.
//
// Expected cardinality of a set of all possible actions should be small (we
// store the cache in memory).
type opsCache struct {
	lock      sync.RWMutex
	doneFlags map[string]bool
}

// Do calls callback only if it wasn't called before.
//
// Works on best effort basis: callback can and will be called multiple times
// (just not the every time 'Do' is called).
//
// Keeps "done" flag in local memory and in memcache (using 'key' as
// identifier). The callback should be idempotent, since it still may be called
// multiple times if multiple processes attempt to execute the action at once.
func (o *opsCache) Do(c context.Context, key string, cb func() error) error {
	// Check the local cache.
	if o.getFlag(key) {
		return nil
	}

	// Check the global cache.
	switch _, err := memcache.GetKey(c, key); {
	case err == nil:
		o.setFlag(key)
		return nil
	case err == memcache.ErrCacheMiss:
		break
	default:
		logging.WithError(err).Warningf(c, "opsCache failed to check memcache, will proceed executing op")
	}

	// Do it.
	if err := cb(); err != nil {
		return err
	}

	// Store in the local cache.
	o.setFlag(key)

	// Store in the global cache. Ignore errors, it's not a big deal.
	item := memcache.NewItem(c, key)
	item.SetValue([]byte("ok"))
	item.SetExpiration(24 * time.Hour)
	if err := memcache.Set(c, item); err != nil {
		logging.WithError(err).Warningf(c, "opsCache failed to write item to memcache")
	}

	return nil
}

func (o *opsCache) getFlag(key string) bool {
	o.lock.RLock()
	defer o.lock.RUnlock()
	return o.doneFlags[key]
}

func (o *opsCache) setFlag(key string) {
	o.lock.Lock()
	defer o.lock.Unlock()
	if o.doneFlags == nil {
		o.doneFlags = map[string]bool{}
	}
	o.doneFlags[key] = true
}
