package fgin

import (
	"fmt"

	"git.vox666.top/vox/fgin/internal/util"
	"github.com/valyala/fasthttp"
)

type HandlerFunc func(ctx *Context)

type Engine struct {
	*RouterGroup
	router router
	groups []*RouterGroup
}

func New() *Engine {
	e := &Engine{router: newRouter()}
	e.RouterGroup = &RouterGroup{engine: e}
	e.groups = []*RouterGroup{e.RouterGroup}

	return e
}

func (e *Engine) Run(addr string) error {
	handlers := func(ctx *fasthttp.RequestCtx) {
		c := newContext(ctx)
		method := util.B2s(ctx.Method())
		fmt.Println(string(ctx.Path()))
		node, params := e.router.findRoute(method, util.B2s(ctx.Path()))
		// todo:buffer拼凑减少小对象
		if node == nil {
			c.Data(fasthttp.StatusNotFound, "", []byte("not found"))

			return
		}

		key := method + "-" + node.Pattern()
		c.Params = params

		if handler, ok := e.router.handlers[key]; ok {
			handler(c)
		} else {
			c.Data(fasthttp.StatusNotFound, "", []byte("not found"))
		}
	}

	return fasthttp.ListenAndServe(addr, handlers)
}
