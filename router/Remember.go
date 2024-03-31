package router

import (
	"github.com/gin-gonic/gin"
	"remember/controller"
)

func Remember() *gin.Engine {
	r := gin.Default()
	// todo 路由注册
	r.GET("/", controller.GetAllUsers)
	return r
}
