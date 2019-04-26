package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"google.golang.org/appengine"
)

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

	env := EnvVar{
		Env: os.Getenv("RO_ENV"),
	}

	ch := CreateRecoutHandler{EnvVar: env}
	r.Post("/recout", ch.ServeHTTP)

	gh := GetRecoutHandler{EnvVar: env}
	r.Get("/recout", gh.ServeHTTP)

	http.Handle("/", r)
	appengine.Main()
}
