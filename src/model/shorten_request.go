package model

import "time"

type GetShortenRequest struct {
	LongUrl   string    `json:"long_url" binding:"required"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	CNew      bool      `json:"c_new"`
}
