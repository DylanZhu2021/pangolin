package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"pangolin/core"
	"pangolin/web/controller"
	"pangolin/web/router"
	"syscall"
	"time"
)

func main() {

	e := new(core.Engine)
	e.Init()

	core.CoreEngine = e

	//初始化服务
	controller.NewServices()

	// 注册路由
	r := router.SetupRouter()
	// 启动服务
	srv := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: r,
	}
	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Println("listen:", err)
		}
	}()

	// 优雅关机
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Server Shutdown:", err)
	}

	log.Println("Server exiting")

}
