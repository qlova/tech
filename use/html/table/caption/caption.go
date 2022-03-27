/*
	Package caption provides the HTML <caption> element.

	The <caption> HTML element specifies the caption (or title)
	of a table.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/caption
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package caption

import (
	"qlova.tech/use/html"
	"qlova.tech/web/tree"
)

// Tag is the html <caption> tag.
const Tag = html.Tag("caption")

// New returns an html <caption> tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}
