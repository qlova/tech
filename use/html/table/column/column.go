/*
	Package column provides the HTML <col> element.

	The <col> HTML element defines a column within a table
	and is used for defining common semantics on all common
	cells. It is generally found within a <colgroup> element.

	<col> allows styling columns using CSS, but only a few
	properties will have an effect on the column (see
	the CSS 2.1 specification for a list).

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/col
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package column

import (
	"fmt"

	"qlova.tech/use/html"
	"qlova.tech/web/tree"
)

// Tag is the html <col> tag.
const Tag = html.Tag("col")

// New returns an html <col> tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}

// Span contains a positive integer indicating the number of
// consecutive columns the <col> element spans. If not present,
// its default value is 1.
type Span uint

// RenderAttr implements the attributes.Renderer interface.
func (s Span) RenderAttr() []byte {
	return []byte(html.Attr("span", fmt.Sprint(s)))
}
