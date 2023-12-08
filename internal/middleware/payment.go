package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-board/internal/controller"
	"github.com/go-board/internal/dto"
	"github.com/go-board/log"
)

type Payment struct {
	payment controller.Payment
}

func NewPayment(payment controller.Payment) *Payment {
	return &Payment{payment: payment}
}

func (s *Payment) CreatePayment(c *gin.Context) {
	var req dto.ReqCreatePayment
	if err := c.ShouldBind(&req); err != nil {
		log.Info(err)
		MwAbortJson(c, http.StatusBadRequest, err)
		return
	}
	c.Set("req", req)
	c.Next()
}
