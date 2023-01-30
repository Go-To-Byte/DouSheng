// Author: BeYoung
// Date: 2023/1/29 17:57
// Software: GoLand

package main

import (
	"context"
	"flag"
	_ "github.com/Go-To-Byte/DouSheng/apps/comment/init"
	"github.com/Go-To-Byte/DouSheng/apps/comment/proto"
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
	c := proto.NewCommentClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	f, err := c.Comment(ctx, &proto.CommentRequest{
		UserId:     0,
		VideoId:    0,
		ActionType: 0,
		Content:    "",
		CommentId:  0,
	})
	zap.S().Infof("comment status: %v", f.StatusMsg)

	if err != nil {
		zap.S().Infof("could not greet: %v", err)
	}
}
