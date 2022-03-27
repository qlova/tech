/*
	Package slot provides the HTML <slot> element.

	The <slot> HTML element—part of the Web Components technology
	suite—is a placeholder inside a web component that you can
	fill with your own markup, which lets you create separate
	\DOM trees and present them together.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/slot
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package slot

import (
	"qlova.tech/use/html"
	"qlova.tech/web/tree"
)

// Tag is the html <slot> tag.
const Tag = html.Tag("slot")

// New returns an html <slot> tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}

// Name of the slot.
type Name string

// RenderAttr implements the attributes.Renderer interface.
func (n Name) RenderAttr() []byte {
	return []byte(html.Attr("name", string(n)))
}
