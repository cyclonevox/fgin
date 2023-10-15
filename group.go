package fgin

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	parent      *RouterGroup
	engine      *Engine
}

func (g *RouterGroup) Group(prefix string) *RouterGroup {
	newGroup := &RouterGroup{
		prefix: g.prefix + prefix,
		parent: g.parent,
		engine: g.engine,
	}

	g.engine.groups = append(g.engine.groups, newGroup)
	return newGroup
}

func (g *RouterGroup) addRoute(method, comp string, handler HandlerFunc) {
	pattern := g.prefix + comp
	g.engine.router.addRoute(method, pattern, handler)
}

func (g *RouterGroup) GET(pattern string, handlerFunc HandlerFunc) {
	g.addRoute("GET", pattern, handlerFunc)
}
func (g *RouterGroup) POST(pattern string, handlerFunc HandlerFunc) {
	g.addRoute("POST", pattern, handlerFunc)
}
func (g *RouterGroup) PUT(pattern string, handlerFunc HandlerFunc) {
	g.addRoute("PUT", pattern, handlerFunc)
}
func (g *RouterGroup) DELETE(pattern string, handlerFunc HandlerFunc) {
	g.addRoute("DELETE", pattern, handlerFunc)
}
