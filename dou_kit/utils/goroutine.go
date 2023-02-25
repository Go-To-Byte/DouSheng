// Author: BeYoung
// Date: 2023/2/24 20:56
// Software: GoLand

package utils

import (
	"context"
	"sync"
	"time"
)

// GoRunGrpc is used run a goroutine of grpc
// You need give context, grpc function and request
// the grpc's return is filled with response and error
// Example:
//
//	wait := sync.WaitGroup{}
//	wait.Add(1)
//	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
//	defer cancel()
//	run := GoRunGrpc{
//		Ctx:      &ctx,
//		Wait:     &wait,
//		Request:  user.NewUserInfoRequest(),
//		Rpc:      h.service.UserInfo,
//	}
//	go run.Run()
type GoRunGrpc struct {
	Err      error
	Ctx      *context.Context
	Wait     *sync.WaitGroup
	Request  any
	Response any
	Rpc      func(ctx context.Context, request any) (response any, err error) // grpc function
}

func NewGoRunGrpc() GoRunGrpc {
	return GoRunGrpc{}
}

func (g GoRunGrpc) Run() {
	defer g.Wait.Done()
	for {
		select {
		case <-(*g.Ctx).Done():
			return
		default:
			g.Response, g.Err = g.Rpc(*g.Ctx, g.Request)
		}
	}
}

// GORUN You need give context, grpc function, request, and response.
//
// rpc is a function:
//
//	func(ctx context.Context, request interface{}) (response interface{}, err error)
//
// response is the return value of rpc function, it must be a pointer.
//
// Example:
//
//	wait := sync.WaitGroup{}
//	wait.Add(1)
//	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
//	go GORUN(ctx, wait, user.NewUserInfoRequest, h.service.UserInfo, userInfoResp{})
func GORUN(
	ctx *context.Context, // context.WithTimeout()
	wait *sync.WaitGroup, // sync.WaitGroup{}
	rpc func(ctx context.Context, request any) (response any, err error),
	req any,
	resp *any,
) (err error) {
	defer wait.Done()
	for {
		select {
		case <-(*ctx).Done():
			return
		default:
			r, e := rpc(*ctx, req)
			resp = &r
			err = e
		}
	}
}

func example() {
	wait := sync.WaitGroup{}
	wait.Add(1)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	run := GoRunGrpc{
		Ctx:     &ctx,
		Wait:    &wait,
		Request: nil,
		Rpc:     nil,
	}
	go run.Run()
}
