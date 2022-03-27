/*
	Package listitem provides the HTML <li> element.

	The <li> HTML element is used to represent an item in a list.
	It must be contained in a parent element: an ordered list
	(<ol>), an unordered list (<ul>), or a menu (<menu>). In
	menus and unordered lists, list items are usually displayed
	using bullet points. In ordered lists, they are usually
	displayed with an ascending counter on the left, such as a
	number or letter.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/li
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package listitem

import (
	"fmt"

	"qlova.tech/use/html"
	"qlova.tech/web/tree"
)

// Tag is the html <li> tag.
const Tag = html.Tag("li")

// New returns a <li> html tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}

// Value indicates the current ordinal value of the list item as
// defined by the <ol> element.
func Value(number uint) html.Attribute {
	return html.Attr("value", fmt.Sprint(number))
}
