/*
	Package sample provides the HTML <samp> element.

	The <samp> HTML element is used to enclose inline text
	which represents sample (or quoted) output from a
	computer program. Its contents are typically rendered
	using the browser's default monospaced font (such as
	Courier or Lucida Console).

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/samp
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package sample

import (
	"qlova.tech/new/tree"
	"qlova.tech/use/html"
)

// Tag is the html <samp> tag.
const Tag = html.Tag("samp")

// New returns an html <samp> tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}
