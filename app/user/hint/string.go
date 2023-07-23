// Package hint provides semantic presentation hints for user interfaces.
package hint

import (
	"qlova.tech/app/show"
)

type String interface {
	HintString(*show.String)
}

type hintStringFunc func(*show.String)

func (fn hintStringFunc) HintString(node *show.String) { fn(node) }

var Title title

type title struct{}

func (title) HintString(node *show.String) {
	node.Hints.Title = true
}
