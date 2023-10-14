package test

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"git.vox666.top/vox/fgin"
)

const testUrl = "http://127.0.0.1:9998"

type testBody []byte

func (t testBody) Read(p []byte) (n int, err error) {
	return 0, io.EOF
}

func TestHTTPBasic(t *testing.T) {

	go func() {
		r := fgin.New()
		r.GET("/", func(ctx *fgin.Context) {
			fmt.Println(ctx.Method(), ctx.Path())
		})
		r.GET("/hello", func(ctx *fgin.Context) {
			fmt.Println(ctx.Method(), ctx.Path())
			ctx.Data(200, []byte("hello world by /hello get"))
		})

		r.UPDATE("/", func(ctx *fgin.Context) {
			fmt.Println(ctx.Method(), ctx.Path())
			ctx.Data(200, []byte("hello world by / Update"))
		})
		r.UPDATE("/hello", func(ctx *fgin.Context) {
			fmt.Println(ctx.Method(), ctx.Path())
			ctx.Data(200, []byte("hello world by /hello Update"))
		})

		r.DELETE("/", func(ctx *fgin.Context) {
			fmt.Println(ctx.Method(), ctx.Path())
			ctx.Data(200, []byte("hello world by / Post"))
		})
		r.DELETE("/hello", func(ctx *fgin.Context) {
			fmt.Println(ctx.Method(), ctx.Path())
			ctx.Data(200, []byte("hello world by /hello delete"))
		})

		r.POST("/", func(ctx *fgin.Context) {
			fmt.Println(ctx.Method(), ctx.Path())
			ctx.Data(200, []byte("hello world by / Post"))
		})
		r.POST("/hello", func(ctx *fgin.Context) {
			fmt.Println(ctx.Method(), ctx.Path())
			ctx.Data(200, []byte("hello world by /hello Post"))
		})

		v2 := r.Group("/v2")
		v2.GET("/", func(ctx *fgin.Context) {
			fmt.Println(ctx.Method(), ctx.Path())
			ctx.Data(200, []byte("hello world by /v2 get"))
		})
		v2.GET("/hello", func(ctx *fgin.Context) {
			fmt.Println(ctx.Method(), ctx.Path())
			ctx.Data(200, []byte("hello world by /v2/hello get"))
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
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
	})

	t.Run("get2", func(t *testing.T) {

		resp, err := http.Get(testUrl + "/hello")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
	})

	t.Run("get1", func(t *testing.T) {

		resp, err := http.Get(testUrl + "/v2")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
	})

	t.Run("get2", func(t *testing.T) {

		resp, err := http.Get(testUrl + "/hello")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
	})

	t.Run("post", func(t *testing.T) {

		resp, err := http.Post(testUrl, "hello", testBody{})
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
	})

	t.Run("post2", func(t *testing.T) {
		resp, err := http.Post(testUrl+"/hello", "test", testBody{})
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
	})
}
