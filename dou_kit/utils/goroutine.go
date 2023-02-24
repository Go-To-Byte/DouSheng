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
	Request  interface{}
	Response interface{}
	Rpc      func(interface{}, interface{}) (interface{}, error) // grpc function
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

func GORUN(ctx *context.Context, wait *sync.WaitGroup,
	rpc func(interface{}, interface{}) (interface{}, error),
	req interface{}) (resp interface{}, err error) {
	defer wait.Done()
	for {
		select {
		case <-(*ctx).Done():
			return
		default:
			resp, err = rpc(*ctx, req)
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
