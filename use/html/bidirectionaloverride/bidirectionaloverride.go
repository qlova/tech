/*
	Package bidirectionaloverride provides the HTML <bdo> element.

	The <bdo> HTML element overrides the current directionality
	of text, so that the text within is rendered in a different
	direction.

	The text's characters are drawn from the starting point in
	the given direction; the individual characters' orientation
	is not affected (so characters don't get drawn backward,
	for example).

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/bdo
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package bidirectionaloverride

import (
	"qlova.tech/use/html"
	"qlova.tech/web/tree"
)

// Tag is the html <bdo> tag.
const Tag = html.Tag("bdo")

// New returns a <bdo> html tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}

// The direction in which text should be rendered in this
// element's contents.
const (
	LeftToRight = html.Attribute("dir=ltr")
	RightToLeft = html.Attribute("dir=rtl")
)
