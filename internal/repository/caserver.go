package repository

import (
	"context"

	"github.com/kafka-go/ent"
	"github.com/kafka-go/ent/user"
	"github.com/the-medium-tech/platform-externals/log"
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
		SetUserID(10).
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
		Where(user.Name(name)).
		QueryCaserver()

	list, err := query.All(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return list, nil
}
