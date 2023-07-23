// Package show provides user interface tree nodes.
package show

import (
	"qlova.tech/app/data"
	"qlova.tech/app/then"
)

type Node interface {
	node()
}

type Layout struct {
	isNode

	Nodes []Node
	Hints struct {
		Row bool
	}
}

type isNode = Node

type String struct {
	isNode

	Value string
	Hints struct {
		Title bool
	}
}

type pickable interface {
	~string
}

type Picker[T pickable] struct {
	isNode

	Value *T
	Hints struct {
	}
}

type Choice struct {
	isNode

	Steps then.Steps
	Hints struct {
		Button  string
		OnClick bool
		OnTouch bool
		OnHover bool
	}
}

type Viewer[T any] struct {
	isNode

	Value *T
	Hints struct {
	}
}

type Looped struct {
	isNode

	Value any
	Loops func(data.Sync) (data.Sync, Layout)
	Hints struct {
	}
}
