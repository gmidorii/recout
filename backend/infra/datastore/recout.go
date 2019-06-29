package datastore

import (
	"context"

	"github.com/gmidorii/recout/backend/infra/entity"
	"github.com/gmidorii/recout/backend/infra/repository"
	"github.com/pkg/errors"
	"go.mercari.io/datastore"
)

type recoutClient struct {
	gClient datastore.Client
	env     string
}

func NewRecoutClient(gClient datastore.Client, env string) (repository.Recout, error) {
	return &recoutClient{
		gClient: gClient,
		env:     env,
	}, nil
}

func (c *recoutClient) Put(ctx context.Context, e entity.Recout) (string, error) {
	generatedUUID, err := generateUUID()
	if err != nil {
		return "", err
	}
	key := generateKey(c.gClient, entity.RecoutEntityName, c.env, generatedUUID)
	if _, err := c.gClient.Put(ctx, key, &e); err != nil {
		return "", errors.Wrap(err, "failed put recout entity to datastore")
	}
	return generatedUUID, nil
}

func (c *recoutClient) Fetch(ctx context.Context, accountID string, offset int, limit int) ([]entity.Recout, error) {
	//TODO: use offset query.
	q := newQuery(c.gClient, entity.RecoutEntityName, c.env).
		Filter("account_id =", accountID).
		Order("-created_at").
		Limit(limit)

	entities := make([]entity.Recout, 0, limit)
	_, err := c.gClient.GetAll(ctx, q, &entities)
	if err != nil {
		return nil, errors.Wrap(err, "failed get all from datastore")
	}
	return entities, nil
}
