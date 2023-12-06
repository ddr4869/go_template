package repository

import (
	"context"

	"github.com/go-board/ent"
	"github.com/go-board/ent/user"
)

type User struct {
	client *ent.Client
}

func NewUser(client *ent.Client) *User {
	return &User{client: client}
}

func (r *User) UserList(ctx context.Context) ([]*ent.User, error) {

	query := r.client.User.Query()
	list, err := query.All(ctx)
	if err != nil {
		//tmp log.Error(err)
		return nil, err
	}
	return list, nil
}

func (r *User) GetUser(ctx context.Context, name string) (*ent.User, error) {

	query := r.client.User.Query().Where(user.Name(name))
	list, err := query.Only(ctx)
	if err != nil {
		//tmp log.Error(err)
		return nil, err
	}
	return list, nil
}

func (r *User) CreateUser(ctx context.Context, name, description string, password []byte) (*ent.User, error) {

	user, err := r.client.User.Create().
		SetName(name).
		SetDescription(description).
		SetPassword(password).
		Save(ctx)

	if err != nil {
		//tmp log.Error(err)
		return nil, err
	}
	return user, nil
}

func (r *User) UserLogin(ctx context.Context) ([]*ent.User, error) {

	query := r.client.User.Query()
	list, err := query.All(ctx)
	if err != nil {
		//tmp log.Error(err)
		return nil, err
	}
	return list, nil
}
