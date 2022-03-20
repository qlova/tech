/*
	Package header provides the HTML <header> element.

	The <header> HTML element represents introductory content, typically
	a group of introductory or navigational aids. It may contain some
	heading elements but also a logo, a search form, an author name,
	and other elements.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/head
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package footer

import (
	"qlova.tech/new/tree"
	"qlova.tech/use/html"
)

// Tag is the html <header> tag.
const Tag = html.Tag("footer")

// New returns a <header> html tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}
