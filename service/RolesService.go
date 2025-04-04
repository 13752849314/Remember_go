package service

import (
	"github.com/gin-gonic/gin"
	"log"
	"remember/common"
	"remember/mapper"
	"remember/utils"
)

var JWTWHITELIST []string

func init() {
	JWTWHITELIST = make([]string, 0)
	JWTWHITELIST = append(JWTWHITELIST, "/user/login")
	JWTWHITELIST = append(JWTWHITELIST, "/user/register")
}

func IsInJwtWhiteList(path string) bool {
	for _, s := range JWTWHITELIST {
		if s == path {
			return true
		}
	}
	return false
}

func JwtCheck(role common.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("JwtCheck")
		s := c.Request.URL.String()
		if IsInJwtWhiteList(s) {
			c.Next()
			return
		}
		jjwt := c.Request.Header.Get("jwt")
		if jjwt == "" {
			c.JSON(200, common.StatusErr().SetMessage("请先登录"))
			c.Abort()
			return
		}
		claims, err := utils.CheckToken(jjwt)
		if err != nil {
			c.JSON(200, common.StatusErr().SetMessage(err.Error()))
			c.Abort()
			return
		}
		username := claims.Username
		info := utils.GetInfo(username)
		if info != jjwt {
			c.JSON(200, common.StatusErr().SetMessage("用户未登录或token过期"))
			c.Abort()
			return
		}
		if !RolesCheck(role, username, c) {
			c.JSON(200, common.StatusErr().SetMessage("权限不够"))
			c.Abort()
			return
		}
		c.Next()
	}
}

func RolesCheck(role common.Role, username string, c *gin.Context) bool {
	log.Println("RolesCheck")
	um := mapper.UserMapper{}
	user := um.GetUserByUsername(username)
	c.Set("user", user)
	return role.Ge(common.ValueOf(user.Roles))
}
