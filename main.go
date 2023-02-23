package main

import (
	"fmt"
	"github.com/TskFok/GinApi/app/router"
	"github.com/TskFok/GinApi/app/utils/conf"
	"net/http"
	"time"
)

func main() {
	conf.InitConf()

	handler := router.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", conf.AppHttpPort),
		Handler:        handler,
		ReadTimeout:    time.Duration(conf.AppReadTimeOut) * time.Second,
		WriteTimeout:   time.Duration(conf.AppWriteTimeOut) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
