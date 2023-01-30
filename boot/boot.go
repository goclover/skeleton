package boot

import (
	"context"
	"github.com/goclover/clover"
	"github.com/goclover/skeleton/httpsrv"
	"github.com/goclover/skeleton/library/resource"
)

func Start() {
	var c = clover.New()

	c.Group(httpsrv.Router)

	_ = c.Run(cs.HTTPServer.Addr)
}

func MustInit(ctx context.Context, _ *config) {
	resource.MustInit(ctx)
}
