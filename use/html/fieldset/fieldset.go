/*
	Package fieldset provides the HTML <fieldset> element.

	The <fieldset> HTML element is used to group several controls as
	well as labels (<label>) within a web form.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/fieldset
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package fieldset

import (
	"qlova.tech/use/html"
	"qlova.tech/web/tree"
)

// Tag is the html <fieldset> tag.
const Tag = html.Tag("fieldset")

// New returns an html <fieldset> tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}

// Name associated with the group.
type Name string

// RenderAttr implements the attributes.Renderer interface.
func (n Name) RenderAttr() []byte {
	return []byte(html.Attr("name", string(n)))
}

/*
	If this Boolean attribute is set, all form controls that a
	re descendants of the <fieldset>, are disabled, meaning
	they are not editable and won't be submitted along with
	the <form>. They won't receive any browsing events, like
	mouse clicks or focus-related events. By default browsers
	display such controls grayed out. Note that form elements
	inside the <legend> element won't be disabled.
*/
const Disabled = html.Attribute("disabled")

/*
	This attribute takes the value of the id attribute of a
	<form> element you want the <fieldset> to be part of,
	even if it is not inside the form. Please note that usage
	of this is confusing — if you want the <input> elements
	inside the <fieldset> to be associated with the form, you
	need to use the form attribute directly on those elements.
	You can check which elements are associated with a form
	via JavaScript, using HTMLFormElement.elements.
*/
func Form(id html.ID) html.Attribute {
	return html.Attr("form", string(id))
}
