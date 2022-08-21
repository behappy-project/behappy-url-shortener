package routers

import (
	"behappy-url-shortener/src/common"
	"behappy-url-shortener/src/handler"
	"github.com/gin-gonic/gin"
)

func ShortenRoute(e *gin.RouterGroup) {
	// 通过long_url获取shorten_url
	e.POST("shorten", common.Wrapper(handler.ShortenUrl))
	// 使用shorten_url进行地址跳转
	e.GET(":short_url", common.Wrapper(handler.HandleShortenUrl))
}
