package main

import (
	"fmt"
	"net/http"

	"github.com/denstepanov/gobonacci/second/handlers"
	"github.com/denstepanov/gobonacci/shared/config"
	"github.com/denstepanov/gobonacci/shared/constants"
	"github.com/denstepanov/gobonacci/shared/proto"
)

func main() {
	cfg := config.GetConfig(constants.SecondService)
	if cfg == nil {
		panic("Config data didn't find.")
	}

	countHandler := proto.NewCountHandlerServer(&handlers.CountHandler{})
	mux := http.NewServeMux()
	mux.Handle(countHandler.PathPrefix(), countHandler)

	err := http.ListenAndServe(fmt.Sprintf("%s:%d", cfg.App.Host, cfg.App.Port), countHandler)
	if err != nil {
		panic(err)
	}
}
