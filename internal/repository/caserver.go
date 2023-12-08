package repository

import (
	"context"

	"github.com/go-board/ent"
	"github.com/go-board/ent/user"
	"github.com/go-board/log"
)

type CaServer struct {
	client *ent.Client
}

func NewCaServer(client *ent.Client) *CaServer {
	return &CaServer{client: client}
}

func (r *CaServer) CaServerList(ctx context.Context) ([]*ent.CaServer, error) {

	query := r.client.CaServer.Query()
	list, err := query.All(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return list, nil
}

func (r *CaServer) CreateCaServer(ctx context.Context, name string) (*ent.CaServer, error) {

	user, err := r.client.CaServer.Create().
		SetName(name).
		SetUserID(name).
		//SetUser("qwe").
		Save(ctx)

	if err != nil {
		log.Error(err)
		return nil, err
	}
	return user, nil
}

func (r *CaServer) UsersCaServer(ctx context.Context, name string) ([]*ent.CaServer, error) {

	query := r.client.User.Query().
		Where(user.ID(name)).
		QueryCaserver()

	list, err := query.All(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return list, nil
}
