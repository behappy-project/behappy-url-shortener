package handler

import (
	"behappy-url-shortener/src/common"
	"behappy-url-shortener/src/model"
	"behappy-url-shortener/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShortenUrl(c *gin.Context) error {
	getShortenR := &model.ShortenRequest{}
	if bindErr := c.ShouldBindJSON(getShortenR); bindErr != nil {
		return common.ParameterError(bindErr.Error())
	}
	err, response := service.ShortenUrl(getShortenR)
	if err != nil {
		return err
	}
	c.JSON(http.StatusOK, common.OK.WithData(response))
	return nil
}

func HandleShortenUrl(c *gin.Context) error {
	shortUrl := c.Param("short_url")
	err, longUrl := service.HandleShortenUrl(shortUrl)
	if err != nil {
		c.HTML(http.StatusOK, "shorten/index.tmpl", gin.H{
			"status": "500",
			"msg":    err.Error(),
		})
		return nil
	}
	c.Redirect(http.StatusMovedPermanently, longUrl)
	return nil
}
