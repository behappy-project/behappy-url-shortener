package cmd

import (
	"behappy-url-shortener/src/common"
	"behappy-url-shortener/src/model"
	"behappy-url-shortener/src/routers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

func startApp() {
	g := gin.Default()
	//gin.SetMode(gin.ReleaseMode)
	// cors
	g.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			// 自定义过滤源站的方法
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))
	// 使用嵌套组实现根路径(when need)
	routerGroup := g.Group("/")
	g.NoMethod(common.HandleNotFound)
	g.NoRoute(common.HandleNotFound)
	routers.ShortenRoute(routerGroup)
	// 启动HTTP服务，默认在0.0.0.0:3000启动服务
	addr := ":" + model.RunOpts.Port
	if err := g.Run(addr); err != nil {
		logrus.Error("startup service failed, err: %v\n\n", err)
	}
}
