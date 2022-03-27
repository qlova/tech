/*
	Package tbody provides the HTML <tbody> element.

	The <tbody> HTML element encapsulates a set of table
	rows (<tr> elements), indicating that they comprise
	the body of the table (<table>).

	The <tbody> element, along with its cousins <thead>
	and <tfoot>, provide useful semantic information
	that can be used when rendering for either screen
	or printer as well as for accessibility purposes.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/tbody
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package body

import (
	"qlova.tech/use/html"
	"qlova.tech/web/tree"
)

// Tag is the html <tbody> tag.
const Tag = html.Tag("tbody")

// New returns an html <tbody> tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}
