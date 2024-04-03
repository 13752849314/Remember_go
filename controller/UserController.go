package controller

import (
	"github.com/gin-gonic/gin"
	"remember/common"
	"remember/entity"
	"remember/service/impl"
)

func GetAllUsers(c *gin.Context) {
	us := impl.UserServiceImpl{}
	users := us.GetAllUsers()
	c.JSON(200, common.StatusOk().AddData("users", users))
}

func GetAllUser(c *gin.Context) {
	us := impl.UserServiceImpl{}
	users := us.GetAllUser()
	c.JSON(200, common.StatusOk().AddData("users", users))
}

func Registration(c *gin.Context) {
	us := impl.UserServiceImpl{}
	user := new(entity.User)
	err := c.ShouldBindJSON(user)
	if err != nil {
		c.JSON(200, common.StatusErr().SetMessage(err.Error()))
		return
	}
	err = us.Registration(user)
	if err != nil {
		c.JSON(200, common.StatusErr().SetMessage(err.Error()))
		return
	}
	c.JSON(200, common.StatusOk().SetMessage("注册成功!"))
}

func Login(c *gin.Context) {
	us := impl.UserServiceImpl{}
	user := new(entity.User)
	err := c.ShouldBindJSON(user)
	if err != nil {
		c.JSON(200, common.StatusErr().SetMessage(err.Error()))
		return
	}
	var token string
	token, err = us.Login(user)
	if err != nil {
		c.JSON(200, common.StatusErr().SetMessage(err.Error()))
		return
	}
	c.JSON(200, common.StatusOk().SetMessage("登录成功").AddData("jwt", token))
}
