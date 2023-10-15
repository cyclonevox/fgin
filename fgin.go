package fgin

import (
	"strings"

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

func Default() *Engine {
	e := &Engine{router: newRouter()}
	e.RouterGroup = &RouterGroup{engine: e}
	e.Use(Logger(), Recovery())
	e.groups = []*RouterGroup{e.RouterGroup}

	return e
}

func (e *Engine) Run(addr string) error {
	handlers := func(ctx *fasthttp.RequestCtx) {
		c := newContext(ctx)

		var middlewares []HandlerFunc
		for _, group := range e.groups {
			if strings.HasPrefix(c.Path, group.prefix) {
				middlewares = append(middlewares, group.middlewares...)
			}
		}
		c.handlers = middlewares

		e.router.handle(c)
	}

	return fasthttp.ListenAndServe(addr, handlers)
}
