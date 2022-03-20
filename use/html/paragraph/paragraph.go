/*
	Package paragraph provides the HTML <p> element.

	The <p> HTML element represents a paragraph. Paragraphs are
	usually represented in visual media as blocks of text separated
	from adjacent blocks by blank lines and/or first-line
	indentation, but HTML paragraphs can be any structural grouping
	of related content, such as images or form fields.

	Paragraphs are block-level elements, and notably will
	automatically close if another block-level element is parsed
	before the closing </p> tag. See "Tag omission" below.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/p
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package paragraph

import (
	"qlova.tech/new/tree"
	"qlova.tech/use/html"
)

// Tag is the html <p> tag.
const Tag = html.Tag("p")

// New returns a <p> html tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}
