package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gmidorii/recout/backend/config"
	"github.com/gmidorii/recout/backend/form"
	"github.com/gmidorii/recout/backend/infra/repository"
	"github.com/gmidorii/recout/backend/injector"
	"github.com/go-chi/render"
	"golang.org/x/xerrors"
	"google.golang.org/appengine"
)

type Recout struct {
	config.Config
}

func (rh Recout) Post(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var form form.Recout
	if err := decoder.Decode(&form); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctn := configToContainer(rh.Config)
	service, err := injector.InitRecoutApp(rh.Config.Client, ctn, ctn.Env)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	ctx := appengine.NewContext(r)
	uid, err := service.Create(ctx, form)
	if err != nil {
		log.Printf("failed recout create service: %+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println(uid)
	fmt.Fprint(w, uid)
	w.WriteHeader(http.StatusOK)
	render.Status(r, http.StatusOK)
}

func (rh Recout) Get(w http.ResponseWriter, r *http.Request) {
	form, err := form.FactoryFetchForm(r.URL.Query())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctn := configToContainer(rh.Config)
	service, err := injector.InitRecoutApp(rh.Config.Client, ctn, ctn.Env)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	res, err := service.Fetch(appengine.NewContext(r), form)
	if err != nil {
		log.Printf("failed service :%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, res)
}

func (rh Recout) GetContinues(w http.ResponseWriter, r *http.Request) {
	form, err := form.FactoryContinues(r.URL.Query())
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctn := configToContainer(rh.Config)
	service, err := injector.InitRecoutApp(rh.Config.Client, ctn, ctn.Env)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	res, err := service.FetchContinues(appengine.NewContext(r), form)
	if err != nil {
		if xerrors.Is(err, repository.NotFoundError{}) {
			log.Println(err)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, res)
}
