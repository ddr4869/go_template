package dto

import "time"

type Payment struct {
}

type ReqCreatePayment struct {
	MovieName string `form:"movie_name" json:"movie_name" binding:"required"`
}

type PaymentDto struct {
	ID          int
	MovieName   string
	UserName    string
	Grade       string
	CreatedDate time.Time
}
