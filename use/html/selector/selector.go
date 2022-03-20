/*
	Package selector provides the HTML <select> element.

	The <select> HTML element represents a control that provides a menu of options.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/select
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package selector

import (
	"qlova.tech/new/tree"
	"qlova.tech/use/html"
)

// Tag is the html <select> tag.
const Tag = html.Tag("select")

// New returns an html <select> tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}
