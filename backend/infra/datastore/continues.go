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
	k := generateKey(c.gClient, entity.ContinuesEntityName, c.env, e.AccountID)
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
	k := generateKey(c.gClient, entity.ContinuesEntityName, c.env, accountID)
	var e entity.Continues
	if err := c.gClient.Get(ctx, k, &e); err != nil {
		if err == datastore.ErrNoSuchEntity {
			return "", entity.Continues{}, xerrors.Errorf("failed fetching entity id=%v: %w", k.Encode(), repository.NotFoundError{})
		}
		return "", entity.Continues{}, xerrors.Errorf("failed continues entity get: %v", err)
	}
	return k.Encode(), e, nil
}
