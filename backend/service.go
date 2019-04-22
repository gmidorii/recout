package main

import (
	"context"
	"time"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"go.mercari.io/datastore/aedatastore"
)

type RecoutService interface {
	Create(form RecoutForm) (string, error)
}

type recoutService struct {
}

func NewRecoutService() RecoutService {
	return &recoutService{}
}

func (r *recoutService) Create(form RecoutForm) (uid string, err error) {
	client, err := aedatastore.FromContext(context.Background())
	if err != nil {
		return "", errors.Wrap(err, "failed init aeclient")
	}
	defer client.Close()

	id, err := uuid.NewV4()
	if err != nil {
		return "", errors.Wrap(err, "failed generate uuid")
	}
	uid = id.String()

	key := client.NameKey("RecoutEntity", uid, nil)
	entity := RecoutEntity{
		Message:   form.Message,
		CreatedAt: time.Now().Unix(),
	}
	_, err = client.Put(context.Background(), key, &entity)
	if err != nil {
		return "", errors.Wrap(err, "failed put datastore")
	}
	return
}
