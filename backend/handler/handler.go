package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gmidorii/recout/backend/app"
	"github.com/gmidorii/recout/backend/config"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	fmt.Fprint(w, "Hello World!!")
}

func configToContainer(config config.Config) app.Container {
	return app.Container{
		Env:       config.Env,
		Now:       time.Now(),
		Location:  config.Location,
		Generator: config.Generator,
	}
}
