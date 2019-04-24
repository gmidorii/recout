package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"google.golang.org/appengine"
)

func main() {
	r := chi.NewRouter()

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
