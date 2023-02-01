// Author: BeYoung
// Date: 2023/1/25 23:50
// Software: GoLand

package main

import "fmt"

func main() {
	// router := models.Router
	// router.GET("/", func(c *gin.Context) {
	// 	time.Sleep(time.Second)
	// 	c.String(http.StatusOK, "Welcome Gin Server")
	// })
	//
	// srv := &http.Server{
	// 	Addr:    ":8080",
	// 	Handler: models.Router,
	// }
	//
	// go func() {
	// 	// 服务连接
	// 	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	// 		log.Fatalf("listen: %s\n", err)
	// 	}
	// }()
	//
	// // 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	// quit := make(chan os.Signal)
	// signal.Notify(quit, os.Interrupt)
	// <-quit
	// log.Println("Shutdown Server ...")
	//
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	// if err := srv.Shutdown(ctx); err != nil {
	// 	log.Fatal("Server Shutdown:", err)
	// }
	// log.Println("Server exiting")
	i64 := make([]int64, 2)
	i64 = append(i64, 1)
	i64 = append(i64, 1)
	i64 = append(i64, 1)
	i64 = append(i64, 1)
	i64 = append(i64, 1)
	i64 = append(i64, 1)
	fmt.Println(len(i64), cap(i64))
}
