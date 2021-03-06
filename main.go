package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-gin-example/models"
	"go-gin-example/pkg/gredis"
	"go-gin-example/pkg/logging"
	"go-gin-example/pkg/setting"
	"go-gin-example/routers"
	"log"
	"net/http"
)

// @title Golang Gin API
// @version 1.0
// @description An example of gin
func main() {
	router := routers.InitRouter()

	setting.Setup()
	models.Setup()
	logging.Setup()
	gredis.SetUp()

	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	s := &http.Server{
		Addr:           endPoint,
		Handler:        router,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	s.ListenAndServe()
}
