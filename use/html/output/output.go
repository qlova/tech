/*
	Package output provides the HTML <output> element.

	The <output> HTML element is a container element into which
	a site or app can inject the results of a calculation or
	the outcome of a user action.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/output
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package output

import (
	"strings"

	"qlova.tech/use/html"
	"qlova.tech/web/tree"
)

// Tag is the html <output> tag.
const Tag = html.Tag("output")

// New returns an html <output> tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}

// For list of other elements' ids, indicating that those
// elements contributed input values to (or otherwise
// affected) the calculation.
func For(ids ...html.ID) html.Attribute {
	var s []string
	for _, id := range ids {
		s = append(s, string(id))
	}
	return html.Attr("for", strings.Join(s, " "))
}

// The <form> element to associate the output with (its form owner).
// The value of this attribute must be the id of a <form> in the same
// document. (If this attribute is not set, the <output> is associated
// with its ancestor <form> element, if any.)
//
// This attribute lets you associate <output> elements to <form>s
// anywhere in the document, not just inside a <form>. It can also
// override an ancestor <form> element.
func Form(id html.ID) html.Attribute {
	return html.Attr("form", string(id))
}

// Name of the output.
type Name string

// RenderAttr implements the attributes.Renderer interface.
func (n Name) RenderAttr() []byte {
	return []byte(html.Attr("name", string(n)))
}
