package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	httpSwagger "github.com/swaggo/http-swagger"

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

	r.Mount("/swagger", httpSwagger.WrapHandler)
	r.Get("/api/init-counts/{num:[1-9]+}", handlers.InitCount)

	err := http.ListenAndServe(fmt.Sprintf("%s:%d", config.App.Host, config.App.Port), r)
	if err != nil {
		panic(err)
	}
}