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
				apiException = ServerErrorWithMsg(e.Error())
			} else {
				apiException = UnknownError()
			}
			c.JSON(http.StatusOK, apiException)
			return
		}
	}
}

// ServerErrorWithMsg 500 错误处理
func ServerErrorWithMsg(msg string) *APIException {
	return newAPIException(http.StatusInternalServerError, msg)
}

// NotFound 404 错误
func NotFound() *APIException {
	return newAPIException(http.StatusNotFound, http.StatusText(http.StatusNotFound))
}

// NotFoundWithMsg 404 错误
func NotFoundWithMsg(msg string) *APIException {
	return newAPIException(http.StatusNotFound, msg)
}

// UnknownError 未知错误
func UnknownError() *APIException {
	return newAPIException(http.StatusForbidden, "UnknownError")
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
