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
	"go.pyspa.org/brbundle"
	"go.pyspa.org/brbundle/brchi"
	"google.golang.org/appengine"
)

const frontPackageName = "front.pb"

func pathRoute(r *chi.Mux, config config.Config) {
	r.Route("/recout", func(r chi.Router) {
		rh := handler.Recout{Config: config}
		r.Post("/", rh.Post)
		r.Get("/", rh.Get)
		r.Get("/continues", rh.GetContinues)
	})

	r.Route("/user", func(r chi.Router) {
		u := handler.User{Config: config}
		r.Post("/", u.Post)
	})

	brbundle.RegisterBundle(frontPackageName)
	r.Get("/*", brchi.Mount())
	r.Get("/", brchi.Mount(brbundle.WebOption{
		SPAFallback: "index.html",
	}))
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

	config, err := config.New()
	if err != nil {
		log.Fatalf("failed new config: %v", err)
	}

	pathRoute(r, config)
	http.Handle("/", r)
	appengine.Main()
}
