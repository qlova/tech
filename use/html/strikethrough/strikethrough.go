/*
	Package strikethrough provides the HTML <s> element.

	The <s> HTML element renders text with a strikethrough,
	or a line through it. Use the <s> element to represent
	things that are no longer relevant or no longer accurate.
	However, <s> is not appropriate when indicating document
	edits; for that, use the <del> and <ins> elements, as
	appropriate.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/s
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package strikethrough

import (
	"qlova.tech/use/html"
	"qlova.tech/web/tree"
)

// Tag is the html <s> tag.
const Tag = html.Tag("s")

// New returns a <s> html tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, html.Tag("s"))...)
}
