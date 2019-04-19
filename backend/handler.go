package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

	log.Println(form)
	w.WriteHeader(http.StatusOK)
}
