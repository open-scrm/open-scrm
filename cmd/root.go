package main

import (
	"context"
	"github.com/open-scrm/open-scrm/cmd/server"
	"github.com/urfave/cli/v2"
	"log"
	"os"
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
	if err := app.RunContext(context.Background(), os.Args); err != nil {
		log.Println(err)
	}
}
