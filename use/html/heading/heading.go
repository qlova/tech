/*
	Package heading provides the HTML <h1>–<h6> elements.

	The <h1> to <h6> HTML elements represent six levels of
	section headings. <h1> is the highest section level
	and <h6> is the lowest.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package h1

import (
	"qlova.tech/new/tree"
	"qlova.tech/use/html"
)

const (
	// One is the html <h1> tag.
	One    = html.Tag("h1")
	Level1 = One

	// Two is the html <h2> tag.
	Two    = html.Tag("h2")
	Level2 = Two

	// Three is the html <h3> tag.
	Three  = html.Tag("h3")
	Level3 = Three

	// Four is the html <h4> tag.
	Four   = html.Tag("h4")
	Level4 = Four

	// Five is the html <h5> tag.
	Five   = html.Tag("h5")
	Level5 = Five

	// Six is the html <h6> tag.
	Six    = html.Tag("h6")
	Level6 = Six
)

// New returns an html <h1> tree node.
// Other heading levels can be specified
// by including a Level argument.
func New(args ...any) tree.Node {
	return html.New(append(args, One)...)
}
