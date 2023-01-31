package handler

import (
	"context"
	"github.com/goclover/clover"
	"github.com/goclover/skeleton/library/handler"
)

type index struct {
}

func (p index) Execute(ctx context.Context, r clover.Request) (interface{}, error) {
	return map[string]interface{}{"hello": "world"}, nil
}

func Index() clover.HandlerFunc {
	return handler.NewHandlerFunc(index{})
}
