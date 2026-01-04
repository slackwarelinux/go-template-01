package routers

import (
	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine) {
	r := router.Group("/api")
	{
		r.GET("/", index)
		r.GET("/ping", ping)
	}
}
