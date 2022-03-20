/*
	Package ruby provides the HTML <ruby> element.

	The <ruby> HTML element represents small annotations that are
	rendered above, below, or next to base text, usually used for
	showing the pronunciation of East Asian characters. It can also
	be used for annotating other kinds of text, but this usage is
	less common.

	The term ruby originated as a unit of measurement used by
	typesetters, representing the smallest size that text can be
	printed on newsprint while remaining legible.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ruby
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package ruby

import (
	"qlova.tech/new/tree"
	"qlova.tech/use/html"
)

// Tag is the html <ruby> tag.
const Tag = html.Tag("ruby")

// New returns an html <ruby> tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}
