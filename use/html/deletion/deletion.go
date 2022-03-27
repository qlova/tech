/*
	Package deletion provides the HTML <del> element.

	The <del> HTML element represents a range of text that has
	been deleted from a document. This can be used when
	rendering "track changes" or source code diff information,
	for example. The <ins> element can be used for the opposite
	purpose: to indicate text that has been added to the document.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/del
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package deletion

import (
	"qlova.tech/use/html"
	"qlova.tech/web/tree"
)

// Tag is the html <del> tag.
const Tag = html.Tag("del")

// New returns an html <del> tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}

// Cite a URI for a resource that explains the change
// (for example, meeting minutes).
func Cite(url string) html.Attribute {
	return html.Attr("cite", url)
}

// This attribute indicates the time and date of the change
// and must be a valid date string with an optional time.
// If the value cannot be parsed as a date with an optional time string,
// the element does not have an associated time stamp.
func DateTime(value string) html.Attribute {
	return html.Attr("datetime", value)
}
