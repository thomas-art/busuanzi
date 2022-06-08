package webutil

import (
	"busuanzi/app/controller"
	"github.com/gin-gonic/gin"
)

func initRoute(r *gin.Engine) {
	{
		r.POST("/api", controller.ApiHandler)
		r.GET("/ping", controller.PingHandler)
	}
}
