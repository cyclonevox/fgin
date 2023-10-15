package test

import (
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"git.vox666.top/vox/fgin"
)

const testUrl = "http://127.0.0.1:9998"

func TestHTTPBasic(t *testing.T) {

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

		panic(r.Run(":9998"))
	}()

	time.Sleep(2 * time.Second)

	t.Run("get1", func(t *testing.T) {

		resp, err := http.Get(testUrl)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		fmt.Println(string(body))
	})

	t.Run("get2", func(t *testing.T) {

		resp, err := http.Get(testUrl + "/hello")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		fmt.Println(string(body))
	})

	t.Run("get3", func(t *testing.T) {

		resp, err := http.Get(testUrl + "/v2")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		fmt.Println(string(body))
	})

	t.Run("get4", func(t *testing.T) {

		resp, err := http.Get(testUrl + "/v1/hello")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		fmt.Println(string(body))
	})
}
