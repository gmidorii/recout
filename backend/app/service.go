package main

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"go.mercari.io/datastore"
)

const pixelaURL = "https://pixe.la/v1/users"

type Container struct {
	Env      string
	Location *time.Location
	Client   datastore.Client
}

type RecoutService interface {
	Create(ctx context.Context, form RecoutForm) (string, error)
	Fetch(ctx context.Context, form FetchForm) ([]RecoutResponse, error)
}

type recoutService struct {
	Ctn Container
}

func NewRecoutService(ctn Container) RecoutService {
	return &recoutService{Ctn: ctn}
}

func (r *recoutService) Create(ctx context.Context, form RecoutForm) (uid string, err error) {
	client := r.Ctn.Client

	id, err := uuid.NewV4()
	if err != nil {
		return "", errors.Wrap(err, "failed generate uuid")
	}
	uid = id.String()

	key := generateKey(client, recoutEntityName, r.Ctn.Env, uid)
	entity := RecoutEntity{
		AccountID: "gmidorii", //TODO: fix to user login account id
		Message:   form.Message,
		CreatedAt: time.Now().In(r.Ctn.Location).Unix(),
	}
	_, err = client.Put(ctx, key, &entity)
	if err != nil {
		return "", errors.Wrap(err, "failed put datastore")
	}
	return
}

func (r *recoutService) Fetch(ctx context.Context, form FetchForm) ([]RecoutResponse, error) {
	client := r.Ctn.Client

	envEntity := generateEntityByEnv(recoutEntityName, r.Ctn.Env)
	q := client.NewQuery(envEntity).Order("-CreatedAt").Limit(form.Limit)

	entities := make([]RecoutEntity, 0, form.Limit)
	_, err := client.GetAll(ctx, q, &entities)
	if err != nil {
		return nil, errors.Wrap(err, "failed get all from datastore")
	}

	responses := make([]RecoutResponse, len(entities))
	for i, e := range entities {
		responses[i] = RecoutResponse{
			Message:   e.Message,
			CreatedAt: JSONTime(time.Unix(e.CreatedAt, 0).In(r.Ctn.Location)),
		}
	}
	return responses, nil
}

func generateEntityByEnv(kind, env string) string {
	return fmt.Sprintf("%v_%v", env, kind)
}

func generateKey(client datastore.Client, kind, env, uid string) datastore.Key {
	return client.NameKey(generateEntityByEnv(kind, env), uid, nil)
}

type User interface {
	Save(ctx context.Context, form UserForm) error
}

type user struct {
	Ctn Container
}

func NewUser(ctn Container) User {
	return &user{Ctn: ctn}
}

func (p *user) Save(ctx context.Context, form UserForm) error {
	client := p.Ctn.Client

	id, err := uuid.NewV4()
	if err != nil {
		return errors.Wrap(err, "failed generate uuid")
	}
	uid := id.String()

	key := generateKey(client, userEntityName, p.Ctn.Env, uid)
	entity := UserEntity{
		AccountID: form.AccountID,
		PixelaURL: fmt.Sprintf("%v/%v/graphs/%v", pixelaURL, form.AccountID, form.Graph),
	}
	if _, err := client.Put(ctx, key, &entity); err != nil {
		return errors.Wrap(err, "failed put user entity")
	}
	return nil
}