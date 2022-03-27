package tree

import "reflect"

type Node []any

type Seed struct {
	*seed
}

type seed struct {
	Body struct {
		Suffix []any
	}

	TreeRenderers map[reflect.Type]any
}

func NewSeed() Seed {
	return Seed{&seed{
		TreeRenderers: make(map[reflect.Type]any),
	}}
}

type Renderer interface {
	RenderTree(Seed) Node
}

func New(args ...any) Node {
	return args
}

func render(tree Node, seed Seed) Node {
	for i, node := range tree {
		switch node := node.(type) {
		case Renderer:
			tree[i] = node.RenderTree(seed)
		case Node:
			tree[i] = render(node, seed)
		}
	}

	return tree
}

func Render(tree Renderer, seed Seed) Node {
	return render(tree.RenderTree(seed), seed)
}
