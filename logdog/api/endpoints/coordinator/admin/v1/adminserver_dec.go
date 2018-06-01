// Code generated by svcdec; DO NOT EDIT

package logdog

import (
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"

	empty "github.com/golang/protobuf/ptypes/empty"
)

type DecoratedAdmin struct {
	// Service is the service to decorate.
	Service AdminServer
	// Prelude is called for each method before forwarding the call to Service.
	// If Prelude returns an error, then the call is skipped and the error is
	// processed via the Postlude (if one is defined), or it is returned directly.
	Prelude func(c context.Context, methodName string, req proto.Message) (context.Context, error)
	// Postlude is called for each method after Service has processed the call, or
	// after the Prelude has returned an error. This takes the the Service's
	// response proto (which may be nil) and/or any error. The decorated
	// service will return the response (possibly mutated) and error that Postlude
	// returns.
	Postlude func(c context.Context, methodName string, rsp proto.Message, err error) error
}

func (s *DecoratedAdmin) SetConfig(c context.Context, req *SetConfigRequest) (rsp *empty.Empty, err error) {
	var newCtx context.Context
	if s.Prelude != nil {
		newCtx, err = s.Prelude(c, "SetConfig", req)
	}
	if err == nil {
		c = newCtx
		rsp, err = s.Service.SetConfig(c, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(c, "SetConfig", rsp, err)
	}
	return
}
