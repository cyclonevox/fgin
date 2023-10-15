package fgin

import (
	"encoding/json"

	"git.vox666.top/vox/fgin/internal/util"
	"github.com/valyala/fasthttp"
)

type Request struct {
	RequestURI []byte
}

type Context struct {
	ctx *fasthttp.RequestCtx

	// todo 切换引擎后可以不再使用fasthttp的Context，减少开销。
	//  	此处的目的是需要像gin使用context字段达成效果
	//      使用Request对象的指针来达到效果
	Method     string
	Path       string
	StatusCode int
	Request
	// todo gin使用了slice进一步减少消耗，后续可以修改
	Params

	handlers []HandlerFunc
	index    int
}

func newContext(ctx *fasthttp.RequestCtx) *Context {
	return &Context{
		ctx:     ctx,
		Method:  util.B2s(ctx.Method()),
		Path:    util.B2s(ctx.Path()),
		Request: Request{RequestURI: ctx.RequestURI()},
		index:   -1,
	}
}

func (c *Context) Next() {
	c.index++
	s := len(c.handlers)
	for ; c.index < s; c.index++ {
		c.handlers[c.index](c)
	}
}

func (c *Context) Param(key string) string {
	return c.Params.ByName(key)
}

func (c *Context) PostForm(key string) string {
	return util.B2s(c.ctx.FormValue(key))
}

func (c *Context) Query() string {
	return util.B2s(c.ctx.URI().QueryString())
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.ctx.SetStatusCode(code)
}

func (c *Context) SetHeader(key string, value string) {
	c.ctx.Request.Header.Set(key, value)
}

func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	if encode, err := json.Marshal(obj); err != nil {
		c.ctx.Error(err.Error(), 500)
	} else {
		c.ctx.SetBody(encode)
	}
}

func (c *Context) Data(code int, contentType string, data []byte) {
	c.Status(code)
	c.ctx.SetBody(data)
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.ctx.SetBody([]byte(html))
}
