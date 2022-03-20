/*
	Package aside provides the HTML <aside> element.

	The <aside> HTML element represents a portion of a document whose
	content is only indirectly related to the document's main content.
	Asides are frequently presented as sidebars or call-out boxes.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/aside
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package aside

import (
	"qlova.tech/new/tree"
	"qlova.tech/use/html"
)

// Tag is the html <aside> tag.
const Tag = html.Tag("aside")

// New returns an html <aside> tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}
