/*
	Package variable provides the HTML <var> element.

	The <var> HTML element represents the name of a variable in a
	mathematical expression or a programming context. It's
	typically presented using an italicized version of the
	current typeface, although that behavior is browser-dependent.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/var
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package variable

import (
	"qlova.tech/new/tree"
	"qlova.tech/use/html"
)

// Tag is the html <var> tag.
const Tag = html.Tag("var")

// New returns an html <var> tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}
