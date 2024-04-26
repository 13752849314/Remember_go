package router

import (
	"github.com/gin-gonic/gin"
	"remember/common"
	"remember/controller"
	"remember/service"
)

func FileRouter(r *gin.Engine) *gin.Engine {
	r.MaxMultipartMemory = 10 << 20 // 10MB

	FileRouterUser := r.Group("/file").Use(service.JwtCheck(common.User))
	FileRouterUser.POST("/upload", controller.UpLoad)
	FileRouterUser.GET("/list", controller.GetFilesList)
	FileRouterUser.GET("/download/:flag", controller.Download)

	return r
}
