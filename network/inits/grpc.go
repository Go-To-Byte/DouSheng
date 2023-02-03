// Author: BeYoung
// Date: 2023/2/1 1:18
// Software: GoLand

package inits

import (
	"fmt"
	"github.com/Go-To-Byte/DouSheng/network/models"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func initGrpc() {
	// TODO: goroutines
	initFeedClient()
	initUserClient()
	initVideoClient()
	initCommentClient()
	initMessageClient()
	initRelationClient()
	initFavoriteClient()
}

func initFeedClient() {
	// TODO
	return
}

func initUserClient() {
	// target := fmt.Sprintf("%v:%v", models.Config.GrpcConfig.Host, models.Config.GrpcConfig.Port)
	targetUser := fmt.Sprintf("consul://%s:%d/%s?wait=14s",
		models.Config.Consul.Host, models.Config.Consul.Port, models.Config.GrpcName.User)
	if dial, err := grpc.Dial(targetUser,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`)); err != nil {
		zap.S().Errorf("grpc dial failed: %v", err)
	} else {
		zap.S().Infof("grpc dial connect: %v", targetUser)
		models.Dials["user"] = dial
	}

}

func initVideoClient() {

	targetVideo := fmt.Sprintf("consul://%v:%v/%v?wait=14s",
		models.Config.Consul.Host, models.Config.Consul.Port, models.Config.GrpcName.Video)
	if dial, err := grpc.Dial(targetVideo,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`)); err != nil {
		zap.S().Errorf("grpc dial failed: %v", err)
	} else {
		zap.S().Infof("grpc dial connect: %v", targetVideo)
		models.Dials["video"] = dial
	}

}

func initCommentClient() {
	targetComment := fmt.Sprintf("consul://%v:%v/%v?wait=14s",
		models.Config.Consul.Host, models.Config.Consul.Port, models.Config.GrpcName.Comment)
	if dial, err := grpc.Dial(targetComment,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`)); err != nil {
		zap.S().Errorf("grpc dial failed: %v", err)
	} else {
		zap.S().Infof("grpc dial connect: %v", targetComment)
		models.Dials["comment"] = dial
	}

}

func initMessageClient() {

	targetMessage := fmt.Sprintf("consul://%v:%v/%v?wait=14s", models.Config.Consul.Host, models.Config.Consul.Port, models.Config.GrpcName.Massage)
	if dial, err := grpc.Dial(targetMessage,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`)); err != nil {
		zap.S().Errorf("grpc dial failed: %v", err)
	} else {
		zap.S().Infof("grpc dial connect: %v", targetMessage)
		models.Dials["message"] = dial
	}

}

func initRelationClient() {
	targetRelation := fmt.Sprintf("consul://%v:%v/%v?wait=14s", models.Config.Consul.Host, models.Config.Consul.Port, models.Config.GrpcName.Relation)
	if dial, err := grpc.Dial(targetRelation,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`)); err != nil {
		zap.S().Errorf("grpc dial failed: %v", err)
	} else {
		zap.S().Infof("grpc dial connect: %v", targetRelation)
		models.Dials["relation"] = dial
	}

}

func initFavoriteClient() {
	targetFavorite := fmt.Sprintf("consul://%v:%v/%v?wait=14s", models.Config.Consul.Host, models.Config.Consul.Port, models.Config.GrpcName.Favorite)
	if dial, err := grpc.Dial(targetFavorite,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`)); err != nil {
		zap.S().Errorf("grpc dial failed: %v", err)
	} else {
		zap.S().Infof("grpc dial connect: %v", targetFavorite)
		models.Dials["favorite"] = dial
	}

}
