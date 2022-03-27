/*
	Package progress provides the HTML <progress> element.

	The <progress> HTML element displays an indicator
	showing the completion progress of a task, typically
	displayed as a progress bar.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/progress
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package progress

import (
	"fmt"

	"qlova.tech/use/html"
	"qlova.tech/web/tree"
)

// Tag is the html <progress> tag.
const Tag = html.Tag("progress")

// New returns an html <progress> tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}

// Max describes how much work the task indicated by the progress
// element requires.
type Max float64

// RenderAttr implements the attributes.Renderer interface.
func (m Max) RenderAttr() []byte {
	return []byte(html.Attr("max", fmt.Sprint(m)))
}

// Value describes how much of the task that has been completed.
// If there is no value attribute, the progress bar is indeterminate;
// this indicates that an activity is ongoing with no indication of
// how long it is expected to take.
type Value float64

// RenderAttr implements the attributes.Renderer interface.
func (v Value) RenderAttr() []byte {
	return []byte(html.Attr("value", fmt.Sprint(v)))
}
