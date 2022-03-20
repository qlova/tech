/*
	Package descriptionlist provides the HTML <dl> element.

	The <dl> HTML element represents a description list. The element
	encloses a list of groups of terms (specified using the <dt>
	element) and descriptions (provided by <dd> elements). Common
	uses for this element are to implement a glossary or to display
	metadata (a list of key-value pairs).

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dl
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package descriptionlist

import (
	"qlova.tech/new/tree"
	"qlova.tech/use/html"
)

// Tag is the html <dl> tag.
const Tag = html.Tag("dl")

// New returns a <dl> html tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}
