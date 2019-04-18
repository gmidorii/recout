package main

import (
	"github.com/go-chi/chi"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", indexHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Default Port %v", port)
	}

	log.Printf("Listening on port %v", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), r))
}

