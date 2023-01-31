package render

import (
	"context"
	"github.com/goclover/clover/render"
)

type ResponseRender func(ctx context.Context, data interface{}, err error) render.Render
