/*
	Package details provides the HTML <details> element.

	The <details> HTML element creates a disclosure widget in which
	information is visible only when the widget is toggled into an
	"open" state. A summary or label must be provided using the
	<summary> element.

	A disclosure widget is typically presented onscreen using a small
	triangle which rotates (or twists) to indicate open/closed status,
	with a label next to the triangle. The contents of the <summary>
	element are used as the label for the disclosure widget.

	A <details> widget can be in one of two states. The default closed
	state displays only the triangle and the label inside <summary>
	(or a user agent-defined default string if no <summary>).

	When the user clicks on the widget or focuses it then presses the
	space bar, it "twists" open, revealing its contents. The common
	use of a triangle which rotates or twists around to represent
	opening or closing the widget is why these are sometimes called
	"twisties".

	You can use CSS to style the disclosure widget, and you can
	programmatically open and close the widget by setting/removing
	its open attribute. Unfortunately, at this time there's no
	built-in way to animate the transition between open and closed.

	By default when closed, the widget is only tall enough to
	display the disclosure triangle and summary. When open, it
	expands to display the details contained within.

	Fully standards-compliant implementations automatically apply
	the CSS display: list-item to the <summary> element. You can
	use this to customize its appearance further.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/details
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package details

import (
	"qlova.tech/use/html"
	"qlova.tech/web/tree"
)

// Tag is the html <details> tag.
const Tag = html.Tag("details")

// New returns an html <details> tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}

// Open indicates whether or not the details — that is, the contents
// of the <details> element — are currently visible. The details are
// shown when this attribute exists, or hidden when this attribute is
// absent. By default this attribute is absent which means the details
// are not visible.
const Open = html.Attribute("open")
