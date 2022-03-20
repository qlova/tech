/*
	Package figure provides the HTML <figure> element.

	The <figure> HTML element represents self-contained content,
	potentially with an optional caption, which is specified using
	the <figcaption> element. The figure, its caption, and its
	contents are referenced as a single unit.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/figure
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package figure

import (
	"qlova.tech/new/tree"
	"qlova.tech/use/html"
)

// Tag is the html <figure> tag.
const Tag = html.Tag("figure")

// New returns a <figure> html tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}
