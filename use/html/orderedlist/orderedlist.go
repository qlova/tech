/*
	Package orderedlist provides the HTML <ol> element.

	The <ol> HTML element represents an ordered list of
	items — typically rendered as a numbered list.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ol
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package orderedlist

import (
	"fmt"

	"qlova.tech/use/html"
	"qlova.tech/web/tree"
)

// Tag is the html <ol> tag.
const Tag = html.Tag("ol")

// New returns a <ol> html tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}

// Reversed specifies that the list's items are in reverse
// order. Items will be numbered from high to low.
const Reversed = html.Attribute("reversed")

// Start is integer to start counting from for the list items.
// Always an Arabic numeral (1, 2, 3, etc.), even when the
// numbering type is letters or Roman numerals. For example,
// to start numbering elements from the letter "d" or the
// Roman numeral "iv," use start="4".
func Start(from uint) html.Attribute {
	return html.Attr("start", fmt.Sprint(from))
}

// Numbering types.
const (
	Lowercase = html.Attribute("type=a")
	Uppercase = html.Attribute("type=A")

	RomanLowercase = html.Attribute("type=i")
	RomanUppercase = html.Attribute("type=I")
)
