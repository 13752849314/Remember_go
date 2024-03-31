package controller

import (
	"github.com/gin-gonic/gin"
	"remember/common"
	"remember/service/impl"
)

func GetAllUsers(c *gin.Context) {
	us := impl.UserServiceImpl{}
	users := us.GetAllUsers()
	c.JSON(200, common.StatusOk().AddData("users", users))
}
