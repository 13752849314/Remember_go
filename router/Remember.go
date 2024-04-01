package router

import (
	"github.com/gin-gonic/gin"
	"remember/controller"
)

func Remember() *gin.Engine {
	r := gin.Default()
	// todo 路由注册
	userR := r.Group("/user")
	userR.GET("/users", controller.GetAllUsers)
	userR.GET("/user", controller.GetAllUser)
	userR.POST("/register", controller.Registration)
	userR.POST("/login", controller.Login)
	return r
}
