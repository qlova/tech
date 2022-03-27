/*
	Package legend provides the HTML <legend> element.

	The <legend> HTML element represents a caption
	for the content of its parent <fieldset>.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/legend
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package legend

import (
	"qlova.tech/use/html"
	"qlova.tech/web/tree"
)

// Tag is the html <legend> tag.
const Tag = html.Tag("legend")

// New returns an html <legend> tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}
