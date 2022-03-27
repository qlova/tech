/*
	Package descriptiondetails provides the HTML <dd> element.

	The <dd> HTML element provides the description, definition, or
	value for the preceding term (<dt>) in a description list (<dl>).

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dd
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package descriptiondetails

import (
	"qlova.tech/use/html"
	"qlova.tech/web/tree"
)

// Tag is the html <dd> tag.
const Tag = html.Tag("dd")

// New returns a <dd> html tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}
