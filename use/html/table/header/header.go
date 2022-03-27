/*
	Package header provides the HTML <th> element.

	The <th> HTML element defines a cell as header of a
	group of table cells. The exact nature of this group
	is defined by the scope and headers attributes.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/th
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package header

import (
	"fmt"
	"strings"

	"qlova.tech/use/html"
	"qlova.tech/web/tree"
)

// Tag is the html <th> tag.
const Tag = html.Tag("th")

// New returns an html <th> tree node.
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

// Abbreviation contains a short abbreviated description of the
// cell's content. Some user-agents, such as speech readers, may
// present this description before the content itself.
type Abbreviation string

// RenderAttr implements the attributes.Renderer interface.
func (s Abbreviation) RenderAttr() []byte {
	return []byte(html.Attr("abbr", string(s)))
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

// Defines the cells that the header (defined in the <th>) element relates to
const (
	// The header relates to all cells of the row it belongs to.
	ScopeRow = html.Attribute("scope=row")

	// The header relates to all cells of the column it belongs to.
	ScopeColumn = html.Attribute("scope=col")

	// The header belongs to a rowgroup and relates to all of its cells.
	// These cells can be placed to the right or the left of the header,
	// depending on the value of the dir attribute in the <table> element.
	ScopeRowGroup = html.Attribute("scope=rowgroup")

	// The header belongs to a colgroup and relates to all of its cells.
	ScopeColumnGroup = html.Attribute("scope=colgroup")
)
