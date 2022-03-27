/*
	Package mark provides the HTML <mark> element.

	The <mark> HTML element represents text which is marked or
	highlighted for reference or notation purposes, due to the
	marked passage's relevance or importance in the enclosing context.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/mark
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package mark

import (
	"qlova.tech/use/html"
	"qlova.tech/web/tree"
)

// Tag is the html <mark> tag.
const Tag = html.Tag("mark")

// New returns an html <mark> tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}
