package fgin

import (
	"strings"

	"git.vox666.top/vox/fgin/internal/trie"
)

type router struct {
	handlers map[string]HandlerFunc
	roots    map[string]*trie.Node
}

func newRouter() router {
	return router{
		handlers: make(map[string]HandlerFunc),
		roots:    make(map[string]*trie.Node),
	}
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)

	if _, ok := r.roots[method]; !ok {
		r.roots[method] = &trie.Node{}
	}
	r.roots[method].Insert(pattern, parts, 0)

	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *router) findRoute(method string, path string) (*trie.Node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}

	n := root.Search(searchParts, 0)
	if n != nil {
		parts := parsePattern(n.Pattern())
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}

	return nil, nil
}

func parsePattern(pattern string) []string {
	ps := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, p := range ps {
		if p != "" {
			parts = append(parts, p)
			// 有星号则直接不用匹配了。
			if p[0] == '*' {
				break
			}
		}
	}

	return parts
}
