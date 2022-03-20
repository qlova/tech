/*
	Package descriptionterm provides the HTML <dt> element.

	The <dt> HTML element specifies a term in a description or definition
	list, and as such must be used inside a <dl> element. It is usually
	followed by a <dd> element; however, multiple <dt> elements in a row
	indicate several terms that are all defined by the immediate next
	<dd> element.

	The subsequent <dd> (Description Details) element provides the
	definition or other related text associated with the term specified
	using <dt>.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dt
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package descriptionterm

import (
	"qlova.tech/new/tree"
	"qlova.tech/use/html"
)

// Tag is the html <dt> tag.
const Tag = html.Tag("dt")

// New returns a <dt> html tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}
