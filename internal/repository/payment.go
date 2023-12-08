package repository

import (
	"context"

	"github.com/go-board/ent"
	"github.com/go-board/log"
)

type Payment struct {
	client *ent.Client
}

func NewPayment(client *ent.Client) *Payment {
	return &Payment{client: client}
}

func (r *Payment) CreatePayment(ctx context.Context, movieName, userName, grade string) (*ent.Payment, error) {

	payment, err := r.client.Payment.Create().
		SetUserName(userName).
		SetUserID(userName).
		SetGrade(grade).
		SetMovieID(movieName).
		Save(ctx)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	return payment, nil
}
