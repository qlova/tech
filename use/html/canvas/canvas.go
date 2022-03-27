/*
	Package canvas provides the HTML <canvas> element.

	Use the HTML <canvas> element with either the canvas scripting
	API or the WebGL API to draw graphics and animations.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/canvas
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package canvas

import (
	"fmt"

	"qlova.tech/use/html"
	"qlova.tech/web/tree"
)

// Tag is the html <canvas> tag.
const Tag = html.Tag("canvas")

// New returns a <canvas> html tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}

// Width is the intrinsic width of the canvas in pixels.
type Width uint

// RenderAttr implements the attr.Renderer interface.
func (w Width) RenderAttr() []byte {
	return []byte(html.Attr("width", fmt.Sprint(w)))
}

// The height of the canvas in pixels.
type Height uint

// RenderAttr implements the attr.Renderer interface.
func (h Height) RenderAttr() []byte {
	return []byte(html.Attr("height", fmt.Sprint(h)))
}
