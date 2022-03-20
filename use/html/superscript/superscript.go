/*
	Package superscript provides the HTML <sup> element.

	The <sup> HTML element specifies inline text which is to
	be displayed as superscript for solely typographical reasons.
	Superscripts are usually rendered with a raised baseline
	using smaller text.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/sup
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package superscript

import (
	"qlova.tech/new/tree"
	"qlova.tech/use/html"
)

// Tag is the html <sup> tag.
const Tag = html.Tag("sup")

// New returns an html <sup> tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}
