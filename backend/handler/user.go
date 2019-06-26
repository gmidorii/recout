package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gmidorii/recout/backend/app"
	"github.com/gmidorii/recout/backend/config"
	"github.com/gmidorii/recout/backend/form"
	"github.com/gmidorii/recout/backend/infra/repository"
	"github.com/gmidorii/recout/backend/injector"
	"github.com/go-chi/render"
	"golang.org/x/xerrors"
	"google.golang.org/appengine"
)

type User struct {
	config.Config
}

func (u User) Get(w http.ResponseWriter, r *http.Request) {
	user, err := form.FactoryUser(r.URL.Query())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	ctn := app.Container{
		Env: u.Config.Env,
	}
	service, err := injector.InitUserApp(u.Config.Client, ctn, ctn.Env)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	res, err := service.Fetch(appengine.NewContext(r), user)
	if err != nil {
		if xerrors.Is(err, repository.NotFoundError{}) {
			w.WriteHeader(http.StatusNotFound)
			log.Println(err)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, res)
}

func (u User) Post(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var form form.User
	if err := decoder.Decode(&form); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	ctn := app.Container{
		Env:      u.Config.Env,
		Location: u.Config.Location,
	}
	service, err := injector.InitUserApp(u.Config.Client, ctn, ctn.Env)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := service.Save(appengine.NewContext(r), form); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusOK)
}
