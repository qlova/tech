/*
	Package meter provides the HTML <meter> element.

	The <meter> HTML element represents either a scalar
	value within a known range or a fractional value.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/meter
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package meter

import (
	"fmt"

	"qlova.tech/new/tree"
	"qlova.tech/use/html"
)

// Tag is the html <meter> tag.
const Tag = html.Tag("meter")

// New returns an html <meter> tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}

// Value sets the current numeric value. This must be between the
// minimum and maximum values (min attribute and max attribute) if
// they are specified. If unspecified or malformed, the value is 0.
// If specified, but not within the range given by the min attribute
// and max attribute, the value is equal to the nearest end of the range.
type Value float64

// RenderAttr implements the attributes.Renderer interface.
func (v Value) RenderAttr() []byte {
	return []byte(html.Attr("value", fmt.Sprint(v)))
}

// Min is the lower numeric bound of the measured range. This must be
// less than the maximum value (max attribute), if specified. If
// unspecified, the minimum value is 0.
type Min float64

// RenderAttr implements the attributes.Renderer interface.
func (v Min) RenderAttr() []byte {
	return []byte(html.Attr("min", fmt.Sprint(v)))
}

// Max is the upper numeric bound of the measured range. This must be
// greater than the minimum value (min attribute), if specified. If
// unspecified, the maximum value is 1.
type Max float64

// RenderAttr implements the attributes.Renderer interface.
func (v Max) RenderAttr() []byte {
	return []byte(html.Attr("max", fmt.Sprint(v)))
}

// Low is the upper numeric bound of the low end of the measured range.
// This must be greater than the minimum value (min attribute), and it
// also must be less than the high value and maximum value (high
// attribute and max attribute, respectively), if any are specified.
// If unspecified, or if less than the minimum value, the low value is
// equal to the minimum value.
type Low float64

// RenderAttr implements the attributes.Renderer interface.
func (v Low) RenderAttr() []byte {
	return []byte(html.Attr("low", fmt.Sprint(v)))
}

// High is the lower numeric bound of the high end of the measured range.
// This must be less than the maximum value (max attribute), and it also
// must be greater than the low value and minimum value (low attribute
// and min attribute, respectively), if any are specified. If unspecified,
// or if greater than the maximum value, the high value is equal to the
// maximum value.
type High float64

// RenderAttr implements the attributes.Renderer interface.
func (v High) RenderAttr() []byte {
	return []byte(html.Attr("high", fmt.Sprint(v)))
}

// Optimum indicates the optimal numeric value. It must be within the range
// (as defined by the min attribute and max attribute). When used with the
// low attribute and high attribute, it gives an indication where along the
// range is considered preferable. For example, if it is between the min
// attribute and the low attribute, then the lower range is considered
// preferred. The browser may color the meter's bar differently depending
// on whether the value is less than or equal to the optimum value.
type Optimum float64

// The <form> element to associate the <meter> element with (its form owner).
// The value of this attribute must be the id of a <form> in the same document.
// If this attribute is not set, the <meter> is associated with its ancestor
// <form> element, if any. This attribute is only used if the <meter> element
// is being used as a form-associated element, such as one displaying a
// range corresponding to an <input type="number">.
func Form(id html.ID) html.Attribute {
	return html.Attr("form", string(id))
}
