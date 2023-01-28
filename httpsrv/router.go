package httpsrv

import (
	"context"
	"github.com/goclover/clover"
	"github.com/goclover/clover/middleware"
	"github.com/goclover/clover/render"
	"net/http"
)

var Router = func(r clover.Router) {
	r.Use(middleware.Logger)
	r.Handle("/", func(c context.Context, r *http.Request) render.Render {
		return render.Text("Hello World")
	})
}
