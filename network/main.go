// Author: BeYoung
// Date: 2023/1/31 22:13
// Software: GoLand

package main

import (
	"context"
	_ "github.com/Go-To-Byte/DouSheng/network/inits"
	"github.com/Go-To-Byte/DouSheng/network/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	defer func(grpcConn map[string]*grpc.ClientConn) {
		for k := range grpcConn {
			err := grpcConn[k].Close()
			if err != nil {
				zap.S().Errorf("Grpc Client Connect Close Failed: %v", k)
				continue
			}
		}
	}(models.Dials)

	router := models.Router
	router.GET("/", func(c *gin.Context) {
		time.Sleep(time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: models.Router,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.S().Infof("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	zap.S().Errorf("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		zap.S().Errorf("Server Shutdown: %v", err)
	}
	zap.S().Infof("Server exiting")
}
