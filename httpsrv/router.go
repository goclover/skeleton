package httpsrv

import (
	"context"
	"github.com/goclover/clover"
	"github.com/goclover/clover/middleware"
	"github.com/goclover/clover/render"
	"github.com/goclover/skeleton/httpsrv/handler"
	"net/http"
)

var Router = func(r clover.Router) {
	r.Use(middleware.Logger)

	r.Method("GET", "/", func(c context.Context, r *http.Request) render.Render {
		return render.Text("Hello World")
	})

	r.Handle("/hello", handler.Index())
}
