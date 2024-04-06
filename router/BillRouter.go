package router

import (
	"github.com/gin-gonic/gin"
	"remember/common"
	"remember/controller"
	"remember/service"
)

func BillRouter(r *gin.Engine) *gin.Engine {
	BillRouterUser := r.Group("/bill").Use(service.JwtCheck(common.User))
	BillRouterUser.POST("/add", controller.AddBill)

	return r
}
