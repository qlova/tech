/*
	Package menu provides the HTML <menu> element.

	The <menu> HTML element is a semantic alternative to <ul>. It
	represents an unordered list of items (represented by <li>
	elements), each of these represent a link or other command that
	the user can activate.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/menu
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package menu

import (
	"qlova.tech/new/tree"
	"qlova.tech/use/html"
)

// Tag is the html <menu> tag.
const Tag = html.Tag("menu")

// New returns a <menu> html tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}
