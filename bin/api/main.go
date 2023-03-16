package main

import (
	"fmt"
	"github.com/TskFok/GinApi/app/global"
	"github.com/TskFok/GinApi/bootstrap"
	"github.com/TskFok/GinApi/router"
	"net/http"
	"time"
)

func main() {
	bootstrap.Init()

	router.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", global.AppHttpPort),
		Handler:        router.Handle,
		ReadTimeout:    time.Duration(global.AppReadTimeOut) * time.Second,
		WriteTimeout:   time.Duration(global.AppWriteTimeOut) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
