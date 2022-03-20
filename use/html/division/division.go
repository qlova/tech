/*
	Package division provides the HTML <div> element.

	The <div> HTML element is the generic container for flow content.
	It has no effect on the content or layout until styled in some
	way using CSS (e.g. styling is directly applied to it, or some
	kind of layout model like Flexbox is applied to its parent
	element).

	As a "pure" container, the <div> element does not inherently
	represent anything. Instead, it's used to group content so it
	can be easily styled using the class or id attributes, marking
	a section of a document as being written in a different
	language (using the lang attribute), and so on.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/div
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package division

import (
	"qlova.tech/new/tree"
	"qlova.tech/use/html"
)

// Tag is the html <div> tag.
const Tag = html.Tag("div")

// New returns a <div> html tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}
