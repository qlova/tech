/*
	Package source provides the HTML <source> element.

	The <source> HTML element specifies multiple media resources
	for the <picture>, the <audio> element, or the <video> element.
	It is an empty element, meaning that it has no content and does
	not have a closing tag. It is commonly used to offer the same
	media content in multiple file formats in order to provide
	compatibility with a broad range of browsers given their
	differing support for image file formats and media file formats.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/portal
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package source

import (
	"fmt"
	"strings"

	"qlova.tech/new/tree"
	"qlova.tech/use/css"
	"qlova.tech/use/html"
)

// Tag is the html <source> tag.
const Tag = html.Tag("source")

// New returns an html <source> tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag, html.Void)...)
}

// MimeType to use to select the plug-in to instantiate.
func MimeType(mimetype string) html.Attribute {
	return html.Attr("type", mimetype)
}

// Source is required if the source element’s parent is an
// <audio> and <video> element, but not allowed if the
// source element’s parent is a <picture> element.
//
// Address of the media resource.
type Source string

// RenderAttr implements the attr.Renderer interface.
func (s Source) RenderAttr() []byte {
	return []byte(html.Attr("src", string(s)))
}

/*
	One or more strings separated by commas, indicating possible image
	sources for the user agent to use. Each string is composed of:

		1. A URL to an image
		2. Optionally, whitespace followed by one of:
			- A width descriptor (a positive integer directly followed by w).
			  The width descriptor is divided by the source size given in the
			  sizes attribute to calculate the effective pixel density.
		   	- A pixel density descriptor (a positive floating point number
			  directly followed by x).

	If no descriptor is specified, the source is assigned the default
	descriptor of 1x.

	It is incorrect to mix width descriptors and pixel density descriptors
	in the same srcset attribute. Duplicate descriptors (for instance, two
	sources in the same srcset which are both described with 2x) are also
	invalid.

	The user agent selects any of the available sources at its discretion.
	This provides them with significant leeway to tailor their selection
	based on things like user preferences or bandwidth conditions.
*/
type SourceSet []Source

// Size used for sizes.
type Size struct {
	MediaCondition css.MediaQuery
	Value          css.Size
}

// Allowed if the source element’s parent is a <picture> element, but not
// allowed if the source element’s parent is an <audio> or <video> element.
//
// Sizes specify the intended display size of the image. User agents use
// the current source size to select one of the sources supplied by the
// srcset attribute, when those sources are described using width (w)
// descriptors. The selected source size affects the intrinsic size of
// the image (the image's display size if no CSS styling is applied).
// If the srcset attribute is absent, or contains no values with a
// width descriptor, then the sizes attribute has no effect.
type Sizes []Size

// RenderAttr implements the attr.Renderer interface.
func (sizes Sizes) RenderAttr() []byte {
	var s []string
	for _, size := range sizes {
		s = append(s, fmt.Sprintf("%s %s", size.MediaCondition, size.Value))
	}
	return []byte(html.Attr("sizes", strings.Join(s, ",")))
}

// Allowed if the source element’s parent is a <picture> element, but
// not allowed if the source element’s parent is an <audio> or
// <video> element. Media query of the resource's intended media.
func Media(query css.MediaQuery) html.Attribute {
	return html.Attr("media", string(query))
}

// Width is the intrinsic width of the image in pixels.
type Width uint

// RenderAttr implements the attr.Renderer interface.
func (w Width) RenderAttr() []byte {
	return []byte(html.Attr("width", fmt.Sprint(w)))
}

// The height of the image in pixels.
type Height uint

// RenderAttr implements the attr.Renderer interface.
func (h Height) RenderAttr() []byte {
	return []byte(html.Attr("height", fmt.Sprint(h)))
}
