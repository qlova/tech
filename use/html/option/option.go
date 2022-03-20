/*
	Package option provides the HTML <option> element.

	The <option> HTML element is used to define an item
	contained in a <select>, an <optgroup>, or a <datalist> element.
	As such, <option> can represent menu items in popups and
	other lists of items in an HTML document.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/option
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package option

import (
	"qlova.tech/new/tree"
	"qlova.tech/use/html"
)

// Tag is the html <option> tag.
const Tag = html.Tag("option")

// New returns an html <option> tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag, html.Void)...)
}

// Disabled when set, means this option is not checkable. Often
// browsers grey out such control and it won't receive any
// browsing event, like mouse clicks or focus-related ones.
// If this attribute is not set, the element can still be
// disabled if one of its ancestors is a disabled <optgroup> element.
const Disabled = html.Attribute("disabled")

// Label is text for the label indicating the meaning of the option.
// If the label attribute isn't defined, its value is that of the
// element text content.
type Label string

// RenderAttr implements the attributes.Renderer interface.
func (l Label) RenderAttr() []byte {
	return []byte(html.Attr("label", string(l)))
}

/*
	Selected if present, this Boolean attribute indicates that the
	option is initially selected. If the <option> element is the
	descendant of a <select> element whose multiple attribute is
	not set, only one single <option> of this <select> element may
	have the selected attribute.
*/
const Selected = html.Attribute("selected")

/*
	Value to be submitted with the form, should this option be
	selected. If this attribute is omitted, the value is taken
	from the text content of the option element.
*/
type Value string

// RenderAttr implements the attributes.Renderer interface.
func (v Value) RenderAttr() []byte {
	return []byte(html.Attr("value", string(v)))
}
