package repository

import (
	"context"

	"github.com/kafka-go/ent"
	"github.com/the-medium-tech/platform-externals/log"
)

type Admin struct {
	client *ent.Client
}

func NewAdmin(client *ent.Client) *Admin {
	return &Admin{client: client}
}

func (r *Admin) AdminList(ctx context.Context) ([]*ent.Admin, error) {

	query := r.client.Admin.Query()
	list, err := query.All(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return list, nil
}

func (r *Admin) CreateAdmin(ctx context.Context, name string) (*ent.Admin, error) {

	Admin, err := r.client.Admin.Create().
		SetName(name).
		Save(ctx)

	if err != nil {
		log.Error(err)
		return nil, err
	}
	return Admin, nil
}
