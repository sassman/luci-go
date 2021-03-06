// Copyright 2019 The LUCI Authors.
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

package lucicfg

import (
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"go.chromium.org/luci/common/api/luci_config/config/v1"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
)

func TestConfigSet(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	Convey("With temp dir", t, func() {
		tmp, err := ioutil.TempDir("", "lucicfg")
		So(err, ShouldBeNil)
		defer os.RemoveAll(tmp)

		path := func(p ...string) string {
			return filepath.Join(append([]string{tmp}, p...)...)
		}

		read := func(p ...string) []byte {
			body, err := ioutil.ReadFile(path(p...))
			So(err, ShouldBeNil)
			return body
		}

		So(os.Mkdir(path("subdir"), 0700), ShouldBeNil)
		So(ioutil.WriteFile(path("a.cfg"), []byte("a\n"), 0600), ShouldBeNil)
		So(ioutil.WriteFile(path("subdir", "b.cfg"), []byte("b\n"), 0600), ShouldBeNil)

		Convey("Reading", func() {
			Convey("Success", func() {
				cfg, err := ReadConfigSet(tmp)
				So(err, ShouldBeNil)
				So(cfg, ShouldResemble, ConfigSet{
					"a.cfg":        []byte("a\n"),
					"subdir/b.cfg": []byte("b\n"),
				})

				So(cfg.Files(), ShouldResemble, []string{
					"a.cfg",
					"subdir/b.cfg",
				})
			})

			Convey("Missing dir", func() {
				_, err := ReadConfigSet(path("unknown"))
				So(err, ShouldNotBeNil)
			})
		})

		Convey("Writing", func() {
			cs := ConfigSet{
				"a":     []byte("111"),
				"dir/a": []byte("222"),
			}
			changed, unchanged, err := cs.Write(tmp)
			So(changed, ShouldResemble, []string{"a", "dir/a"})
			So(unchanged, ShouldHaveLength, 0)
			So(err, ShouldBeNil)

			So(read("a"), ShouldResemble, []byte("111"))
			So(read("dir/a"), ShouldResemble, []byte("222"))

			cs["a"] = []byte("333")
			changed, unchanged, err = cs.Write(tmp)
			So(changed, ShouldResemble, []string{"a"})
			So(unchanged, ShouldResemble, []string{"dir/a"})
			So(err, ShouldBeNil)

			So(read("a"), ShouldResemble, []byte("333"))
		})

		Convey("DiscardChangesToUntracked", func() {
			original, _ := ReadConfigSet(tmp)

			generated := func() ConfigSet {
				return ConfigSet{
					"a.cfg":        []byte("new a\n"),
					"subdir/b.cfg": []byte("new b\n"),
				}
			}

			Convey("No untracked", func() {
				cs := generated()
				So(cs.DiscardChangesToUntracked(ctx, []string{"*"}, "-"), ShouldBeNil)
				So(cs, ShouldResemble, generated())
			})

			Convey("Untracked files are restored from disk", func() {
				cs := generated()
				So(cs.DiscardChangesToUntracked(ctx, []string{"!b.cfg"}, tmp), ShouldBeNil)
				So(cs, ShouldResemble, ConfigSet{
					"a.cfg":        generated()["a.cfg"],
					"subdir/b.cfg": original["subdir/b.cfg"],
				})
			})

			Convey("Untracked files are discarded when dumping to stdout", func() {
				cs := generated()
				So(cs.DiscardChangesToUntracked(ctx, []string{"!b.cfg"}, "-"), ShouldBeNil)
				So(cs, ShouldResemble, ConfigSet{
					"a.cfg": generated()["a.cfg"],
				})
			})

			Convey("Untracked files are discarded if don't exist on disk", func() {
				cs := ConfigSet{"c.cfg": []byte("generated")}
				So(cs.DiscardChangesToUntracked(ctx, []string{"!c.cfg"}, tmp), ShouldBeNil)
				So(cs, ShouldResemble, ConfigSet{})
			})
		})
	})

	Convey("Digests", t, func() {
		cs := ConfigSet{"a": nil, "b": []byte("123")}
		So(cs.Digests(), ShouldResemble, map[string]string{
			"a": "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
			"b": "a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3",
		})
	})

	Convey("Validation", t, func() {
		validator := testValidator{
			res: &ValidationResult{
				Messages: []*ValidationMessage{
					{Severity: "ERROR", Text: "Boo"},
				},
			},
		}

		cfg := ConfigSet{"a.cfg": []byte("aaa"), "b.cfg": []byte{0, 1, 2}}

		res, err := cfg.Validate(ctx, "config set name", &validator)
		So(err, ShouldBeNil)
		So(res, ShouldResemble, validator.res)

		So(validator.req, ShouldResemble, &ValidationRequest{
			ConfigSet: "config set name",
			Files: []*config.LuciConfigValidateConfigRequestMessageFile{
				{Path: "a.cfg", Content: "YWFh"},
				{Path: "b.cfg", Content: "AAEC"},
			},
		})
	})

	Convey("Overall error check", t, func() {
		result := func(level ...string) *ValidationResult {
			res := &ValidationResult{}
			for _, l := range level {
				res.Messages = append(res.Messages, &ValidationMessage{
					Severity: l,
					Text:     "boo",
				})
			}
			res.Log(ctx) // for code coverage
			return res
		}

		// Fail on warnings = false.
		So(result().OverallError(false), ShouldBeNil)
		So(result("INFO", "WARNING").OverallError(false), ShouldBeNil)
		So(result("INFO", "ERROR").OverallError(false), ShouldErrLike, "some files were invalid")

		// Fail on warnings = true.
		So(result().OverallError(true), ShouldBeNil)
		So(result("INFO").OverallError(true), ShouldBeNil)
		So(result("INFO", "WARNING", "ERROR").OverallError(true), ShouldErrLike, "some files were invalid")
		So(result("INFO", "WARNING").OverallError(true), ShouldErrLike, "some files had validation warnings")
	})
}

type testValidator struct {
	req *ValidationRequest // captured request
	res *ValidationResult  // a reply to send
}

func (t *testValidator) Validate(ctx context.Context, req *ValidationRequest) (*ValidationResult, error) {
	t.req = req
	return t.res, nil
}
