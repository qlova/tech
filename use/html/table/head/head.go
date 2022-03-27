/*
	Package head provides the HTML <thead> element.

	The <thead> HTML element defines a set of rows defining the
	head of the columns of the table.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/thead
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package head

import (
	"qlova.tech/use/html"
	"qlova.tech/web/tree"
)

// Tag is the html <thead> tag.
const Tag = html.Tag("thead")

// New returns an html <thead> tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}
