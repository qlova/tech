/*
	Package foot provides the HTML <tfoot> element.

	The <tfoot> HTML element defines a set of rows
	summarizing the columns of the table.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/tfoot
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package foot

import (
	"qlova.tech/new/tree"
	"qlova.tech/use/html"
)

// Tag is the html <tfoot> tag.
const Tag = html.Tag("tfoot")

// New returns an html <tfoot> tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}
