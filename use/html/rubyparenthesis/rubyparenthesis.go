/*
	Package rubyparenthesis provides the HTML <rp> element.

	The <rp> HTML element is used to provide fall-back parentheses
	for browsers that do not support display of ruby annotations
	using the <ruby> element. One <rp> element should enclose each
	of the opening and closing parentheses that wrap the <rt>
	element that contains the annotation's text.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/rp
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package rubyparenthesis

import (
	"qlova.tech/use/html"
	"qlova.tech/web/tree"
)

// Tag is the html <rp> tag.
const Tag = html.Tag("rp")

// New returns an html <rp> tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, html.Tag("rp"))...)
}
