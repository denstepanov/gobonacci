package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/denstepanov/gobonacci/first/handlers"
	"github.com/denstepanov/gobonacci/shared/config"
	"github.com/denstepanov/gobonacci/shared/constants"
)

func main() {
	config := config.GetConfig(constants.FirstService)
	if config == nil {
		panic("Config data didn't find.")
	}

	r := chi.NewRouter()

	r.Get("/api/init-counts/{num}", handlers.InitCount)

	err := http.ListenAndServe(fmt.Sprintf("%s:%d", config.App.Host, config.App.Port), r)
	if err != nil {
		panic(err)
	}
}
