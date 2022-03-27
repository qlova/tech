/*
	Package keyboard provides the HTML <kbd> element.

	The <kbd> HTML element represents a span of inline text
	denoting textual user input from a keyboard, voice input,
	or any other text entry device. By convention, the user
	agent defaults to rendering the contents of a <kbd>
	element using its default monospace font, although this
	is not mandated by the HTML standard.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/kbd
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package keyboard

import (
	"qlova.tech/use/html"
	"qlova.tech/web/tree"
)

// Tag is the html <kbd> tag.
const Tag = html.Tag("kbd")

// New returns an html <kbd> tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, html.Tag("kbd"))...)
}
