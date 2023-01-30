// Author: BeYoung
// Date: 2023/1/29 17:57
// Software: GoLand

package main

import (
	"context"
	"flag"
	_ "github.com/Go-To-Byte/DouSheng/apps/message/init"
	"github.com/Go-To-Byte/DouSheng/apps/message/proto"
	"go.uber.org/zap"
	"time"

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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// r1, err := c.Follow(ctx, &pb.FollowRequest{
	// 	UserId:     1619766175401512960,
	// 	ToUserId:   1619974131258757120,
	// 	ActionType: 1,
	// })
	// zap.S().Infof("Follow: %+v", r1)
	//
	// r1, err = c.Follow(ctx, &pb.FollowRequest{
	// 	UserId:     1618553379024277504,
	// 	ToUserId:   1619974131258757120,
	// 	ActionType: 1,
	// })
	// zap.S().Infof("Follow: %+v", r1)
	//
	// r1, err = c.Follow(ctx, &pb.FollowRequest{
	// 	UserId:     1619974131258757120,
	// 	ToUserId:   1618553379024277504,
	// 	ActionType: 1,
	// })
	// zap.S().Infof("Follow: %+v", r1)

	if err != nil {
		zap.S().Infof("could not greet: %v", err)
	}
}
