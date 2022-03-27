/*
	Package head provides the HTML <head> element.

	The <head> HTML element contains machine-readable information
	(metadata) about the document, like its title, scripts, and style
	sheets.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/head
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package head

import (
	"qlova.tech/use/html"
	"qlova.tech/web/tree"
)

// Tag is the html <head> tag.
const Tag = html.Tag("head")

// New returns a <head> html tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}
