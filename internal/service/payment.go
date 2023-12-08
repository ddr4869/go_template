package service

import (
	"github.com/gin-gonic/gin"
	"github.com/go-board/ent"
	"github.com/go-board/internal/dto"
	"github.com/go-board/internal/repository"
)

type Payment struct {
	payment repository.Payment
}

func NewPayment(payment repository.Payment) *Payment {
	return &Payment{payment: payment}
}

func (s *Payment) CreatePayment(c *gin.Context, movieName, userName, grade string) (*dto.PaymentDto, error) {

	payment, err := s.payment.CreatePayment(c, movieName, userName, grade)
	if err != nil {
		return nil, err
	}
	return convertPaymentEntToDto(payment), nil
}

func convertPaymentEntToDto(entPayment *ent.Payment) *dto.PaymentDto {

	dto_payment := dto.PaymentDto{
		ID:          entPayment.ID,
		MovieName:   entPayment.MovieID,
		UserName:    entPayment.UserName,
		Grade:       entPayment.Grade,
		CreatedDate: entPayment.CreatedDate,
	}

	return &dto_payment
}
