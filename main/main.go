package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"hydra/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Logo
	fmt.Printf("Starting Hydra Server ...\n")
	fmt.Printf("Version 1.0.0 \n")

	// 参数
	port := flag.String("port", "9527", "port to listen on.")
	flag.Parse()

	// 初始化 Gin
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	// Print the routes
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("%-7s %s", httpMethod, absolutePath)
	}

	// 注册路由
	routers.Register(router)

	// 启动 Server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", *port),
		Handler: router,
	}
	go func() {
		fmt.Printf("🦊️ Listen on 0.0.0.0:%s\n", *port)
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err.Error())
		}
	}()

	// 等待退出信号
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 关闭服务器
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}
