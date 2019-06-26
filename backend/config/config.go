package config

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gmidorii/recout/backend/app"
	"github.com/rs/xid"
	"go.mercari.io/datastore"
	"go.mercari.io/datastore/clouddatastore"
)

const timeZone = "Asia/Tokyo"

type Config struct {
	Env       string
	Location  *time.Location
	Client    datastore.Client
	Generator app.RandomGenerator
}

type randomXid struct{}

func (x *randomXid) Do() string {
	guid := xid.New()
	return guid.String()
}

func New() (Config, error) {
	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		return Config{}, fmt.Errorf("failed location setting: %v", err)
	}

	client, err := clouddatastore.FromContext(context.Background())
	if err != nil {
		return Config{}, fmt.Errorf("faild create datastore client: %v", err)
	}
	return Config{
		Env:       os.Getenv("RO_ENV"),
		Location:  loc,
		Client:    client,
		Generator: &randomXid{},
	}, nil
}
