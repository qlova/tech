/*
	Package group provides the HTML <optgroup> element.

	The <optgroup> HTML element creates a grouping of options
	within a <select> element.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/optgroup
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package group

import (
	"qlova.tech/new/tree"
	"qlova.tech/use/html"
)

// Tag is the html <optgroup> tag.
const Tag = html.Tag("optgroup")

// New returns an html <optgroup> tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}

// Disabled when set, means none of the items in this option group is
// selectable. Often browsers grey out such control and it won't
// receive any browsing events, like mouse clicks or focus-related ones.
const Disabled = html.Attribute("disabled")

// Label is the name of the group of options, which the browser can use
// when labeling the options in the user interface. This attribute is
// mandatory if this element is used.
type Label string

// RenderAttr implements the attributes.Renderer interface.
func (l Label) RenderAttr() []byte {
	return []byte(html.Attr("label", string(l)))
}
