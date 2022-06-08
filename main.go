package main

import (
	"busuanzi/process/webutil"
)

func main() {
<<<<<<< HEAD

	if !config.C.Web.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	// middleware
	if config.C.Web.Log {
		r.Use(gin.Logger())
	}
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	// routers
	{
		r.POST("/api", controller.ApiHandler)
		r.GET("/ping", controller.PingHandler)
		r.NoRoute(controller.NoRouteHandler)
	}

	// start server
	log.SetOutput(os.Stdout)
	log.Println("server listen on port:", config.C.Web.Address)
	err := r.Run(config.C.Web.Address)
	if err != nil {
		log.Fatalf("web服务启动失败: %s", err)
	}
=======
	webutil.Init()
>>>>>>> dev
}
