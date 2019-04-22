package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"google.golang.org/appengine"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", indexHandler)
	r.Post("/recout", createRecoutHandler)

	http.Handle("/", r)
	appengine.Main()
}
