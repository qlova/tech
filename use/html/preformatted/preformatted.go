/*
	Package preformatted provides the HTML <pre> element.

	The <pre> HTML element represents preformatted text which
	is to be presented exactly as written in the HTML file.
	The text is typically rendered using a non-proportional,
	or "monospaced, font. Whitespace inside this element is
	displayed as written.

	If you have reserved characters such as <, >, &, " in
	literal strings in the <pre> tag, these characters will be
	escaped using their respective HTML entity.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/pre
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package preformatted

import (
	"qlova.tech/new/tree"
	"qlova.tech/use/html"
)

// Tag is the html <pre> tag.
const Tag = html.Tag("pre")

// New returns a <pre> html tree node.
func New(args ...any) tree.Node {
	for i, arg := range args {
		if s, ok := arg.(string); ok {
			args[i] = html.Escape(s)
		}
	}

	return html.New(append(args, Tag)...)
}
