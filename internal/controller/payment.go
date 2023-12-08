package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-board/internal/dto"
	"github.com/go-board/internal/service"
)

type Payment struct {
	payment service.Payment
}

func NewPayment(payment service.Payment) *Payment {
	return &Payment{payment: payment}
}

func (s *Payment) CreatePayment(c *gin.Context) {
	req := c.MustGet("req").(dto.ReqCreatePayment)
	token := c.MustGet("token").(*dto.AccessClaims)

	payment, err := s.payment.CreatePayment(c, req.MovieName, token.Username, token.UserGrade)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, payment)
}
