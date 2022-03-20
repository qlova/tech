/*
	Package dialog provides the HTML <dialog> element.

	The <dialog> HTML element represents a dialog box or other
	interactive component, such as a dismissible alert, inspector,
	or subwindow.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dialog
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package dialog

import (
	"qlova.tech/new/tree"
	"qlova.tech/use/html"
)

// Tag is the html <dialog> tag.
const Tag = html.Tag("dialog")

// New returns a <dialog> html tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}

// Open indicates that the dialog is active and can be interacted
// with. When the open attribute is not set, the dialog shouldn't
// be shown to the user.
const Open = html.Attribute("open")
