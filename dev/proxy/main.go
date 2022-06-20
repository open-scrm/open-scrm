package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.clearcode.cn/wx/xrequest"
	"gopkg.in/resty.v1"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	go func() {
		defer func() {
			recover()
		}()

		xrequest.StartDNSServer()
	}()
	g := gin.New()
	g.Use(gin.Recovery())
	gin.SetMode(gin.ReleaseMode)
	g.NoRoute(handler)

	g.Run(":9090")
}

func handler(ctx *gin.Context) {
	fmt.Println("< ", ctx.Request.Method, ctx.Request.RequestURI)

	req := xrequest.HomeClient().SetRedirectPolicy(resty.FlexibleRedirectPolicy(3)).R()
	for h := range ctx.Request.Header {
		req.SetHeader(h, ctx.GetHeader(h))
		fmt.Println("< ", fmt.Sprintf("%s: %s", h, ctx.GetHeader(h)))
	}
	var response *resty.Response
	var err error

	fmt.Println("< ")
	switch ctx.Request.Method {
	case http.MethodGet, http.MethodDelete, http.MethodHead, http.MethodOptions:
		fmt.Println("> ", ctx.Request.Method, "http://work.mrj.com:30080"+ctx.Request.RequestURI)
		response, err = req.Get("http://work.mrj.com:30080" + ctx.Request.RequestURI)
	case http.MethodPost, http.MethodPut, http.MethodPatch:
		data, _ := ioutil.ReadAll(ctx.Request.Body)
		req.Method = ctx.Request.Method
		req.SetBody(data)
		fmt.Println("< ", string(data))
		fmt.Println("< ")
		fmt.Println("> ", ctx.Request.Method, "http://work.mrj.com:30080"+ctx.Request.RequestURI)
		response, err = req.Execute(req.Method, "http://work.mrj.com:30080"+ctx.Request.RequestURI)
	}

	if err != nil {
		log.Println("get error", err)
		ctx.String(200, "错误: %v", err.Error())
		return
	}
	fmt.Println("> ", response.Status())
	for h := range response.Header() {
		fmt.Println("> ", h, response.Header().Get(h))
		ctx.Header(h, response.Header().Get(h))
	}
	fmt.Println("> ", string(response.Body()))
	ctx.String(response.StatusCode(), string(response.Body()))
}
