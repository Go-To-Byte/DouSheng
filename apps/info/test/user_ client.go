// Author: BeYoung
// Date: 2023/1/29 17:57
// Software: GoLand

package main

import (
	"context"
	"flag"
	"go.uber.org/zap"
	"time"

	_ "github.com/Go-To-Byte/DouSheng/apps/user/init"
	pb "github.com/Go-To-Byte/DouSheng/apps/user/proto"
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
	c := pb.NewUserClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r1, err := c.Register(ctx, &pb.RegisterRequest{
		Username: "atest",
		Password: "apptest",
	})
	zap.S().Infof("Registered: %+v", r1)

	r2, err := c.Login(ctx, &pb.LoginRequest{
		Username: "atest",
		Password: "apptest",
	})
	zap.S().Infof("Registered: %+v", r2)

	if err != nil {
		zap.S().Infof("could not greet: %v", err)
	}
}
