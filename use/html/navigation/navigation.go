/*
	Package navigation provides the HTML <nav> element.

	The <nav> HTML element represents a section of a page whose purpose
	is to provide navigation links, either within the current document
	or to other documents. Common examples of navigation sections are
	menus, tables of contents, and indexes.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/nav
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package navigation

import (
	"qlova.tech/new/tree"
	"qlova.tech/use/html"
)

// Tag is the html <nav> tag.
const Tag = html.Tag("nav")

// New returns a <nav> html tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}
