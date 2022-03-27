/*
	Package bringattention provides the HTML <b> element.

	The <b> HTML element is used to draw the reader's attention to
	the element's contents, which are not otherwise granted special
	importance. This was formerly known as the Boldface element,
	and most browsers still draw the text in boldface. However,
	you should not use <b> for styling text; instead, you should
	use the CSS font-weight property to create boldface text,
	or the <strong> element to indicate that text is of special
	importance.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/b
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package bringattention

import (
	"qlova.tech/use/html"
	"qlova.tech/web/tree"
)

// Tag is the html <b> tag.
const Tag = html.Tag("b")

// New returns a <b> html tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}
