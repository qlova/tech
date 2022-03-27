/*
	Package linebreak provides the HTML <br> element.

	The <br> HTML element produces a line break in text
	(carriage-return). It is useful for writing a poem
	or an address, where the division of lines is significant.

	As you can see from the above example, a <br> element is
	included at each point where we want the text to break.
	The text after the <br> begins again at the start of the
	next line of the text block.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/br
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package linebreak

import (
	"qlova.tech/use/html"
	"qlova.tech/web/tree"
)

// Tag is the html <br> tag.
const Tag = html.Tag("br")

// Line break is an html <br> element.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag, html.Void)...)
}
