/*
	Package textarea provides the HTML <textarea> element.

	The <textarea> HTML element represents a multi-line plain-text
	editing control, useful when you want to allow users to enter
	a sizeable amount of free-form text, for example a comment on
	a review or feedback form.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/textarea
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package textarea

import (
	"strconv"

	"qlova.tech/new/tree"
	"qlova.tech/use/html"
)

// Tag is the html <textarea> tag.
const Tag = html.Tag("textarea")

// New returns an html <textarea> tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}

// Columns sets the visible width of the text control, in average
// character widths. If it is specified, it must be a positive
// integer. If it is not specified, the default value is 20.
type Columns uint

// RenderAttr implements attributes.Renderer
func (c Columns) RenderAttr() []byte {
	return []byte(html.Attr("cols", strconv.FormatUint(uint64(c), 10)))
}

// Rows are the number of visible text lines for the control. If
// it is specified, it must be a positive integer. If it is not
// specified, the default value is 2.
type Rows uint

// RenderAttr implements attributes.Renderer
func (r Rows) RenderAttr() []byte {
	return []byte(html.Attr("rows", strconv.FormatUint(uint64(r), 10)))
}

// Specifies whether the <textarea> is subject to spell checking by the underlying browser/OS.
const (
	SpellcheckOn      = html.Attribute("spellcheck=on")
	SpellcheckDefault = html.Attribute("spellcheck=default")
	SpellcheckOff     = html.Attribute("spellcheck=off")
)

// Indicates how the control wraps text.
const (
	WrapHard = html.Attribute("wrap=hard")
	WrapSoft = html.Attribute("wrap=soft")
)
