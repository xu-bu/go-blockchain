package main

import (
	_"context"
	_"errors"
	_"log"
	_"net/http"
	_"os"
	_"os/signal"
	_"syscall"
	_"time"

	"github.com/gin-gonic/gin"
	_"github.com/julienschmidt/httprouter"
)

func main() {
	// 创建一个默认的路由引擎
	app := gin.Default()
	// r.GET("/", test)
	// post 请求在 runApi中要使用body中的x-www
	app.Handle("POST", "/", test)
	// 默认8080端口
    app.Run()
}

// 在这里分发路由,处理函数在handler.go中，同一个包中不同文件的函数可以直接调用，不用import
// func newRouter() *httprouter.Router {
// 	mux := httprouter.New()
// 	mux.GET("/createNFT", createNFT)
// 	mux.POST("/", test)

// 	return mux
// }

// func main() {

// 	srv := &http.Server{
// 		Addr:    ":10101",
// 		Handler: newRouter(),
// 	}

// 	idleConnsClosed := make(chan struct{})
// 	go func() {
// 		sigint := make(chan os.Signal, 1)
// 		signal.Notify(sigint, os.Interrupt)
// 		signal.Notify(sigint, syscall.SIGTERM)
// 		<-sigint

// 		log.Println("service interrupt received")

// 		log.Println("http server shutting down")
// 		time.Sleep(5 * time.Second)

// 		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
// 		defer cancel()

// 		if err := srv.Shutdown(ctx); err != nil {
// 			log.Printf("http server shutdown error: %v", err)
// 		}

// 		log.Println("shutdown complete")

// 		close(idleConnsClosed)

// 	}()

// 	log.Printf("Starting server on port 10101")
// 	if err := srv.ListenAndServe(); err != nil {
// 		if !errors.Is(err, http.ErrServerClosed) {
// 			log.Fatalf("fatal http server failed to start: %v", err)
// 		}
// 	}

// 	<-idleConnsClosed
// 	log.Println("Service Stop")

// }


