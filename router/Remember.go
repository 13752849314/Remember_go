package router

import (
	"github.com/gin-gonic/gin"
)

func Remember() *gin.Engine {
	r := gin.Default()

	r = UserRouter(r)

	return r
}
