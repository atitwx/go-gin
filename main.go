package main

import (
	"go-gin/repository"
	"go-gin/routes"

	"go-gin/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	repository.InitDatabase()
	router := gin.Default()
	router.Use(middleware.LoggerMiddleware())
	// 创建两个路由组
	Router := router.Group("/api")
	RouterPub := router.Group("/public")
	routes.RouterGroupApp.SysUserApi.InitSysUserRouter(Router, RouterPub)
	router.Run(":8001")
}
