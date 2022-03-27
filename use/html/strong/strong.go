/*
	Package strong provides the HTML <strong> element.

	The <strong> HTML element indicates that its contents have
	strong importance, seriousness, or urgency. Browsers
	typically render the contents in bold type.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/strong
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package strong

import (
	"qlova.tech/use/html"
	"qlova.tech/web/tree"
)

// Tag is the html <strong> tag.
const Tag = html.Tag("strong")

// New returns an html <strong> tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}
