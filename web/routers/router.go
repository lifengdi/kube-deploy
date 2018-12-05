package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv := r.Group("/api/v1")

	{
		apiv.POST("/hello")
		apiv.GET("/hello")
	}

	return r
}
