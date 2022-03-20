/*
	Package data provides the HTML <data> element.

	The <data> HTML element links a given piece of content
	with a machine-readable translation. If the content is
	time- or date-related, the <time> element must be used.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/data
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package data

import (
	"qlova.tech/new/tree"
	"qlova.tech/use/html"
)

// Tag is the html <data> tag.
const Tag = html.Tag("data")

// New returns a <data> html tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}

// Value pecifies the machine-readable translation of the
// content of the element.
func Value(value string) html.Attribute {
	return html.Attr("value", value)
}
