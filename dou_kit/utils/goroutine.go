// Author: BeYoung
// Date: 2023/2/24 20:56
// Software: GoLand

package utils

import (
	"context"
	"errors"
	"sync"
	"time"
)

// GoRunGrpc is used run a goroutine of grpc
// You need give context, grpc function and request
// the grpc's return is filled with response and error
// Example:
//
//	func example() {
//		wait := sync.WaitGroup{}
//		wait.Add(1)
//		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
//		defer cancel()
//		run := NewGoRunGrpc(&ctx, &wait, nil, nil)
//		go run.Run()
//		wait.Wait()
//	}
type GoRunGrpc struct {
	ctx      *context.Context
	wait     *sync.WaitGroup
	rpc      func(ctx context.Context, request any) (response any, err error) // grpc function
	req      any
	Err      error
	Response any
}

func NewGoRunGrpc(
	waitGroup *[]*sync.WaitGroup,
	rpc func(ctx context.Context, request any) (response any, err error),
	req any,
) (goRun GoRunGrpc, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	wait := sync.WaitGroup{}
	wait.Add(1)
	if waitGroup == nil {
		err = errors.New("waitGroup is nil")
		return GoRunGrpc{Err: err}, err
	}
	*waitGroup = append(*waitGroup, &wait)
	return GoRunGrpc{
		ctx:  &ctx,
		wait: &wait,
		rpc:  rpc,
		req:  req,
	}, nil
}

func (g GoRunGrpc) Run() {
	defer g.wait.Done()
	for {
		select {
		case <-(*g.ctx).Done():
			return
		default:
			g.Response, g.Err = g.rpc(*(g.ctx), g.req)
			return
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
//	go GORUN(ctx, wait, h.service.UserInfo, user.NewUserInfoRequest, userInfoResp{})
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
			return
		}
	}
}

func example() {
	var waitGroup []*sync.WaitGroup
	run, _ := NewGoRunGrpc(&waitGroup, nil, nil)
	go run.Run()
	for _, wait := range waitGroup {
		(*wait).Wait()
	}
}
