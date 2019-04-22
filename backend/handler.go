package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"google.golang.org/appengine"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	fmt.Fprint(w, "Hello World!!")
}

type RecoutForm struct {
	Message string `json:message`
}

func createRecoutHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var form RecoutForm
	if err := decoder.Decode(&form); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	service := NewRecoutService()
	ctx := appengine.NewContext(r)
	uid, err := service.Create(ctx, form)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println(uid)
	fmt.Fprint(w, uid)
	w.WriteHeader(http.StatusOK)
}
