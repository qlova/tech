/*
	Package horizontalrule provides the HTML <hr> element.

	The <hr> HTML element represents a thematic break between
	paragraph-level elements: for example, a change of scene
	in a story, or a shift of topic within a section.

	Historically, this has been presented as a horizontal rule
	or line. While it may still be displayed as a horizontal
	rule in visual browsers, this element is now defined in
	semantic terms, rather than presentational terms, so if
	you wish to draw a horizontal line, you should do so
	using appropriate CSS.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/hr
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package horizontalrule

import (
	"qlova.tech/new/tree"
	"qlova.tech/use/html"
)

// Tag is the html <hr> tag.
const Tag = html.Tag("hr")

// New returns a <hr> html tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag, html.Void)...)
}
