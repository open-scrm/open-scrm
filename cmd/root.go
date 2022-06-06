package main

import (
	"context"
	"github.com/open-scrm/open-scrm/cmd/server"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"os/signal"
)

const (
	appName = "app"
)

var (
	app = cli.NewApp()
)

func init() {
	app.Name = appName
	app.Commands = []*cli.Command{
		server.ServeCommand,
	}
}

func main() {
	gracefulShutdown(func(ctx context.Context) {
		if err := app.RunContext(ctx, os.Args); err != nil {
			log.Println(err)
		}
	})
}

func gracefulShutdown(process func(ctx context.Context), callbacks ...func()) {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		process(ctx)
		cancel()
	}()
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Kill, os.Interrupt)

	defer func() {
		for _, fn := range callbacks {
			fn()
		}
	}()

	select {
	case <-ctx.Done():
		cancel()
	case <-sig:
		cancel()
	}
}
