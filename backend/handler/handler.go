package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gmidorii/recout/backend/app"
	"github.com/gmidorii/recout/backend/config"
	"github.com/gmidorii/recout/backend/form"
	"github.com/gmidorii/recout/backend/injector"
	"github.com/go-chi/render"
	"google.golang.org/appengine"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	fmt.Fprint(w, "Hello World!!")
}

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

	ctn := app.Container{
		Env:      rh.Config.Env,
		Now:      time.Now(),
		Location: rh.Config.Location,
	}
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

	ctn := app.Container{
		Env:      rh.Config.Env,
		Location: rh.Config.Location,
	}
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

	ctn := app.Container{
		Env:      rh.Config.Env,
		Now:      time.Now(),
		Location: rh.Config.Location,
	}
	service, err := injector.InitRecoutApp(rh.Config.Client, ctn, ctn.Env)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	res, err := service.FetchContinues(appengine.NewContext(r), form)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, res)
}

type User struct {
	config.Config
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
