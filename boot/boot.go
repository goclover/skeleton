package boot

import (
	"cashier/httpsrv"
	"cashier/library/resource"
	"context"
	"github.com/goclover/clover"
)

func Start() {
	var c = clover.New()

	c.Group(httpsrv.Router)

	_ = c.Run(cs.HTTPServer.Addr)
}

func MustInit(ctx context.Context, _ *config) {
	resource.MustInit(ctx)
}
