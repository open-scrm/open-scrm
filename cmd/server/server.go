package server

import (
	"context"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/open-scrm/open-scrm/configs"
	"github.com/open-scrm/open-scrm/internal"
	"github.com/open-scrm/open-scrm/internal/global"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-lib/metrics/prometheus"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
	"os"
	"os/signal"
)

var (
	ServeCommand = &cli.Command{
		Name:        "serve",
		Description: "run open-scrm http server",
		Action: func(c *cli.Context) error {
			var e error
			gracefulShutdown(func(ctx context.Context) {
				c.Context = ctx
				if err := run(c); err != nil {
					e = err
					return
				}
			}, func() {
				if err := internal.StopHttpServer(c.Context); err != nil {
					e = err
				}
			})
			return e
		},
	}

	log = logrus.New().WithField("module", "server")
)

func run(ctx *cli.Context) error {
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

	if err := global.InitGlobal(ctx.Context, configs.Get()); err != nil {
		return err
	}

	data, _ := yaml.Marshal(configs.Get())
	fmt.Println("-- 加载配置信息 --")
	fmt.Println(string(data))

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

	return internal.RunHttpServer(ctx.Context)
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
