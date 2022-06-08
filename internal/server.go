package internal

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/open-scrm/open-scrm/configs"
	"github.com/open-scrm/open-scrm/internal/controller/addressbook"
	"github.com/open-scrm/open-scrm/internal/controller/callbaclcontroller"
	"github.com/open-scrm/open-scrm/internal/controller/configcontroller"
	"github.com/open-scrm/open-scrm/lib/log"
	"net"
	"net/http"
)

var (
	httpServer *http.Server
)

func RunHttpServer(ctx context.Context) error {
	config := configs.Get()
	_ = config

	g := gin.New()
	g.Use(gin.Recovery())
	g.Use(log.Trace())
	g.Use(log.Logger())

	api := g.Group("/api/v1")
	{
		{
			addressBook := api.Group("/addressbook")
			addressBook.POST("/sync", addressbook.SyncCorpStructure)
		}

		{
			config := api.Group("/config")
			config.POST("/talent", configcontroller.UpdateTalentInfo)
		}
	}

	g.Any("/callback/addressbook", callbaclcontroller.AddressBookCallback)

	httpServer = &http.Server{
		Addr:    config.Web.Addr,
		Handler: g,
		BaseContext: func(listener net.Listener) context.Context {
			return ctx
		},
	}

	log.WithContext(ctx).Infof("starting http server at: %v", config.Web.Addr)

	return httpServer.ListenAndServe()
}

func StopHttpServer(ctx context.Context) error {
	if httpServer != nil {
		return httpServer.Shutdown(ctx)
	}
	return nil
}
