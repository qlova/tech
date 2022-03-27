/*
	Package section provides the HTML <blockquote> element.

	The <blockquote> HTML element indicates that the enclosed text
	is an extended quotation. Usually, this is rendered visually by
	indentation (see Notes for how to change it). A URL for the
	source of the quotation may be given using the cite attribute,
	while a text representation of the source can be given using
	the <cite> element.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/blockquote
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package blockquote

import (
	"qlova.tech/use/html"
	"qlova.tech/web/tree"
)

// Tag is the html <blockquote> tag.
const Tag = html.Tag("blockquote")

// New returns a <blockquote> html tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}

// Cite a URL that designates a source document or message for the information
// quoted. This attribute is intended to point to information explaining the
// context or the reference for the quote.
func Cite(cite string) html.Attribute {
	return html.Attr("cite", cite)
}
