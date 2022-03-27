/*
	Package template provides the HTML <template> element.

	The <template> HTML element is a mechanism for holding HTML
	that is not to be rendered immediately when a page is loaded
	but may be instantiated subsequently during runtime using
	JavaScript.

	Think of a template as a content fragment that is being
	stored for subsequent use in the document. While the parser
	does process the contents of the <template> element while
	loading the page, it does so only to ensure that those
	contents are valid; the element's contents are not rendered,
	however.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/template
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package template

import (
	"qlova.tech/use/html"
	"qlova.tech/web/tree"
)

// Tag is the html <template> tag.
const Tag = html.Tag("template")

// New returns an html <template> tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}
