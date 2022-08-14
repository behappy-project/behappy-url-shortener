package common

import (
	"github.com/gin-gonic/gin"
)

// APIException api错误的结构体
type APIException struct {
	*Response
}

// 实现接口
func (e *APIException) Error() string {
	return e.Msg
}

func newAPIException(code int, msg string) *APIException {
	return &APIException{
		&Response{
			Code: code,
			Msg:  msg,
		},
	}
}

type HandlerFunc func(c *gin.Context) error
