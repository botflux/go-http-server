package routing

import "strings"

type httpRadixTree struct {
	path      string
	nextPaths []*httpRadixTree
}

func newRadix() *httpRadixTree {
	return &httpRadixTree{
		path:      "",
		nextPaths: []*httpRadixTree{},
	}
}

func (h *httpRadixTree) isLeaf() bool {
	return len(h.nextPaths) == 0
}

func (h *httpRadixTree) isPrefix(path []string) bool {
	return len(path) > 0 && path[0] == h.path
}

func (h *httpRadixTree) Search(path []string) *httpRadixTree {
	isPathEmpty := len(path) == 0

	if isPathEmpty {
		return h
	}

	for _, next := range h.nextPaths {
		if next.isPrefix(path) {
			if result := next.Search(path[1:]); result != nil {
				return result
			}
		}
	}

	return nil
}

func (h *httpRadixTree) Insert(path []string) {
	for _, next := range h.nextPaths {
		if next.isPrefix(path) {
			next.Insert(path[1:])
			return
		}
	}

	curr := h

	for _, chunk := range path {
		newNode := &httpRadixTree{
			path:      chunk,
			nextPaths: []*httpRadixTree{},
		}

		curr.nextPaths = append(curr.nextPaths, newNode)
		curr = newNode
	}
}

type Router struct {
	radixPerMethod map[string]*httpRadixTree
}

func NewRouter() *Router {
	return &Router{
		radixPerMethod: map[string]*httpRadixTree{},
	}
}

func (router *Router) Add(method string, url string) {
	path := strings.Split(url, "/")

	if router.radixPerMethod[method] == nil {
		router.radixPerMethod[method] = newRadix()
	}

	router.radixPerMethod[method].Insert(path)
}

func (router *Router) Dispatch(method string, url string) bool {
	path := strings.Split(url, "/")

	if router.radixPerMethod[method] == nil {
		return false
	}

	node := router.radixPerMethod[method].Search(path)

	return node != nil
}
