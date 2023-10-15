package fgin

import (
	"encoding/json"

	"git.vox666.top/vox/fgin/internal/util"
	"github.com/valyala/fasthttp"
)

type Context struct {
	ctx    *fasthttp.RequestCtx
	Params map[string]string
}

func newContext(ctx *fasthttp.RequestCtx) *Context {
	return &Context{
		ctx: ctx,
	}
}

func (c *Context) Method() string {
	return util.B2s(c.ctx.Method())
}

func (c *Context) Path() string {
	return util.B2s(c.ctx.Path())
}

func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

func (c *Context) PostForm(key string) string {
	return util.B2s(c.ctx.FormValue(key))
}

func (c *Context) Query() string {
	return util.B2s(c.ctx.URI().QueryString())
}

func (c *Context) Status(code int) {
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
