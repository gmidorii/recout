package main

import (
	"log"
	"net/http"

	"github.com/gmidorii/recout/backend/config"
	"github.com/gmidorii/recout/backend/handler"
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

	r.Get("/", handler.IndexHandler)

	config, err := config.New()
	if err != nil {
		log.Fatalf("failed new config: %v", err)
	}

	ch := handler.CreateRecout{Config: config}
	r.Post("/recout", ch.ServeHTTP)

	gh := handler.GetRecout{Config: config}
	r.Get("/recout", gh.ServeHTTP)

	ph := handler.PostUser{Config: config}
	r.Post("/user", ph.ServeHTTP)

	http.Handle("/", r)
	appengine.Main()
}
