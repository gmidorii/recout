package datastore

import (
	"context"

	"github.com/gmidorii/recout/backend/infra/entity"
	"github.com/gmidorii/recout/backend/infra/repository"
	"github.com/pkg/errors"
	"go.mercari.io/datastore"
	"golang.org/x/xerrors"
)

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
	key := generateKey(u.gClient, entity.UserEntityName, u.env, e.AccountID)
	if _, err := u.gClient.Put(ctx, key, &e); err != nil {
		return "", errors.Wrap(err, "failed put recout entity to datastore")
	}
	return key.Encode(), nil
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
	key := generateKey(u.gClient, entity.UserEntityName, u.env, accountID)
	var e entity.User
	if err := u.gClient.Get(ctx, key, &e); err != nil {
		if err == datastore.ErrNoSuchEntity {
			return entity.User{}, xerrors.Errorf("not found entity account_id=%v: %w", accountID, repository.NotFoundError{})
		}
		return entity.User{}, xerrors.Errorf("failed get entity account_id=%v: %w", accountID, err)
	}
	return e, nil
}

func (u *userClient) Delete(ctx context.Context, accountID string) error {
	key := generateKey(u.gClient, entity.UserEntityName, u.env, accountID)
	if err := u.gClient.Delete(ctx, key); err != nil {
		return xerrors.Errorf("failed delete datastore entity account_id=%v: %w", accountID, err)
	}
	return nil
}
