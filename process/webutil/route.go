package webutil

import (
	"busuanzi/app/controller"
	"busuanzi/app/middleware"
	"github.com/gin-gonic/gin"
)

func initRoute(r *gin.Engine) {
	{
		r.POST("/api", middleware.Identity(), controller.ApiHandler)
		r.GET("/ping", controller.PingHandler)
	}
}
