package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"go-gin-backend-admin/model"
	"go-gin-backend-admin/pkg/logging"
	"go-gin-backend-admin/pkg/setting"
	"go-gin-backend-admin/pkg/util"
	"go-gin-backend-admin/router"
)

func init() {
	setting.Setup()
	model.Setup()
	logging.Setup()
	util.Setup()
}

// @title Golang Gin API
// @version 1.0
// @description An example of gin Admin
// @termsOfService https://github.com/liuvigongzuoshi/go-gin-backend-admin
// @license.name MIT
// @license.url https://github.com/liuvigongzuoshi/go-gin-backend-admin/blob/master/LICENSE
// @host localhost:8000
// @host api.ywbang.top
func main() {
	gin.SetMode(setting.ServerSetting.AppMode)

	routerInit := router.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	httpPort := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	server := &http.Server{
		Addr:         httpPort,
		Handler:      routerInit,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	log.Printf("[info] start http server listening %s", setting.AppSetting.PrefixUrl)

	server.ListenAndServe()
}
