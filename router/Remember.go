package router

import (
	"github.com/gin-gonic/gin"
	"remember/controller"
	"remember/service"
)

func Remember() *gin.Engine {
	r := gin.Default()
	// todo 路由注册
	userR := r.Group("/user").Use(service.JwtCheck())
	userR.GET("/users", controller.GetAllUsers)
	userR.GET("/user", controller.GetAllUser)
	userP := r.Group("/user")
	userP.POST("/register", controller.Registration)
	userP.POST("/login", controller.Login)
	return r
}
