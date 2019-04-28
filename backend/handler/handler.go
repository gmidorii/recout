package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gmidorii/recout/backend/app"
	"github.com/gmidorii/recout/backend/config"
	"github.com/gmidorii/recout/backend/form"
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

type CreateRecout struct {
	config.Config
}

func (c CreateRecout) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var form form.Recout
	if err := decoder.Decode(&form); err != nil {
		log.Println(err)
		render.Status(r, http.StatusBadRequest)
		return
	}

	ctn := app.Container{
		Env:      c.Config.Env,
		Location: c.Config.Location,
		Client:   c.Config.Client,
	}
	service := app.NewRecoutService(ctn)
	ctx := appengine.NewContext(r)
	uid, err := service.Create(ctx, form)
	if err != nil {
		log.Println(err)
		render.Status(r, http.StatusInternalServerError)
		return
	}

	log.Println(uid)
	fmt.Fprint(w, uid)
	w.WriteHeader(http.StatusOK)
	render.Status(r, http.StatusOK)
}

type GetRecout struct {
	config.Config
}

func (g GetRecout) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	form, err := form.FactoryFetchForm(r.URL.Query())
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		return
	}

	ctn := app.Container{
		Env:      g.Config.Env,
		Location: g.Config.Location,
		Client:   g.Config.Client,
	}
	service := app.NewRecoutService(ctn)
	res, err := service.Fetch(appengine.NewContext(r), form)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		log.Printf("failed service :%v", err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, res)
}

type PostUser struct {
	config.Config
}

func (p PostUser) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var form form.User
	if err := decoder.Decode(&form); err != nil {
		render.Status(r, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	ctn := app.Container{
		Env:      p.Config.Env,
		Location: p.Config.Location,
		Client:   p.Config.Client,
	}
	service := app.NewUser(ctn)
	if err := service.Save(appengine.NewContext(r), form); err != nil {
		log.Println(err)
		render.Status(r, http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusOK)
}
