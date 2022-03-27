/*
	Package group provides the HTML <colgroup> element.

	The <colgroup> HTML element defines a group of columns
	within a table.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/colgroup
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package group

import (
	"fmt"

	"qlova.tech/use/html"
	"qlova.tech/web/tree"
)

// Tag is the html <colgroup> tag.
const Tag = html.Tag("colgroup")

// New returns a <colgroup> html tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}

// Span contains a positive integer indicating the number of
// consecutive columns the <colgroup> element spans. If not
// present, its default value is 1.
//
// This attribute is applied on the attributes of the column
// group, it has no effect on the CSS styling rules associated
// with it or, even more, to the cells of the column's members
// of the group.
//
// The span attribute is not permitted if there are one or more
// <col> elements within the <colgroup>.
type Span uint

// RenderAttr implements the attributes.Renderer interface.
func (s Span) RenderAttr() []byte {
	return []byte(html.Attr("span", fmt.Sprint(s)))
}
