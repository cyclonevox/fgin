package test

import (
	"fmt"
	"time"

	"git.vox666.top/vox/fgin"
)

const testUrl = "http://127.0.0.1:9998"
const testAddr = ":9998"

func setUpTestServer(addr ...string) {
	go func() {
		r := fgin.New()
		// r.GET("/", func(ctx *fgin.Context) {
		// 	fmt.Println(ctx.Method(), ctx.Path())
		// })
		r.GET("/hello", func(ctx *fgin.Context) {
			fmt.Println(ctx.Method, ctx.Request.RequestURI)
			ctx.Data(200, "", []byte("hello world by /hello get"))
		})

		v1 := r.Group("/v1")
		v1.GET("/", func(ctx *fgin.Context) {
			fmt.Println(ctx.Method, ctx.Request.RequestURI)
		})
		v1.GET("/hello", func(ctx *fgin.Context) {
			fmt.Println(ctx.Method, ctx.Request.RequestURI)
			ctx.Data(200, "", []byte("hello world by /hello get"))
		})

		v2 := r.Group("/v2")
		v2.GET("/", func(ctx *fgin.Context) {
			fmt.Println(ctx.Method, ctx.Request.RequestURI)
			ctx.Data(200, "", []byte("hello world by /v2 get"))
		})
		v2.GET("/hello", func(ctx *fgin.Context) {
			fmt.Println(ctx.Method, ctx.Request.RequestURI)
			ctx.Data(200, "", []byte("hello world by /v2/hello get"))
		})

		v3 := r.Group("/v3").Use(fgin.Logger(), fgin.Recovery())
		v3.GET("/hello", func(ctx *fgin.Context) {
			fmt.Println(ctx.Method, ctx.Request.RequestURI)
			ctx.Data(200, "", []byte("hello world by /v3/hello get"))
		})
		v3.GET("/panic", func(ctx *fgin.Context) {
			fmt.Println(ctx.Method, ctx.Request.RequestURI)
			var a []string
			ctx.Data(200, "", []byte("hello world by /v3/hello get"+a[0]))
		})

		if len(addr) > 0 {
			panic(r.Run(addr[0]))
		} else {
			panic(r.Run(testAddr))
		}
	}()

	time.Sleep(2 * time.Second)
}
