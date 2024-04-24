package router

import (
	"github.com/gin-gonic/gin"
	"remember/common"
	"remember/controller"
	"remember/service"
)

func FileRouter(r *gin.Engine) *gin.Engine {
	FileRouterUser := r.Group("/file").Use(service.JwtCheck(common.User))
	FileRouterUser.POST("/upload", controller.UpLoad)
	FileRouterUser.GET("/list", controller.GetFilesList)

	return r
}
