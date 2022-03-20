/*
	Package label provides the HTML <label> element.

	The <label> HTML element represents a caption for
	an item in a user interface.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/label
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package label

import (
	"qlova.tech/new/tree"
	"qlova.tech/use/html"
)

// Tag is the html <label> tag.
const Tag = html.Tag("label")

// New returns a <label> html tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}

/*
	The value of the for attribute must be a single id for a labelable
	form-related element in the same document as the <label> element.
	So, any given label element can be associated with only one form control.

	The first element in the document with an id attribute matching the
	value of the for attribute is the labeled control for this label
	element — if the element with that id is actually a labelable
	element. If it is not a labelable element, then the for attribute
	has no effect. If there are other elements that also match the id
	value, later in the document, they are not considered.

	Multiple label elements can be given the same value for their for
	attribute; doing so causes the associated form control (the form
	control that for value references) to have multiple labels.
*/
func For(input html.ID) html.Attribute {
	return html.Attr("for", string(input))
}
