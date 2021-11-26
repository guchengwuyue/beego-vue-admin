package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
	"yixiang.co/go-mall/app/models"
	"yixiang.co/go-mall/pkg/base"
	"yixiang.co/go-mall/pkg/global"
	"yixiang.co/go-mall/pkg/jwt"
	"yixiang.co/go-mall/pkg/logging"
	"yixiang.co/go-mall/pkg/redis"
	"yixiang.co/go-mall/routers"
)

func init() {
	global.YSHOP_VP = base.Viper()
	models.Setup()
	logging.Setup()
	redis.Setup()
	jwt.Setup()
}

// @title gin-shop  API
// @version 1.0
// @description gin-shop商城后台管理系统
// @termsOfService https://gitee.com/guchengwuyue/gin-shop
// @license.name apache2
func main() {
	gin.SetMode(global.YSHOP_CONFIG.Server.RunMode)

	routersInit := routers.InitRouter()
	endPoint := fmt.Sprintf(":%d", global.YSHOP_CONFIG.Server.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: maxHeaderBytes,
	}



	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()

}
