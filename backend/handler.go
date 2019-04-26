package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/render"
	"google.golang.org/appengine"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	fmt.Fprint(w, "Hello World!!")
}

type CreateRecoutHandler struct {
	Config
}

func (c CreateRecoutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var form RecoutForm
	if err := decoder.Decode(&form); err != nil {
		log.Println(err)
		render.Status(r, http.StatusBadRequest)
		return
	}

	ctn := Container{
		Env:      c.Config.Env,
		Location: c.Config.Location,
	}
	service := NewRecoutService(ctn)
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

type GetRecoutHandler struct {
	Config
}

func (g GetRecoutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	form, err := FactoryFetchForm(r.URL.Query())
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		return
	}

	ctn := Container{
		Env:      g.Config.Env,
		Location: g.Config.Location,
	}
	service := NewRecoutService(ctn)
	res, err := service.Fetch(appengine.NewContext(r), form)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		log.Printf("failed service :%v", err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, res)
}
