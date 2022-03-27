/*
	Package span provides the HTML <span> element.

	The <span> HTML element is a generic inline container for
	phrasing content, which does not inherently represent anything.
	It can be used to group elements for styling purposes (using
	the class or id attributes), or because they share attribute
	values, such as lang. It should be used only when no other
	semantic element is appropriate. <span> is very much like a
	<div> element, but <div> is a block-level element whereas a
	<span> is an inline element.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/span
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package span

import (
	"qlova.tech/use/html"
	"qlova.tech/web/tree"
)

// Tag is the html <span> tag.
const Tag = html.Tag("span")

// New returns a <span> html tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}
