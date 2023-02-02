package resource

import (
	"context"
	"sync"
)

var (
	once      = sync.Once{}
	initFuncs = []initFunc{
		initGORM,
		initRedis,
	}
	// Conf configuration dir
	Conf = "conf"
)

type initFunc func(ctx context.Context, cf string) error

func MustInit(ctx context.Context) {
	once.Do(func() {
		for _, f := range initFuncs {
			if err := f(ctx, Conf); err != nil {
				panic(err)
			}
		}
	})
}
