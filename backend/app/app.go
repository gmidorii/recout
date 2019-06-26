package app

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gmidorii/recout/backend/form"
	"github.com/gmidorii/recout/backend/infra/entity"
	"github.com/gmidorii/recout/backend/infra/pixela"
	"github.com/gmidorii/recout/backend/infra/repository"
	"github.com/gmidorii/recout/backend/response"
	"github.com/pkg/errors"
	"github.com/rs/xid"
	"golang.org/x/xerrors"
)

const (
	pixelaURL         = "https://pixe.la/v1/users"
	pixelaHeaderToken = "X-USER-TOKEN"

	durationDay = 24 * time.Hour
	hoursPerDay = 24
	idLen       = 16
)

type Container struct {
	Env      string
	Now      time.Time
	Location *time.Location
}

func NewContainer(env string, now time.Time, location time.Location) Container {
	return Container{
		Env:      env,
		Now:      now,
		Location: &location,
	}
}

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
	//TODO: fix to user login account id
	const accountID = "gmidorii"

	userEntity, err := r.repoUser.Get(ctx, accountID)
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
		AccountID: "gmidorii",
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
			//TODO: put new entity.
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
	entities, err := r.repoRecout.Fetch(ctx, 0, form.Limit)
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
	_, e, err := r.repoContinues.Get(ctx, form.AccountID)
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

func subDate(before, after time.Time) int {
	trBefore := before.Truncate(durationDay)
	trAfter := after.Truncate(durationDay)
	hours24Divisible := trAfter.Sub(trBefore).Hours()
	return int(hours24Divisible) / hoursPerDay
}

type User interface {
	Fetch(ctx context.Context, form form.User) (response.User, error)
	Save(ctx context.Context, form form.User) error
}

type user struct {
	ctn          Container
	repoUser     repository.User
	pixelaClient pixela.Client
}

func NewUser(ctn Container, repoUser repository.User, client pixela.Client) User {
	return &user{
		ctn:          ctn,
		repoUser:     repoUser,
		pixelaClient: client,
	}
}

func (u *user) Fetch(ctx context.Context, form form.User) (response.User, error) {
	//XXX: check all user.
	//us, err := u.repoUser.Fetch(ctx, 0, 100)
	//if err != nil {
	//	return response.User{}, err
	//}
	//log.Printf("all user: %v\n", us)

	accountID := fmt.Sprintf("rec%v", strings.ToLower(form.AccountID))

	userEntity, err := u.repoUser.Get(ctx, accountID)
	if err != nil {
		return response.User{}, err
	}

	log.Println(userEntity)
	return response.User{
		AccountID:   userEntity.AccountID,
		PixelaGraph: userEntity.PixelaGraph,
	}, nil
}

func (p *user) Save(ctx context.Context, form form.User) error {
	guid := xid.New()
	token := guid.String()
	accountID := fmt.Sprintf("rec%v", strings.ToLower(form.AccountID))

	pixelaEntity := pixela.User{
		Token:               token,
		UserName:            accountID,
		AgreeTermsOfService: pixela.Yes,
		NotMinor:            pixela.Yes,
	}
	if err := p.pixelaClient.CreateUser(pixelaEntity); err != nil {
		return xerrors.Errorf("failed create pixela user:%v", err)
	}

	guid = xid.New()
	graphID := guid.String()[:idLen]
	graphName := generateGraphName(accountID)
	if err := p.pixelaClient.CreateGraph(graphID, graphName, accountID, token); err != nil {
		return xerrors.Errorf("failed create pixela graph:%v", err)
	}

	entity := entity.User{
		AccountID:   accountID,
		PixelaGraph: graphID,
		PixelaToken: token,
	}

	if _, err := p.repoUser.Put(ctx, entity); err != nil {
		return xerrors.Errorf("faild create entity user: %v", err)
	}
	return nil
}

func generateGraphName(origin string) string {
	return strings.ReplaceAll(strings.ToLower(origin), " ", "")
}
