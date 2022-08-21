package model

import (
	"github.com/golang-module/carbon/v2"
)

type ShortenRequest struct {
	LongUrl   string          `json:"long_url" binding:"required"`
	StartDate carbon.DateTime `json:"start_date"`
	EndDate   carbon.DateTime `json:"end_date"`
	// 如果为true,则将覆盖已有url,生成新的
	CNew bool `json:"c_new"`
}

type ShortenResponse struct {
	LongUrl string `json:"long_url" mapstructure:"long_url"`
	Hash    string `json:"hash" mapstructure:"hash"`
}
