package datastore

import (
	"context"

	"github.com/gmidorii/recout/backend/infra/entity"
	"github.com/gmidorii/recout/backend/infra/repository"
	"go.mercari.io/datastore"
	"golang.org/x/xerrors"
)

type continuesClient struct {
	gClient datastore.Client
	env     string
}

func NewContinuesClient(gClient datastore.Client, env string) repository.Continues {
	return &continuesClient{
		gClient: gClient,
		env:     env,
	}
}

func (c *continuesClient) Put(ctx context.Context, e entity.Continues) error {
	k := c.gClient.IncompleteKey(generateEntityByEnv(entity.ContinuesEntityName, c.env), nil)
	if _, err := c.gClient.Put(ctx, k, &e); err != nil {
		return err
	}
	return nil
}

func (c *continuesClient) PutKey(ctx context.Context, encodedKey string, e entity.Continues) error {
	key, err := c.gClient.DecodeKey(encodedKey)
	if err != nil {
		return xerrors.Errorf("failed decode key: %v", err)
	}
	if _, err := c.gClient.Put(ctx, key, &e); err != nil {
		return err
	}
	return nil
}

func (c *continuesClient) Get(ctx context.Context, accountID string) (string, entity.Continues, error) {
	q := c.gClient.NewQuery(generateEntityByEnv(entity.ContinuesEntityName, c.env)).
		Filter("account_id = ", accountID).
		Limit(1)

	// user entity is only by account_id.
	entities := make([]entity.Continues, 0, 1)
	keys, err := c.gClient.GetAll(ctx, q, &entities)
	if err != nil {
		return "", entity.Continues{}, xerrors.Errorf("failed continues entity get: %v", err)
	}
	if len(entities) == 0 {
		return "", entity.Continues{}, repository.NotFoundError{}
	}
	return keys[0].Encode(), entities[0], nil
}
