package main

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"go.mercari.io/datastore"
	"go.mercari.io/datastore/aedatastore"
)

type Container struct {
	Env string
}

type RecoutService interface {
	Create(ctx context.Context, form RecoutForm) (string, error)
}

type recoutService struct {
	Ctn Container
}

func NewRecoutService(ctn Container) RecoutService {
	return &recoutService{Ctn: ctn}
}

func (r *recoutService) Create(ctx context.Context, form RecoutForm) (uid string, err error) {
	client, err := aedatastore.FromContext(ctx)
	if err != nil {
		return "", errors.Wrap(err, "failed init aeclient")
	}
	defer client.Close()

	id, err := uuid.NewV4()
	if err != nil {
		return "", errors.Wrap(err, "failed generate uuid")
	}
	uid = id.String()

	key := generateKey(client, "RecoutEntity", r.Ctn.Env, uid)
	entity := RecoutEntity{
		Message:   form.Message,
		CreatedAt: time.Now().Unix(),
	}
	_, err = client.Put(ctx, key, &entity)
	if err != nil {
		return "", errors.Wrap(err, "failed put datastore")
	}
	return
}

func generateKey(client datastore.Client, kind, env, uid string) string {
	return client.NameKey(fmt.Sprintf("%v_%v", env, kind), uid, nil)
}
