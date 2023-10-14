package fgin

import (
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

// func (e *Engine) GET(pattern string, handlerFunc HandlerFunc) {
// 	e.router.addRoute("GET", pattern, handlerFunc)
// }
// func (e *Engine) POST(pattern string, handlerFunc HandlerFunc) {
// 	e.router.addRoute("POST", pattern, handlerFunc)
// }
// func (e *Engine) UPDATE(pattern string, handlerFunc HandlerFunc) {
// 	e.router.addRoute("UPDATE", pattern, handlerFunc)
// }
// func (e *Engine) DELETE(pattern string, handlerFunc HandlerFunc) {
// 	e.router.addRoute("DELETE", pattern, handlerFunc)
// }

func (e *Engine) Run(addr string) error {
	handlers := func(ctx *fasthttp.RequestCtx) {
		c := newContext(ctx)
		method := util.B2s(ctx.Method())
		node, params := e.router.findRoute(method, util.B2s(ctx.Path()))
		// todo:buffer拼凑减少小对象
		key := method + "-" + node.Pattern()
		c.Params = params

		if handler, ok := e.router.handlers[key]; ok {
			handler(c)
		} else {
			ctx.Error("not found", fasthttp.StatusNotFound)
		}
	}

	return fasthttp.ListenAndServe(addr, handlers)
}
