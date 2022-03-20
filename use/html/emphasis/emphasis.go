/*
	Package emphasis provides the HTML <em> element.

	The <em> HTML element marks text that has stress emphasis.
	The <em> element can be nested, with each level of nesting
	indicating a greater degree of emphasis.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/em
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package emphasis

import (
	"qlova.tech/new/tree"
	"qlova.tech/use/html"
)

// Tag is the html <em> tag.
const Tag = html.Tag("em")

// New returns an html <em> tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}
