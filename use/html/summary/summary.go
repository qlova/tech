/*
	Package summary provides the HTML <summary> element.

	The <summary> HTML element specifies a summary, caption,
	or legend for a <details> element's disclosure box.
	Clicking the <summary> element toggles the state of
	the parent <details> element open and closed.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/summary
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package summary

import (
	"qlova.tech/use/html"
	"qlova.tech/web/tree"
)

// Tag is the html <summary> tag.
const Tag = html.Tag("summary")

// New returns an html <summary> tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}
