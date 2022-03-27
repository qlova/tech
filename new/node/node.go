package node

import "qlova.tech/web/tree"

func Get[T any](node tree.Node) T {
	var empty T
	for _, arg := range node {
		if v, ok := arg.(T); ok {
			return v
		}
	}
	return empty
}
