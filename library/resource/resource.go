package resource

import (
	"context"
	"sync"
)

var (
	once      = sync.Once{}
	initFuncs = []initFunc{
		initGORM,
	}
)

type initFunc func(ctx context.Context) error

func MustInit(ctx context.Context) {
	once.Do(func() {
		for _, f := range initFuncs {
			if err := f(ctx); err != nil {
				panic(err)
			}
		}
	})
}
