/*
	Package code provides the HTML <code> element.

	The <code> HTML element displays its contents styled in a
	fashion intended to indicate that the text is a short
	fragment of computer code. By default, the content text
	is displayed using the user agent's default monospace font.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/code
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package code

import (
	"qlova.tech/new/tree"
	"qlova.tech/use/html"
)

// Tag is the html <code> tag.
const Tag = html.Tag("code")

// New returns an html <code> tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}
