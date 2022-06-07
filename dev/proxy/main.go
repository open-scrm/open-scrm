package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.clearcode.cn/wx/xrequest"
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

	req := xrequest.HomeClient().R()
	for h := range ctx.Request.Header {
		req.SetHeader(h, ctx.GetHeader(h))
		fmt.Println("< ", fmt.Sprintf("%s: %s", h, ctx.GetHeader(h)))
	}
	fmt.Println("< ")
	switch ctx.Request.Method {
	case http.MethodGet:
		fmt.Println("> ", ctx.Request.Method, "http://work.mrj.com:30080"+ctx.Request.RequestURI)
		response, err := req.Get("http://work.mrj.com:30080" + ctx.Request.RequestURI)
		if err != nil {
			log.Println("get error", err)
			return
		}
		fmt.Println("> ", response.StatusCode())
		fmt.Println("> ", string(response.Body()))
		ctx.String(response.StatusCode(), string(response.Body()))
	case http.MethodPost:
		data, _ := ioutil.ReadAll(ctx.Request.Body)
		req.Method = ctx.Request.Method
		req.SetBody(data)
		fmt.Println("< ", string(data))
		fmt.Println("< ")
		fmt.Println("> ", ctx.Request.Method, "http://work.mrj.com:30080"+ctx.Request.RequestURI)
		response, err := req.Post("http://work.mrj.com:30080" + ctx.Request.RequestURI)
		if err != nil {
			log.Println("post error", err)
			return
		}
		fmt.Println("> ", response.StatusCode())
		fmt.Println("> ", string(response.Body()))
		ctx.String(response.StatusCode(), string(response.Body()))
	}
}
