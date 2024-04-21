package router

import (
	"github.com/gin-gonic/gin"
	"remember/common"
	"remember/controller"
	"remember/service"
)

func UserRouter(r *gin.Engine) *gin.Engine {
	userRouterAdmins := r.Group("/user").Use(service.JwtCheck(common.Admins))
	userRouterAdmins.GET("/users", controller.GetAllUsers)

	userRouterAdmin := r.Group("/user").Use(service.JwtCheck(common.Admin))
	userRouterAdmin.GET("/user", controller.GetAllUser)
	userRouterAdmin.POST("/add", controller.AddUser)

	userRouterUser := r.Group("/user").Use(service.JwtCheck(common.User))
	userRouterUser.POST("/logout", controller.Logout)
	userRouterUser.POST("/delete", controller.Delete)
	userRouterUser.POST("/changeP", controller.ChangePassword)
	userRouterUser.GET("/info", controller.GetUserInfo)
	userRouterUser.POST("/changeI", controller.ChangeUserInfo)

	userRouter := r.Group("/user")
	userRouter.POST("/register", controller.Registration)
	userRouter.POST("/login", controller.Login)

	return r
}
