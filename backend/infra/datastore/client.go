package datastore

import (
	"context"
	"fmt"

	"github.com/gmidorii/recout/backend/infra/entity"
	"github.com/gmidorii/recout/backend/infra/repository"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"go.mercari.io/datastore"
	"go.mercari.io/datastore/aedatastore"
)

type recoutClient struct {
	gClient datastore.Client
	env     string
}

func NewRecoutClient(ctx context.Context, env string) (repository.Recout, error) {
	gClient, err := aedatastore.FromContext(ctx)
	if err != nil {
		return &recoutClient{}, err
	}
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
	if _, err := c.gClient.Put(ctx, key, e); err != nil {
		return "", errors.Wrap(err, "failed put recout entity to datastore")
	}
	return generatedUUID, nil
}

func (c *recoutClient) Fetch(ctx context.Context, offset int, limit int) ([]entity.Recout, error) {
	//TODO: use offset query.
	q := c.gClient.NewQuery(generateEntityByEnv(entity.RecoutEntityName, c.env)).Order("-CreatedAt").Limit(limit)

	entities := make([]entity.Recout, 0, limit)
	_, err := c.gClient.GetAll(ctx, q, &entities)
	if err != nil {
		return nil, errors.Wrap(err, "failed get all from datastore")
	}
	return entities, nil
}

type userClient struct {
	gClient datastore.Client
	env     string
}

func NewUserClient(ctx context.Context, env string) (repository.User, error) {
	gClient, err := aedatastore.FromContext(ctx)
	if err != nil {
		return &userClient{}, err
	}
	return &userClient{
		gClient: gClient,
		env:     env,
	}, nil
}

func (u *userClient) Put(ctx context.Context, e entity.User) (string, error) {
	generatedUUID, err := generateUUID()
	if err != nil {
		return "", err
	}
	key := generateKey(u.gClient, entity.UserEntityName, u.env, generatedUUID)
	if _, err := u.gClient.Put(ctx, key, e); err != nil {
		return "", errors.Wrap(err, "failed put recout entity to datastore")
	}
	return generatedUUID, nil
}

func (u *userClient) Fetch(ctx context.Context, offset int, limit int) ([]entity.User, error) {
	//TODO: use offset query.
	q := u.gClient.NewQuery(generateEntityByEnv(entity.RecoutEntityName, u.env)).Order("-CreatedAt").Limit(limit)

	entities := make([]entity.User, 0, limit)
	_, err := u.gClient.GetAll(ctx, q, &entities)
	if err != nil {
		return nil, errors.Wrap(err, "failed get all from datastore")
	}
	return entities, nil
}

func generateUUID() (string, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return "", errors.Wrap(err, "failed generate uuid")
	}
	return id.String(), nil
}

func generateEntityByEnv(kind, env string) string {
	return fmt.Sprintf("%v_%v", env, kind)
}

func generateKey(client datastore.Client, kind, env, uid string) datastore.Key {
	return client.NameKey(generateEntityByEnv(kind, env), uid, nil)
}
