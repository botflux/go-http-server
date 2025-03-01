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

	if isPathEmpty && h.isLeaf() {
		return h
	}

	if h.isLeaf() || isPathEmpty {
		return nil
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
	radix *httpRadixTree
}

func NewRouter() *Router {
	r := newRadix()

	return &Router{
		radix: r,
	}
}

func (router *Router) Add(url string) {
	path := strings.Split(url, "/")

	router.radix.Insert(path)
}

func (router *Router) Dispatch(url string) bool {
	path := strings.Split(url, "/")

	node := router.radix.Search(path)

	return node != nil
}
