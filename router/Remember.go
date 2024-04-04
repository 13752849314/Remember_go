package router

import (
	"github.com/gin-gonic/gin"
	"remember/common"
	"remember/controller"
	"remember/service"
)

func Remember() *gin.Engine {
	r := gin.Default()
	// todo 路由注册
	userRouterAdmins := r.Group("/user").Use(service.JwtCheck(common.Admins))
	userRouterAdmins.GET("/users", controller.GetAllUsers)

	userRouterAdmin := r.Group("/user").Use(service.JwtCheck(common.Admin))
	userRouterAdmin.GET("/user", controller.GetAllUser)

	userRouterUser := r.Group("/user")
	userRouterUser.POST("/register", controller.Registration)
	userRouterUser.POST("/login", controller.Login)
	return r
}
