package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/TskFok/GinApi/app/global"
	"github.com/TskFok/GinApi/app/process"
	"github.com/TskFok/GinApi/bootstrap"
	"github.com/TskFok/GinApi/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//守护进程
	args := os.Args

	if len(args) != 1 && args[1] == "bg" {
		process.InitProcess()
	}

	bootstrap.Init()

	router.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", global.AppHttpPort),
		Handler:        router.Handle,
		ReadTimeout:    time.Duration(global.AppReadTimeOut) * time.Second,
		WriteTimeout:   time.Duration(global.AppWriteTimeOut) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
