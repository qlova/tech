/*
	Package definition provides the HTML <dfn> element.

	The <dfn> HTML element is used to indicate the term being
	defined within the context of a definition phrase or
	sentence. The <p> element, the <dt>/<dd> pairing, or the
	<section> element which is the nearest ancestor of the
	<dfn> is considered to be the definition of the term.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dfn
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package definition

import (
	"qlova.tech/new/tree"
	"qlova.tech/use/html"
)

// Tag is the html <dfn> tag.
const Tag = html.Tag("dfn")

// New returns a <dfn> html tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}
