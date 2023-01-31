package handler

import (
	"context"
	"github.com/goclover/clover"
	"github.com/goclover/clover/render"
	cr "github.com/goclover/skeleton/library/render"
	"net/http"
)

type Executor interface {
	Execute(ctx context.Context, r clover.Request) (interface{}, error)
}

type Handler struct {
	e      Executor
	render cr.ResponseRender
}

func NewHandlerFunc(e Executor) clover.HandlerFunc {
	h := Handler{e: e}
	return h.do
}

func (h *Handler) do(c context.Context, r *http.Request) render.Render {
	var data, err = h.e.Execute(c, clover.NewRequest(r))
	if h.render == nil {
		h.render = cr.JSON
	}
	return h.render(c, data, err)
}
