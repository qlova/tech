/*
	Package subscript provides the HTML <sub> element.

	The <sub> HTML element specifies inline text which should
	be displayed as subscript for solely typographical reasons.
	Subscripts are typically rendered with a lowered baseline
	using smaller text.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/sub
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package subscript

import (
	"qlova.tech/new/tree"
	"qlova.tech/use/html"
)

// Tag is the html <sub> tag.
const Tag = html.Tag("sub")

// New returns an html <sub> tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}
