package router

import (
	"github.com/gin-gonic/gin"
	"remember/controller"
)

func Remember() *gin.Engine {
	r := gin.Default()
	// todo 路由注册
	r.GET("/", controller.GetAllUsers)
	r.POST("/register", controller.Registration)
	r.POST("/login", controller.Login)
	return r
}
