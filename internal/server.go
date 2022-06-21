package internal

import (
	"context"
	"github.com/clearcodecn/swaggos"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/open-scrm/open-scrm/configs"
	"github.com/open-scrm/open-scrm/internal/controller/addressbook"
	"github.com/open-scrm/open-scrm/internal/controller/callbaclcontroller"
	"github.com/open-scrm/open-scrm/internal/controller/configcontroller"
	"github.com/open-scrm/open-scrm/internal/controller/customer"
	"github.com/open-scrm/open-scrm/internal/vo"
	"github.com/open-scrm/open-scrm/lib/log"
	"github.com/open-scrm/open-scrm/lib/session"
	"github.com/open-scrm/open-scrm/pkg/addressbook/mapper"
	"github.com/open-scrm/open-scrm/pkg/response"
	addressbook2 "github.com/open-scrm/open-scrm/pkg/response/addressbook"
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

	doc := swaggos.Default()
	doc.APIKeyAuth("Authorization", "header")
	doc.Response(200, response.NewResponse(struct{}{}))

	authRouter(g, doc)
	// 文档
	g.GET("/api/doc.json", gin.WrapH(doc))
	g.Any("/doc/*path", gin.WrapH(swaggos.UI("/doc", "http://localhost:8080/api/doc.json")))

	api := g.Group("/api/v1", session.Auth())
	apiDocGroup := doc.Group("/api/v1")
	{
		{
			addressBookDocGroup := apiDocGroup.Group("/addressbook").Tag("addressbook")

			addressBook := api.Group("/addressbook")
			addressBook.POST("/sync", addressbook.SyncCorpStructure)
			addressBookDocGroup.Post("/sync").Body(struct{}{}).Description("同步组织架构,从企微员工部门信息")

			addressBook.POST("/dept/list", addressbook.DepartmentList)
			addressBookDocGroup.Post("/dept/list").Body(new(vo.AddressBookListDeptRequest)).JSON(new(mapper.DepartmentTree)).Description("获取部门树形结构")

			addressBook.POST("/user/list", addressbook.UserList)
			addressBookDocGroup.Post("/user/list").Body(new(vo.UserListRequest)).JSON(new(addressbook2.UserListResponse)).Description("获取员工列表")
		}
		{
			customerRouter := api.Group("/customer", session.Auth())
			{
				customerRouter.POST("/syncall", customer.SyncAll)
				customerRouter.POST("/list", customer.ListCustomer)
			}
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
