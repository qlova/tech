/*
	Package datalist provides the HTML <datalist> element.

	The <datalist> HTML element contains a set of <option> elements
	that represent the permissible or recommended options available
	to choose from within other controls.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/datalist
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package datalist

import (
	"qlova.tech/new/tree"
	"qlova.tech/use/html"
)

// Tag is the html <datalist> tag.
const Tag = html.Tag("datalist")

// New returns a <datalist> html tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}
