/*
	Package section provides the HTML <section> element.

	The <section> HTML element represents a generic standalone section
	of a document, which doesn't have a more specific semantic element
	to represent it. Sections should always have a heading, with very
	few exceptions.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/section
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package section

import (
	"qlova.tech/use/html"
	"qlova.tech/web/tree"
)

// Tag is the html <section> tag.
const Tag = html.Tag("section")

// New returns a <section> html tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}
