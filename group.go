package fgin

type Group struct {
	prefix      string
	middlewares []HandlerFunc
	parent      *Group
	engine      *Engine
}

func (g *Group) Group(prefix string) *Group {
	newGroup := &Group{
		prefix: g.prefix + prefix,
		parent: g.parent,
		engine: g.engine,
	}

	g.engine.groups = append(g.engine.groups, newGroup)
	return newGroup
}

func (g *Group) addRoute(method, comp string, handler HandlerFunc) {
	pattern := g.prefix + comp
	g.engine.router.addRoute(method, pattern, handler)
}

func (g *Group) GET(pattern string, handlerFunc HandlerFunc) {
	g.addRoute("GET", pattern, handlerFunc)
}
func (g *Group) POST(pattern string, handlerFunc HandlerFunc) {
	g.addRoute("POST", pattern, handlerFunc)
}
func (g *Group) UPDATE(pattern string, handlerFunc HandlerFunc) {
	g.addRoute("UPDATE", pattern, handlerFunc)
}
func (g *Group) DELETE(pattern string, handlerFunc HandlerFunc) {
	g.addRoute("DELETE", pattern, handlerFunc)
}
