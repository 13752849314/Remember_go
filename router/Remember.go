package router

import (
	"github.com/gin-gonic/gin"
	"remember/config"
)

func Remember() *gin.Engine {
	r := gin.Default()
	r.Use(config.Cors())

	r = UserRouter(r)
	r = BillRouter(r)
	r = FileRouter(r)

	return r
}
