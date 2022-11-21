package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/denstepanov/gobonacci/shared/config"
	"github.com/denstepanov/gobonacci/shared/constants"
	"github.com/denstepanov/gobonacci/shared/proto"
)

func Send(prev, curr []byte) (*proto.CountDto, error) {
	cfg := config.GetConfig(constants.SecondService)

	client := proto.NewCountHandlerProtobufClient(
		fmt.Sprintf("%s://%s:%d", cfg.App.Protocol, cfg.App.Host, cfg.App.Port),
		&http.Client{},
	)

	ctx := context.Background()

	result, err := client.CountNext(ctx, &proto.CountDto{Prev: prev, Curr: curr})
	return result, err
}
