package app

import (
	"context"
	"time"

	"github.com/gmidorii/recout/backend/form"
	"github.com/gmidorii/recout/backend/infra/entity"
	"github.com/gmidorii/recout/backend/infra/pixela"
	"github.com/gmidorii/recout/backend/infra/repository"
	"github.com/gmidorii/recout/backend/response"
	"github.com/pkg/errors"
	"golang.org/x/xerrors"
)

type Recout interface {
	Create(ctx context.Context, form form.Recout) (string, error)
	Fetch(ctx context.Context, form form.RecoutFetch) ([]response.RecoutFetch, error)
	FetchContinues(ctx context.Context, form form.RecoutContinues) (response.RecoutContinues, error)
}

type recout struct {
	ctn           Container
	repoRecout    repository.Recout
	repoUser      repository.User
	repoContinues repository.Continues
	pixelaClient  pixela.Client
}

func NewRecout(
	ctn Container,
	repoRecout repository.Recout,
	repoUser repository.User,
	repoContinues repository.Continues,
	pixelaClient pixela.Client,
) Recout {
	return &recout{
		ctn:           ctn,
		repoRecout:    repoRecout,
		repoUser:      repoUser,
		repoContinues: repoContinues,
		pixelaClient:  pixelaClient,
	}
}

func (r *recout) Create(ctx context.Context, form form.Recout) (uid string, err error) {
	accountID := encodeAccountID(form.AccountID)

	_, userEntity, err := r.repoUser.Get(ctx, accountID)
	if err != nil {
		return "", errors.Wrapf(err, "failed fetching user entity id = %v", accountID)
	}

	if err := r.pixelaClient.Increment(userEntity.AccountID, userEntity.PixelaToken, userEntity.PixelaGraph); err != nil {
		return "", errors.Wrapf(
			err,
			"failed pixela graph increment (user=%v, graph=%v)",
			userEntity.AccountID, userEntity.PixelaGraph,
		)
	}

	entityRecout := entity.Recout{
		AccountID: accountID,
		Message:   form.Message,
		CreatedAt: r.ctn.Now.In(r.ctn.Location).Unix(),
	}

	uid, err = r.repoRecout.Put(ctx, entityRecout)
	if err != nil {
		return "", xerrors.Errorf("failed put recout: %w", err)
	}

	continuesKey, continuesEntity, err := r.repoContinues.Get(ctx, accountID)
	if err != nil {
		switch err.(type) {
		case repository.NotFoundError:
			e := entity.Continues{
				AccountID: accountID,
				LastDate:  r.ctn.Now.Format(entity.DateLayout),
				Count:     1,
			}
			if err := r.repoContinues.Put(ctx, e); err != nil {
				return "", xerrors.Errorf("failed init put continues entity: %w", err)
			}
			return "", nil
		default:
			return "", xerrors.Errorf("failed get continues entity: %w", err)
		}
	}

	lastDate, err := time.Parse(entity.DateLayout, continuesEntity.LastDate)
	if err != nil {
		return "", xerrors.Errorf("%v is not %v layout: %w", continuesEntity.LastDate, entity.DateLayout, err)
	}
	day := subDate(lastDate, r.ctn.Now)
	switch day {
	case 0:
		// do nothing.
		return uid, nil
	case 1:
		continuesEntity.LastDate = r.ctn.Now.Format(entity.DateLayout)
		continuesEntity.Count += 1
	default:
		continuesEntity.LastDate = r.ctn.Now.Format(entity.DateLayout)
		continuesEntity.Count = 1
	}

	if err := r.repoContinues.PutKey(ctx, continuesKey, continuesEntity); err != nil {
		return "", xerrors.Errorf("failed update continues entity: %w", err)
	}
	return uid, nil
}

func (r *recout) Fetch(ctx context.Context, form form.RecoutFetch) ([]response.RecoutFetch, error) {
	entities, err := r.repoRecout.Fetch(ctx, encodeAccountID(form.AccountID), 0, form.Limit)
	if err != nil {
		return nil, errors.Wrap(err, "failed fetch recout entity from datastore")
	}

	responses := make([]response.RecoutFetch, len(entities))
	for i, e := range entities {
		responses[i] = response.RecoutFetch{
			Message:   e.Message,
			CreatedAt: response.JSONTime(time.Unix(e.CreatedAt, 0).In(r.ctn.Location)),
		}
	}
	return responses, nil
}

func (r *recout) FetchContinues(ctx context.Context, form form.RecoutContinues) (response.RecoutContinues, error) {
	_, e, err := r.repoContinues.Get(ctx, encodeAccountID(form.AccountID))
	if err != nil {
		return response.RecoutContinues{}, errors.Wrap(err, "failed fetch continues entity")
	}
	lastDate, err := time.Parse(entity.DateLayout, e.LastDate)
	if err != nil {
		return response.RecoutContinues{}, errors.Wrapf(err, "%v is not %v layout.", e.LastDate, entity.DateLayout)
	}
	if date := subDate(lastDate, r.ctn.Now); date > 1 {
		return response.RecoutContinues{Count: 0}, nil
	}

	return response.RecoutContinues{Count: e.Count}, nil
}
