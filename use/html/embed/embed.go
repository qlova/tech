/*
	Package embed provides the HTML <embed> element.

	The <embed> HTML element embeds external content at the
	specified point in the document. This content is provided
	by an external application or other source of interactive
	content such as a browser plug-in.

	Keep in mind that most modern browsers have deprecated and
	removed support for browser plug-ins, so relying upon
	<embed> is generally not wise if you want your site to be
	operable on the average user's browser.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/embed
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package embed

import (
	"fmt"

	"qlova.tech/use/html"
	"qlova.tech/web/tree"
)

// Tag is the html <embed> tag.
const Tag = html.Tag("embed")

// New returns a new <embed> element.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag, html.Void)...)
}

// Width of the resource.
type Width uint

// RenderAttr implements the attr.Renderer interface.
func (w Width) RenderAttr() []byte {
	return []byte(html.Attr("width", fmt.Sprint(w)))
}

// Height of the resource.
type Height uint

// RenderAttr implements the attr.Renderer interface.
func (h Height) RenderAttr() []byte {
	return []byte(html.Attr("height", fmt.Sprint(h)))
}

// Source URL of the resource being embedded.
type Source string

// RenderAttr implements the attr.Renderer interface.
func (s Source) RenderAttr() []byte {
	return []byte(html.Attr("src", string(s)))
}

// MimeType to use to select the plug-in to instantiate.
func MimeType(mimetype string) html.Attribute {
	return html.Attr("type", mimetype)
}
