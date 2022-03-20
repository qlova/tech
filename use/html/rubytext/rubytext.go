/*
	Package rubytext provides the HTML <rt> element.

	The <rt> HTML element specifies the ruby text component
	of a ruby annotation, which is used to provide pronunciation,
	translation, or transliteration information for East Asian
	typography. The <rt> element must always be contained within
	a <ruby> element.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/rt
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package rubytext

import (
	"qlova.tech/new/tree"
	"qlova.tech/use/html"
)

// Tag is the html <rt> tag.
const Tag = html.Tag("rt")

// New returns an html <rt> tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}
