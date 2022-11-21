package handlers

import (
	"context"

	"github.com/twitchtv/twirp"

	"github.com/denstepanov/gobonacci/shared/proto"
	"github.com/denstepanov/gobonacci/shared/services"
)

type CountHandler struct {
	Count proto.CountDto
}

func (c *CountHandler) CountNext(ctx context.Context, dto *proto.CountDto) (*proto.CountDto, error) {
	if dto == nil {
		return dto, twirp.InvalidArgument.Error("CountDto is empty.")
	}

	prev, curr := services.Count(dto.Prev, dto.Curr)
	return &proto.CountDto{Curr: curr, Prev: prev}, nil
}
