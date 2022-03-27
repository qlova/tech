/*
	Package area provides the HTML <area> element.

	The <area> HTML element defines an area inside an image map
	that has predefined clickable areas. An image map allows
	geometric areas on an image to be associated with hypertext
	link.

	This element is used only within a <map> element.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/area
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package area

import (
	"fmt"
	"strings"

	"qlova.tech/use/html"
	"qlova.tech/use/html/link"
	"qlova.tech/web/tree"
)

// Tag is the html <area> tag.
const Tag = html.Tag("area")

// New returns an html <area> tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag, html.Void)...)
}

// Alt is a text string alternative to display on browsers that do not
// display images. The text should be phrased so that it presents the
// user with the same kind of choice as the image would offer when
// displayed without the alternative text. This attribute is required
// only if the href attribute is used.
func Alt(value string) html.Attribute {
	return html.Attr("alt", value)
}

// Vertex is a coordinate pair.
type Vertex struct {
	X, Y int
}

// Rect is a shape that specifies the coordinates of the top-left and
// bottom-right corner of the rectangle.
type Rect struct {
	TopLeft, BottomRight Vertex
}

// RenderAttr implements attr.Renderer
func (r Rect) RenderAttr() []byte {
	return []byte(fmt.Sprintf(`shape=rect coords=%d,%d,%d,%d`, r.TopLeft.X, r.TopLeft.Y, r.BottomRight.X, r.BottomRight.Y))
}

// Circle is a shape that specifies the coordinates of the center and
// radius of the circle.
type Circle struct {
	Vertex
	Radius int
}

// RenderAttr implements attr.Renderer
func (c Circle) RenderAttr() []byte {
	return []byte(fmt.Sprintf(`shape=circle coords=%d,%d,%d`, c.X, c.Y, c.Radius))
}

// Poly is a shape that specifies the coordinates of the vertices of
// the polygon. If the first and last coordinate pairs are not the same,
// the browser will add the last coordinate pair to close the polygon
type Poly struct {
	Coords []Vertex
}

// RenderAttr implements attr.Renderer
func (p Poly) RenderAttr() []byte {
	var s []string
	for _, c := range p.Coords {
		s = append(s, fmt.Sprintf("%d", c.X))
		s = append(s, fmt.Sprintf("%d", c.Y))
	}
	return []byte(fmt.Sprintf(`shape=poly coords=%s`, strings.Join(s, ",")))
}

/*
	Causes the browser to treat the linked URL as a download the browser
	will suggest a filename/extension, generated from various sources:

		- The Content-Disposition HTTP header
		- The final segment in the URL path
		- The media type (from the Content-Type header, the start of a
		  data: URL, or Blob.type for a blob: URL)
*/
const Download = html.Attribute("download")

// Filename causes the browser to treat the linked URL as a download
// the browser the given filename is suggested / and \ characters are
// converted to underscores (_). Filesystems may forbid other
// characters in filenames, so browsers will adjust the suggested
// name if necessary.
func Filename(filename string) html.Attribute {
	return html.Attr("download", filename)
}

/*
	Href sets the URL that the hyperlink points to. Links are not
	restricted to HTTP-based URLs — they can use any URL scheme
	supported by browsers:

		- Sections of a page with fragment URLs
		- Pieces of media files with media fragments
		- Telephone numbers with tel: URLs
		- Email addresses with mailto: URLs
*/
func Href(href string) html.Attribute {
	return html.Attr("href", href)
}

// Hints at the human language of the linked URL. No built-in
// functionality. Allowed values are the same as the global lang attribute.
func HrefLang(lang string) html.Attribute {
	return html.Attr("hreflang", lang)
}

// Ping when the link is followed, the browser will send POST requests
// with the body PING to the URLs. Typically for tracking.
func Ping(urls ...string) html.Attribute {
	return html.Attr("ping", strings.Join(urls, " "))
}

// How much of the referrer to send when following the link.
const (
	// The Referer header will not be sent.
	NoReferrer = html.Attribute("referrerpolicy=no-referrer")

	// The Referer header will not be sent to origins without TLS (HTTPS).
	NoReferrerWhenDowngrade = html.Attribute("referrerpolicy=no-referrer-when-downgrade")

	// The sent referrer will be limited to the origin of the
	// referring page: its scheme, host, and port.
	ReferrerOrigin = html.Attribute("referrerpolicy=origin")

	// The referrer sent to other origins will be limited to the scheme,
	// the host, and the port. Navigations on the same origin will still include the path.
	ReferrerOriginWhenCrossOrigin = html.Attribute("referrerpolicy=origin-when-cross-origin")

	// A referrer will be sent for same origin, but cross-origin
	// requests will contain no referrer information.
	ReferrerSameOrigin = html.Attribute("referrerpolicy=same-origin")

	// Only send the origin of the document as the referrer when
	// the protocol security level stays the same (HTTPS→HTTPS),
	// but don't send it to a less secure destination (HTTPS→HTTP).
	ReferrerStrictOrigin = html.Attribute("referrerpolicy=strict-origin")

	// (default) Send a full URL when performing a same-origin request,
	// only send the origin when the protocol security level stays the
	// same (HTTPS→HTTPS), and send no header to a less secure
	// destination (HTTPS→HTTP).
	ReferrerStrictOriginWhenCrossOrigin = html.Attribute("referrerpolicy=strict-origin-when-cross-origin")

	// The referrer will include the origin and the path (but not the
	// fragment, password, or username). This value is unsafe, because
	// it leaks origins and paths from TLS-protected resources to insecure origins.
	ReferrerUnsafeUrl = html.Attribute("referrerpolicy=unsafe-url")
)

// Rel is the relationship of the linked URL as space-separated link types.
func Rel(types ...link.Type) html.Attribute {
	var s []string
	for _, t := range types {
		s = append(s, string(t))
	}
	return html.Attr("rel", strings.Join(s, " "))
}

// Target defines a keyword or author-defined name of the default browsing context
// to show the results of navigation from <a>, <area>, or <form> elements without
// explicit target attributes.
func Target(target link.Target) html.Attribute {
	return html.Attr("target", string(target))
}
