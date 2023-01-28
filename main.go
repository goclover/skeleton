package main

import (
	"cashier/boot"
	"context"
	"flag"
)

var cf = flag.String("c", "conf/app.toml", "Application configuration file")

func main() {
	flag.Parse()

	var cs = boot.MustLoadConfig(*cf)
	var ctx, cc = context.WithCancel(context.Background())
	defer cc()

	boot.MustInit(ctx, cs)
	boot.Start()
}
