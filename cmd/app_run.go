package cmd

import (
	"behappy-url-shortener/src/common"
	"behappy-url-shortener/src/model"
	"behappy-url-shortener/src/routers"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil/strutil"
	"github.com/sirupsen/logrus"
	"html/template"
	"time"
)

func startApp() {
	gin.SetMode(gin.ReleaseMode)
	g := gin.Default()
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
	//加载静态文件
	g.Static("/static", "templates/static")
	// Golang模板文件中使用自定义函数方法,需要在LoadHTMLGlob之前
	g.SetFuncMap(template.FuncMap{
		"getStatus": func(arg string) []string {
			args := make([]string, 0)
			for _, val := range arg {
				args = append(args, string(val))
			}
			return args
		},
		"indexPlus": func(index int) string {
			indexPlus, _ := strutil.ToString(index + 1)
			return indexPlus
		},
	})
	//模板解析
	g.LoadHTMLGlob("templates/**/*")
	routers.ShortenRoute(routerGroup)
	// 注册pprof相关路由
	pprof.Register(g)
	// 启动HTTP服务，默认在0.0.0.0:3000启动服务
	addr := ":" + model.RunOpts.Port
	if err := g.Run(addr); err != nil {
		logrus.Error("startup service failed, err: %v\n\n", err)
	}
}
