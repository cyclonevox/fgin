package fgin

type group struct {
	prefix      string
	middlewares []HandlerFunc
	parent      *group
	engine      *Engine
}

func (g *group) Group(prefix string) *group {
	newGroup := &group{
		prefix: g.prefix + prefix,
		parent: g.parent,
		engine: g.engine,
	}

	g.engine.groups = append(g.engine.groups, newGroup)
	return newGroup
}

func (g *group) addRoute(method, comp string, handler HandlerFunc) {
	pattern := g.prefix + comp
	g.engine.router.addRoute(method, pattern, handler)
}

func (g *group) GET(pattern string, handlerFunc HandlerFunc) {
	g.addRoute("GET", pattern, handlerFunc)
}
func (g *group) POST(pattern string, handlerFunc HandlerFunc) {
	g.addRoute("POST", pattern, handlerFunc)
}
func (g *group) UPDATE(pattern string, handlerFunc HandlerFunc) {
	g.addRoute("UPDATE", pattern, handlerFunc)
}
func (g *group) DELETE(pattern string, handlerFunc HandlerFunc) {
	g.addRoute("DELETE", pattern, handlerFunc)
}
