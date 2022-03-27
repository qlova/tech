/*
	Package wordbreak provides the HTML <wbr> element.

	The <wbr> HTML element represents a word break opportunity—a
	position within text where the browser may optionally break
	a line, though its line-breaking rules would not otherwise
	create a break at that location.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/wbr
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package wordbreak

import (
	"qlova.tech/use/html"
	"qlova.tech/web/tree"
)

// Tag is the html <wbr> tag.
const Tag = html.Tag("wbr")

// New returns a word break opportunity.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag, html.Void)...)
}
