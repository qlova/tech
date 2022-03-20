/*
	Package unarticulated provides the HTML <u> element.

	The <u> HTML element represents a span of inline text which
	should be rendered in a way that indicates that it has a
	non-textual annotation. This is rendered by default as a
	simple solid underline, but may be altered using CSS.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/u
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package unarticulated

import (
	"qlova.tech/new/tree"
	"qlova.tech/use/html"
)

// Tag is the html <u> tag.
const Tag = html.Tag("u")

// New returns an html <u> tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}
