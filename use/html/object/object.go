/*
	Package object provides the HTML <object> element.

	The <object> HTML element represents an external resource, which
	can be treated as an image, a nested browsing context, or a
	resource to be handled by a plugin.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/object
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package object

import (
	"fmt"

	"qlova.tech/new/tree"
	"qlova.tech/use/html"
)

// Tag is the html <object> tag.
const Tag = html.Tag("object")

// New returns an html <object> tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}

// Form element, if any, that the object element is associated with
// (its form owner). The value of the attribute must be an ID of a
// <form> element in the same document.
func Form(id html.ID) tree.Node {
	return New(html.ID(id))
}

// Width of the object.
type Width uint

// RenderAttr implements the attr.Renderer interface.
func (w Width) RenderAttr() []byte {
	return []byte(html.Attr("width", fmt.Sprint(w)))
}

// Height of the object.
type Height uint

// RenderAttr implements the attr.Renderer interface.
func (h Height) RenderAttr() []byte {
	return []byte(html.Attr("height", fmt.Sprint(h)))
}

// Name of valid browsing context (HTML5), or the name
// of the control (HTML 4).
type Name string

// RenderAttr implements the attr.Renderer interface.
func (n Name) RenderAttr() []byte {
	return []byte(html.Attr("name", string(n)))
}

// MimeType to use to select the plug-in to instantiate.
func MimeType(mimetype string) html.Attribute {
	return html.Attr("type", mimetype)
}

// Map is the partial URL (starting with #) of an image map associated
// with the element.
type Map string

// RenderAttr implements the attr.Renderer interface.
func (m Map) RenderAttr() []byte {
	return []byte(html.Attr("usemap", string(m)))
}
