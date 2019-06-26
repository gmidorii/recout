package app

import (
	"context"
	"log"

	"github.com/gmidorii/recout/backend/form"
	"github.com/gmidorii/recout/backend/infra/entity"
	"github.com/gmidorii/recout/backend/infra/pixela"
	"github.com/gmidorii/recout/backend/infra/repository"
	"github.com/gmidorii/recout/backend/response"
	"golang.org/x/xerrors"
)

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

	accountID := toAccountID(form.AccountID)

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
	token := p.ctn.Generator.Do(20)
	accountID := toAccountID(form.AccountID)

	pixelaEntity := pixela.User{
		Token:               token,
		UserName:            accountID,
		AgreeTermsOfService: pixela.Yes,
		NotMinor:            pixela.Yes,
	}
	if err := p.pixelaClient.CreateUser(pixelaEntity); err != nil {
		return xerrors.Errorf("failed create pixela user:%v", err)
	}

	graphID := p.ctn.Generator.Do(10)
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
