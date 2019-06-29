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

func (u *userClient) Get(ctx context.Context, accountID string) (string, entity.User, error) {
	q := u.gClient.NewQuery(generateEntityByEnv(entity.UserEntityName, u.env)).
		Filter("account_id = ", accountID).
		Limit(1)

	// user entity is only by account_id.
	entities := make([]entity.User, 0, 1)
	keys, err := u.gClient.GetAll(ctx, q, &entities)
	if err != nil {
		return "", entity.User{}, errors.Wrap(err, "failed user get.")
	}
	if len(entities) != 1 {
		return "", entity.User{}, xerrors.Errorf("failed fetching entity id=%v: %w", accountID, repository.NotFoundError{})
	}
	return keys[0].Encode(), entities[0], nil
}

func (u *userClient) Delete(ctx context.Context, encodeKey string) error {
	key, err := u.gClient.DecodeKey(encodeKey)
	if err != nil {
		return xerrors.Errorf("failed decoded key: %v", err)
	}
	if err := u.gClient.Delete(ctx, key); err != nil {
		return xerrors.Errorf("failed delete datastore entity key=%v: %v", encodeKey, err)
	}
	return nil
}
