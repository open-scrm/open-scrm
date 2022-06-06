package server

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/open-scrm/open-scrm/configs"
	"github.com/open-scrm/open-scrm/internal/global"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-lib/metrics/prometheus"
	"github.com/urfave/cli/v2"
)

var (
	ServeCommand = &cli.Command{
		Name:        "serve",
		Description: "run open-scrm http server",
		Action:      run,
	}

	log = logrus.New().WithField("module", "server")
)

func run(ctx *cli.Context) error {
	viper.SetDefault("web.addr", "127.0.0.1:8080")
	viper.SetDefault("web.static", "web/static")
	viper.SetDefault("web.view", "web/view/*/*.gohtml")

	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return err
	}

	if err := configs.ReloadConfig(); err != nil {
		return err
	}
	viper.OnConfigChange(func(in fsnotify.Event) {
		if err := configs.ReloadConfig(); err != nil {
			log.Errorf("reloadConfig failed: %v", err)
		}
	})
	viper.WatchConfig()

	if err := global.InitGlobal(ctx.Context, &configs.Get().Redis); err != nil {
		return err
	}

	metricsFactory := prometheus.New()
	tracer, closeTracer, err := config.Configuration{ServiceName: "open-scrm"}.
		NewTracer(
			config.Metrics(metricsFactory),
		)
	if err != nil {
		panic(err)
	}
	opentracing.SetGlobalTracer(tracer)
	defer closeTracer.Close()

	return nil
	//server := server2.NewSever()
	//return server.Run(ctx.Context)
}
