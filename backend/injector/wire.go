//+build wireinject

package injector

import (
	"github.com/gmidorii/recout/backend/app"
	lds "github.com/gmidorii/recout/backend/infra/datastore"
	"github.com/gmidorii/recout/backend/infra/pixela"
	"github.com/google/wire"
	"go.mercari.io/datastore"
)

func InitRecoutApp(gClient datastore.Client, ctn app.Container, env string) (r app.Recout, err error) {
	wire.Build(lds.NewRecoutClient, lds.NewUserClient, lds.NewContinuesClient, pixela.NewClient, app.NewRecout)
	return
}

func InitUserApp(gClient datastore.Client, ctn app.Container, env string) (r app.User, err error) {
	wire.Build(lds.NewUserClient, app.NewUser, pixela.NewClient)
	return
}
