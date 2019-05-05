package repository

import (
	"context"

	"github.com/gmidorii/recout/backend/infra/entity"
)

type Recout interface {
	Put(ctx context.Context, e entity.Recout) (string, error)
	Fetch(ctx context.Context, offset, limit int) ([]entity.Recout, error)
}

type User interface {
	Put(ctx context.Context, e entity.User) (string, error)
	Get(ctx context.Context, accountID string) (entity.User, error)
	Fetch(ctx context.Context, offset, limit int) ([]entity.User, error)
}