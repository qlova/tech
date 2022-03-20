/*
	Package figcaption provides the HTML <figcaption> element.

	The <figcaption> HTML element represents a caption or legend
	describing the rest of the contents of its parent <figure> element.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/figcaption
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package figcaption

import (
	"qlova.tech/new/tree"
	"qlova.tech/use/html"
)

// Tag is the html <figcaption> tag.
const Tag = html.Tag("figcaption")

// New returns a <figcaption> html tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}
