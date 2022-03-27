/*
	Package footer provides the HTML <footer> element.

	The <footer> HTML element represents a footer for its nearest sectioning
	content or sectioning root element. A <footer> typically contains
	information about the author of the section, copyright data or links
	to related documents.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/head
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package footer

import (
	"qlova.tech/use/html"
	"qlova.tech/web/tree"
)

// Tag is the html <footer> tag.
const Tag = html.Tag("footer")

// New returns a <footer> html tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}
