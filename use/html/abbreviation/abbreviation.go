/*
	Package abbreviation provides the HTML <abbr> element.

	The <abbr> HTML element represents an abbreviation or acronym;
	the optional title attribute can provide an expansion or
	description for the abbreviation. If present, title must
	contain this full description and nothing else.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/abbr
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package abbreviation

import (
	"qlova.tech/new/tree"
	"qlova.tech/use/html"
)

// Tag is the html <abbr> tag.
const Tag = html.Tag("abbr")

// New returns a <abbr> html tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}
