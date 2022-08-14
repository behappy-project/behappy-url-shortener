package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Wrapper 统一异常处理
func Wrapper(handler HandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		var (
			err error
		)
		err = handler(c)
		if err != nil {
			var apiException *APIException
			// APIException异常
			if h, ok := err.(*APIException); ok {
				apiException = h
			} else if e, ok := err.(error); ok {
				if gin.Mode() == "debug" {
					// 错误
					apiException = ServerErrorWithMsg(e.Error())
				} else {
					// 未知错误
					apiException = UnknownError(e.Error())
				}
			} else {
				apiException = ServerError()
			}
			c.JSON(http.StatusOK, apiException)
			return
		}
	}
}

// ServerError 500 错误处理
func ServerError() *APIException {
	return newAPIException(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
}

// ServerErrorWithMsg 500 错误处理
func ServerErrorWithMsg(msg string) *APIException {
	return newAPIException(http.StatusInternalServerError, msg)
}

// NotFound 404 错误
func NotFound() *APIException {
	return newAPIException(http.StatusNotFound, http.StatusText(http.StatusNotFound))
}

// UnknownError 未知错误
func UnknownError(message string) *APIException {
	return newAPIException(http.StatusForbidden, message)
}

// ParameterError 参数错误
func ParameterError(message string) *APIException {
	return newAPIException(http.StatusBadRequest, message)
}

// HandleNotFound method/request not found处理
func HandleNotFound(c *gin.Context) {
	handleErr := NotFound()
	c.JSON(handleErr.Code, handleErr)
}
