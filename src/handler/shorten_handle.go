package handler

import (
	"behappy-url-shortener/src/common"
	"behappy-url-shortener/src/model"
	"github.com/gin-gonic/gin"
)

func GetShortenUrl(c *gin.Context) error {
	getShortenR := &model.GetShortenRequest{}
	if bindErr := c.ShouldBindJSON(getShortenR); bindErr != nil {
		return common.ParameterError(bindErr.Error())
	}
	return nil
}

func HandleShortenUrl(c *gin.Context) error {
	return nil
}
