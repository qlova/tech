/*
	Package td provides the HTML <td> element.

	The <td> HTML element defines a cell of a table
	that contains data. It participates in the table model.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/td
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package data

import (
	"fmt"
	"strings"

	"qlova.tech/new/tree"
	"qlova.tech/use/html"
)

// Tag is the html <td> tag.
const Tag = html.Tag("td")

// New returns an html <td> tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}

// Headers contains a list of space-separated strings, each
// corresponding to the id attribute of the <th> elements that
// apply to this element.
type Headers []html.ID

// RenderAttr implements the attributes.Renderer interface.
func (h Headers) RenderAttr() []byte {
	var s []string
	for _, h := range h {
		s = append(s, string(h))
	}
	return []byte(html.Attr("headers", strings.Join(s, " ")))
}

// ColumnSpan contains a non-negative integer value that indicates
// for how many columns the cell extends. Its default value is 1.
// Values higher than 1000 will be considered as incorrect and will
// be set to the default value (1).
type ColumnSpan uint

// RenderAttr implements the attributes.Renderer interface.
func (s ColumnSpan) RenderAttr() []byte {
	return []byte(html.Attr("colspan", fmt.Sprint(s)))
}

// RowSpan contains a non-negative integer value that indicates for
// how many rows the cell extends. Its default value is 1; if its
// value is set to 0, it extends until the end of the table section
// (<thead>, <tbody>, <tfoot>, even if implicitly defined), that the
// cell belongs to. Values higher than 65534 are clipped down to 65534.
type RowSpan uint

// RenderAttr implements the attributes.Renderer interface.
func (s RowSpan) RenderAttr() []byte {
	return []byte(html.Attr("rowspan", fmt.Sprint(s)))
}
