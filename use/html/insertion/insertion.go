/*
	Package insertion provides the HTML <ins> element.

	The <ins> HTML element represents a range of text that has been
	added to a document. You can use the <del> element to similarly
	represent a range of text that has been deleted from the document.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ins
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package insertion

import (
	"qlova.tech/use/html"
	"qlova.tech/web/tree"
)

// Tag is the html <ins> tag.
const Tag = html.Tag("ins")

// New returns an html <ins> tree node.
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
