package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"remember/common"
	"remember/entity"
	"remember/service/impl"
	"remember/utils"
)

var us *impl.UserServiceImpl

func init() {
	us = new(impl.UserServiceImpl)
}

func GetAllUsers(c *gin.Context) {
	users := us.GetAllUsers()
	c.JSON(200, common.StatusOk().AddData("users", users))
}

func GetAllUser(c *gin.Context) {
	users := us.GetAllUser()
	c.JSON(200, common.StatusOk().AddData("user", users))
}

func Registration(c *gin.Context) {
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

func Logout(c *gin.Context) {
	jjwt := c.Request.Header.Get("jwt")
	claims, _ := utils.CheckToken(jjwt)
	err := us.Logout(claims.Username, jjwt)
	if err != nil {
		c.JSON(200, common.StatusErr().SetMessage(err.Error()))
		return
	}
	c.JSON(200, common.StatusOk().SetMessage("退出成功"))
}

func Delete(c *gin.Context) {
	jjwt := c.Request.Header.Get("jwt")
	claims, _ := utils.CheckToken(jjwt)
	controllerUser := claims.Username
	body := make(map[string]interface{})
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(200, common.StatusErr().SetMessage(err.Error()))
		return
	}
	username := body["username"].(string)
	if username == "" {
		c.JSON(200, common.StatusErr().SetMessage("未输入要执行目标用户名"))
		return
	}
	err = us.Delete(controllerUser, username)
	if err != nil {
		c.JSON(200, common.StatusErr().SetMessage(err.Error()))
		return
	}
	c.JSON(200, common.StatusOk().SetMessage("用户："+controllerUser+"成功删除用户："+username))
}

func ChangePassword(c *gin.Context) {
	cp := new(common.ChangeUserP)
	err := c.ShouldBindJSON(cp)
	if err != nil {
		c.JSON(200, common.StatusErr().SetMessage(err.Error()))
		return
	}
	user, _ := c.Get("user")
	jjwt := c.Request.Header.Get("jwt")
	err = us.ChangePassword(user.(*entity.User), jjwt, cp.OldPassword, cp.NewPassword)
	if err != nil {
		c.JSON(200, common.StatusErr().SetMessage(err.Error()))
		return
	}
	c.JSON(200, common.StatusOk().SetMessage("密码修改成功，请重新登录"))
}

func GetUserInfo(c *gin.Context) {
	user, _ := c.Get("user")
	userInfo := us.GetUserInfo(user.(*entity.User))
	c.JSON(200, common.StatusOk().SetMessage("获取成功").AddData("info", userInfo))
}

func ChangeUserInfo(c *gin.Context) {
	ci := new(common.ChangeUserI)
	err := c.ShouldBindJSON(ci)
	if err != nil {
		c.JSON(200, common.StatusErr().SetMessage(err.Error()))
		return
	}
	user, _ := c.Get("user")
	mp := utils.Struct2Map(ci)
	err = us.ChangeUserInfo(user.(*entity.User), mp)
	if err != nil {
		c.JSON(200, common.StatusErr().SetMessage(err.Error()))
		return
	}
	log.Println("修改信息为：", mp)
	c.JSON(200, common.StatusOk().SetMessage("信息修改成功"))
}
