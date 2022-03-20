/*
	Package noscript provides the HTML <noscript> element.

	The <noscript> HTML element defines a section of HTML to be
	inserted if a script type on the page is unsupported or if
	scripting is currently turned off in the browser.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/noscript
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package noscript

import (
	"qlova.tech/new/tree"
	"qlova.tech/use/html"
)

// Tag is the html <noscript> tag.
const Tag = html.Tag("noscript")

// New returns a <noscript> html tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}
