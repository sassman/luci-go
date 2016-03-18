// Copyright 2016 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package buildbot

import (
	"github.com/julienschmidt/httprouter"

	"github.com/luci/luci-go/server/discovery"
	"github.com/luci/luci-go/server/middleware"
	"github.com/luci/luci-go/server/prpc"

	"github.com/luci/luci-go/common/prpc/talk/buildbot/proto"
)

func InstallAPIRoutes(router *httprouter.Router, base middleware.Base) {
	server := &prpc.Server{}
	buildbot.RegisterBuildbotServer(server, &buildbotService{})
	discovery.Enable(server)
	server.InstallHandlers(router, base)
}
