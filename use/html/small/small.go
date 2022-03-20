/*
	Package small provides the HTML <small> element.

	The <small> HTML element represents side-comments and
	small print, like copyright and legal text, independent
	of its styled presentation. By default, it renders text
	within it one font-size smaller, such as from small to x-small.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/small
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package small

import (
	"qlova.tech/new/tree"
	"qlova.tech/use/html"
)

// Tag is the html <small> tag.
const Tag = html.Tag("small")

// New returns an html <small> tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}
