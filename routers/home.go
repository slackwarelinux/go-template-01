package routers

import (
	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	c.JSON(200, gin.H{
		"err": 0,
		"msg": "OK",
	})
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"err": 0,
		"msg": "pong",
	})
}
