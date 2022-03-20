/*
	Package quote provides the HTML <q> element.

	The <q> HTML element indicates that the enclosed text is a
	short inline quotation. Most modern browsers implement this
	by surrounding the text in quotation marks. This element is
	intended for short quotations that don't require paragraph
	breaks; for long quotations use the <blockquote> element.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/q
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package quote

import (
	"qlova.tech/new/tree"
	"qlova.tech/use/html"
)

// Tag is the html <q> tag.
const Tag = html.Tag("q")

// New returns an html <q> tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}

// Cite a URL that designates a source document or message for the information
// quoted. This attribute is intended to point to information explaining the
// context or the reference for the quote.
func Cite(cite string) html.Attribute {
	return html.Attr("cite", cite)
}
