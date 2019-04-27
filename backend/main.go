package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"go.mercari.io/datastore"
	"go.mercari.io/datastore/aedatastore"
	"google.golang.org/appengine"
)

const timeZone = "Asia/Tokyo"

type Config struct {
	Env      string
	Location *time.Location
	Client   datastore.Client
}

func NewConfig() (Config, error) {
	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		return Config{}, fmt.Errorf("failed location setting: %v", err)
	}

	client, err := aedatastore.FromContext(context.Background())
	if err != nil {
		return Config{}, fmt.Errorf("faild create datastore client: %v", err)
	}
	return Config{
		Env:      os.Getenv("RO_ENV"),
		Location: loc,
		Client:   client,
	}, nil
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	r.Use(c.Handler)

	r.Get("/", indexHandler)

	config, err := NewConfig()
	if err != nil {
		log.Fatalf("failed new config: %v", err)
	}

	ch := CreateRecoutHandler{Config: config}
	r.Post("/recout", ch.ServeHTTP)

	gh := GetRecoutHandler{Config: config}
	r.Get("/recout", gh.ServeHTTP)

	ph := PostPixelaHandler{Config: config}
	r.Post("/pixela", ph.ServeHTTP)

	http.Handle("/", r)
	appengine.Main()
}
