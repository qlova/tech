/*
	Package imagemap provides the HTML <map> element.

	The <map> HTML element is used with <area> elements
	to define an image map (a clickable link area).

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/map
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package imagemap

import (
	"qlova.tech/use/html"
	"qlova.tech/web/tree"
)

// Tag is the html <map> tag.
const Tag = html.Tag("map")

// New returns a <map> html tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}

// Name attribute gives the map a name so that it can be
// referenced. The attribute must be present and must have
// a non-empty value with no space characters. The value of
// the name attribute must not be equal to the value of the
// name attribute of another <map> element in the same
// document. If the id attribute is also specified, both
// attributes must have the same value.
type Name string

// RenderAttr implements the attr.Renderer interface.
func (n Name) RenderAttr() []byte {
	return []byte(html.Attr("name", string(n)))
}
