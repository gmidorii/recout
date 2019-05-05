package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gmidorii/recout/backend/form"
	"github.com/gmidorii/recout/backend/infra/entity"
	"github.com/gmidorii/recout/backend/infra/repository"
	"github.com/gmidorii/recout/backend/response"
	"github.com/pkg/errors"
)

const (
	pixelaURL         = "https://pixe.la/v1/users"
	pixelaHeaderToken = "X-USER-TOKEN"
)

type Container struct {
	Env      string
	Location *time.Location
}

func NewContainer(env string, location time.Location) Container {
	return Container{
		Env:      env,
		Location: &location,
	}
}

type Recout interface {
	Create(ctx context.Context, form form.Recout) (string, error)
	Fetch(ctx context.Context, form form.RecoutFetch) ([]response.RecoutFetch, error)
}

type recout struct {
	ctn        Container
	repoRecout repository.Recout
	repoUser   repository.User
}

func NewRecout(ctn Container, repoRecout repository.Recout, repoUser repository.User) Recout {
	return &recout{
		ctn:        ctn,
		repoRecout: repoRecout,
		repoUser:   repoUser,
	}
}

func (r *recout) Create(ctx context.Context, form form.Recout) (uid string, err error) {
	//TODO: fix to user login account id
	const accountID = "gmidorii"
	log.Println(r.ctn.Env)

	userEntity, err := r.repoUser.Get(ctx, accountID)
	if err != nil {
		return "", errors.Wrapf(err, "failed fetching user entity id = %v", accountID)
	}

	url := fmt.Sprintf("%v/%v/graphs/%v/increment", pixelaURL, userEntity.AccountID, userEntity.PixelaGraph)
	log.Println(url)
	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		return "", errors.Wrap(err, "failed new http request")
	}
	req.Header.Add(pixelaHeaderToken, userEntity.PixelaToken)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	entity := entity.Recout{
		AccountID: "gmidorii",
		Message:   form.Message,
		CreatedAt: time.Now().In(r.ctn.Location).Unix(),
	}

	return r.repoRecout.Put(ctx, entity)
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

type User interface {
	Save(ctx context.Context, form form.User) error
}

type user struct {
	ctn      Container
	repoUser repository.User
}

func NewUser(ctn Container, repoUser repository.User) User {
	return &user{
		ctn:      ctn,
		repoUser: repoUser,
	}
}

func (p *user) Save(ctx context.Context, form form.User) error {
	entity := entity.User{
		AccountID:   form.AccountID,
		PixelaGraph: form.PixelaGraph,
		PixelaToken: form.PixelaToken,
	}
	if _, err := p.repoUser.Put(ctx, entity); err != nil {
		return errors.Wrap(err, "failed create user")
	}
	return nil
}
