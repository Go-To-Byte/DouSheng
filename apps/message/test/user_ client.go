// Author: BeYoung
// Date: 2023/1/29 17:57
// Software: GoLand

package main

import (
	"context"
	"flag"
	"fmt"
	_ "github.com/Go-To-Byte/DouSheng/apps/message/init"
	"github.com/Go-To-Byte/DouSheng/apps/message/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		zap.S().Infof("did not connect: %v", err)
	}
	defer conn.Close()
	c := proto.NewChatClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i := 0; i < 5; i++ {
		r, err := c.MessageAction(ctx, &proto.MessageRequest{
			UserId:     1619766175401512960,
			ToUserId:   1619974131258757120,
			ActionType: 0,
			Content:    fmt.Sprintf("hello go %v", i),
		})
		if err != nil {
			zap.S().Infof("message(%v) error: %+v", i, err)
		}
		zap.S().Infof("message(%v) ok: %+v", i, r)
	}

	for i := 5; i < 10; i++ {
		r, err := c.MessageAction(ctx, &proto.MessageRequest{
			UserId:     1619974131258757120,
			ToUserId:   1619766175401512960,
			ActionType: 0,
			Content:    fmt.Sprintf("hello go %v", i),
		})
		if err != nil {
			zap.S().Infof("message(%v) error: %+v", i, err)
		}
		zap.S().Infof("message(%v) ok: %+v", i, r)
	}

	h, err := c.MessageHistory(ctx, &proto.MessageListRequest{
		UserId:   1619974131258757120,
		ToUserId: 1619766175401512960,
	})
	for i := 0; i < len(h.MessageList); i++ {
		zap.S().Infof("message(%v)", h.MessageList[i])
	}
	if err != nil {
		zap.S().Infof("could not greet: %v", err)
	}
}
