/*
	Package citation provides the HTML <cite> element.

	The <cite> HTML element is used to describe a reference to
	a cited creative work, and must include the title of that
	work. The reference may be in an abbreviated form according
	to context-appropriate conventions related to citation metadata.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/cite
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package citation

import (
	"qlova.tech/new/tree"
	"qlova.tech/use/html"
)

// Tag is the html <cite> tag.
const Tag = html.Tag("cite")

// New returns a <cite> html tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}
