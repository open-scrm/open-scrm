package internal

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/open-scrm/open-scrm/configs"
	"github.com/open-scrm/open-scrm/internal/controller/addressbook"
	"github.com/open-scrm/open-scrm/lib/log"
)

func RunHttpServer(ctx context.Context) {
	config := configs.Get()
	_ = config

	g := gin.New()
	g.Use(gin.Recovery())
	g.Use(log.Trace())
	g.Use(log.Logger())

	api := g.Group("/api/v1")
	{
		addressBook := api.Group("/addressbook")
		addressBook.POST("/sync", addressbook.SyncCorpStructure)
	}
}
