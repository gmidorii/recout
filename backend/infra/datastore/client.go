package datastore

import (
	"context"

	"github.com/gmidorii/recout/backend/infra/entity"
	"github.com/gmidorii/recout/backend/infra/repository"
	"github.com/pkg/errors"
	"go.mercari.io/datastore"
	"golang.org/x/xerrors"
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
	q := c.gClient.
		NewQuery(generateEntityByEnv(entity.RecoutEntityName, c.env)).
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

type userClient struct {
	gClient datastore.Client
	env     string
}

func NewUserClient(gClient datastore.Client, env string) (repository.User, error) {
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
	if _, err := u.gClient.Put(ctx, key, &e); err != nil {
		return "", errors.Wrap(err, "failed put recout entity to datastore")
	}
	return generatedUUID, nil
}

func (u *userClient) Fetch(ctx context.Context, offset int, limit int) ([]entity.User, error) {
	//TODO: use offset query.
	q := u.gClient.NewQuery(generateEntityByEnv(entity.UserEntityName, u.env)).Order("-CreatedAt").Limit(limit)

	entities := make([]entity.User, 0, limit)
	_, err := u.gClient.GetAll(ctx, q, &entities)
	if err != nil {
		return nil, errors.Wrap(err, "failed get all from datastore")
	}
	return entities, nil
}

func (u *userClient) Get(ctx context.Context, accountID string) (entity.User, error) {
	q := u.gClient.NewQuery(generateEntityByEnv(entity.UserEntityName, u.env)).
		Filter("account_id = ", accountID).
		Limit(1)

	// user entity is only by account_id.
	entities := make([]entity.User, 0, 1)
	if _, err := u.gClient.GetAll(ctx, q, &entities); err != nil {
		return entity.User{}, errors.Wrap(err, "failed user get.")
	}
	if len(entities) != 1 {
		return entity.User{}, xerrors.Errorf("failed fetching entity id=%v: %w", accountID, repository.NotFoundError{})
	}
	return entities[0], nil
}

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

func (c *continuesClient) PutKey(ctx context.Context, key string, e entity.Continues) error {
	k := c.gClient.NameKey(generateEntityByEnv(entity.ContinuesEntityName, c.env), key, nil)
	if _, err := c.gClient.Put(ctx, k, &e); err != nil {
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
	return keys[0].Name(), entities[0], nil
}
