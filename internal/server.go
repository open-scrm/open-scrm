package internal

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/open-scrm/open-scrm/configs"
	"github.com/open-scrm/open-scrm/internal/controller/addressbook"
	"github.com/open-scrm/open-scrm/internal/controller/callbaclcontroller"
	"github.com/open-scrm/open-scrm/internal/controller/configcontroller"
	"github.com/open-scrm/open-scrm/lib/log"
	"github.com/open-scrm/open-scrm/lib/session"
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

	conf := cors.DefaultConfig()
	conf.AllowAllOrigins = true
	conf.AllowHeaders = []string{"*"}
	g.Use(cors.New(conf))

	g.Use(log.Trace())
	g.Use(log.Logger())

	authRouter(g)

	api := g.Group("/api/v1", session.Auth())
	{
		{
			addressBook := api.Group("/addressbook")
			addressBook.POST("/sync", addressbook.SyncCorpStructure)
			addressBook.POST("/dept/list", addressbook.DepartmentList)
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
