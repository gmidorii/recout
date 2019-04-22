package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"google.golang.org/appengine"
)

func main() {
	r := chi.NewRouter()
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
