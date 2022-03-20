/*
	Package idiomatic provides the HTML <i> element.

	The <i> HTML element represents a range of text that is set
	off from the normal text for some reason, such as idiomatic
	text, technical terms, taxonomical designations, among others.
	Historically, these have been presented using italicized type,
	which is the original source of the <i> naming of this element.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/i
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package idiomatic

import (
	"qlova.tech/new/tree"
	"qlova.tech/use/html"
)

// Tag is the html <i> tag.
const Tag = html.Tag("i")

// New returns a <i> html tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}
