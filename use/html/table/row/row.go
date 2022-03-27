/*
	Package row provides the HTML <tr> element.

	The <tr> HTML element defines a row of cells in a table.
	The row's cells can then be established using a mix of
	<td> (data cell) and <th> (header cell) elements.

	To provide additional control over how cells fit into
	(or span across) columns, both <th> and <td> support
	the colspan attribute, which lets you specify how many
	columns wide the cell should be, with the default being 1.
	Similarly, you can use the rowspan attribute on cells to
	indicate they should span more than one table row.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/tr
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package row

import (
	"qlova.tech/use/html"
	"qlova.tech/web/tree"
)

// Tag is the html <tr> tag.
const Tag = html.Tag("tr")

// New returns an html <tr> tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}
