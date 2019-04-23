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

	ch := CreateRecoutHandler{
		EnvVar: EnvVar{
			Env: os.Getenv("RO_ENV"),
		},
	}
	r.Post("/recout", ch.ServeHTTP)

	http.Handle("/", r)
	appengine.Main()
}
