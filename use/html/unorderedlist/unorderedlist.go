/*
	Package unorderedlist provides the HTML <ul> element.

	The <ul> HTML element represents an unordered list
	of items, typically rendered as a bulleted list.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ul
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package unorderedlist

import (
	"qlova.tech/use/html"
	"qlova.tech/web/tree"
)

// Tag is the html <ul> tag.
const Tag = html.Tag("ul")

// New returns a <ul> html tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}
