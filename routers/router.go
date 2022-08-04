package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/kbsonlong/gin-wechat-bot/routers/api/v1"
)

func InitRouter() gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode("debug")

	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/webhook", v1.Webhook)
	}
	return *r
}
