/*
	Package image provides the HTML <img> element.

	The <img> HTML element embeds an image into the document.

	The above example shows usage of the <img> element:

		- The src attribute is required, and contains the
		  path to the image you want to embed.
		- The alt attribute holds a text description of the
		  image, which isn't mandatory but is incredibly useful
		  for accessibility — screen readers read this
		  description out to their users so they know what the
		  image means. Alt text is also displayed on the page
		  if the image can't be loaded for some reason: for
		  example, network errors, content blocking, or linkrot.

	There are many other attributes to achieve various purposes:

		- Referrer/CORS control for security and privacy: see
		  crossorigin and referrerpolicy.
		- Use both width and height to set the intrinsic size
		  of the image, allowing it to take up space before it
		  loads, to mitigate content layout shifts.
		- Responsive image hints with sizes and srcset (see also
		  the <picture> element and our Responsive images tutorial).

	Supported Image Formats

	The HTML standard doesn't list what image formats to support,
	so user agents may support different formats.

	The image file formats that are most commonly used on the web are:

		- APNG (Animated Portable Network Graphics) — Good choice for
		  lossless animation sequences (GIF is less performant)
		- AVIF (AV1 Image File Format) — Good choice for both images
		  and animated images due to high performance.
		- GIF (Graphics Interchange Format) — Good choice for simple
		  images and animations.
		- JPEG (Joint Photographic Expert Group image) — Good choice
		  for lossy compression of still images (currently the most popular).
		- PNG (Portable Network Graphics) — Good choice for lossy
		  compression of still images (slightly better quality than JPEG).
		- SVG (Scalable Vector Graphics) — Vector image format. Use for
		  images that must be drawn accurately at different sizes.
		- WebP (Web Picture format) — Excellent choice for both images
		  and animated images

	Formats like WebP and AVIF are recommended as they perform much better than PNG, JPEG, GIF for both still and animated images. WebP is widely supported while AVIF lacks support in Safari.

	SVG remains the recommended format for images that must be drawn accurately at different sizes.

	Image Loading Errors

	If an error occurs while loading or rendering an image, and an
	onerror event handler has been set on the error event, that
	event handler will get called. This can happen in a number
	of situations, including:

		- The src attribute is empty ("") or null.
		- The src URL is the same as the URL of the page the user
		  is currently on.
		- The image is corrupted in some way that prevents it
		  from being loaded.
		- The image's metadata is corrupted in such a way that
		  it's impossible to retrieve its dimensions, and no
		  dimensions were specified in the <img> element's attributes.
		- The image is in a format not supported by the user agent.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/img
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package image

import (
	"fmt"
	"strings"

	"qlova.tech/use/css"
	"qlova.tech/use/html"
	"qlova.tech/web/tree"
)

// Tag is the html <img> tag.
const Tag = html.Tag("img")

// New returns an HTML <img> tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag, html.Void)...)
}

// Alt is used to define an alternative text
// description of the image. Omitting alt altogether indicates
// that the image is a key part of the content and no textual
// equivalent is available. Setting this attribute to an empty
// string (alt="") indicates that this image is not a key part
// of the content (it's decoration or a tracking pixel), and
// that non-visual browsers may omit it from rendering. Visual
// browsers will also hide the broken image icon if the alt is
// empty and the image failed to display. This attribute is also
// used when copying and pasting the image to text, or saving a
// linked image to a bookmark.
type Alt string

// RenderAttr implements the attr.Renderer interface.
func (a Alt) RenderAttr() []byte {
	return []byte(html.Attr("alt", string(a)))
}

/*
	These enumerated attributes indicates whether CORS must be used when
	fetching the resource. CORS-enabled images can be reused in the
	<canvas> element without being tainted. The allowed values are:
*/
const (
	/*
		A cross-origin request (i.e. with an Origin HTTP header) is performed,
		but no credential is sent (i.e. no cookie, X.509 certificate, or HTTP
		Basic authentication). If the server does not give credentials to the
		origin site (by not setting the Access-Control-Allow-Origin HTTP
		header) the resource will be tainted and its usage restricted.
	*/
	CrossOriginAnonymous = html.Attribute("crossorigin=anonymous")

	/*
		A cross-origin request (i.e. with an Origin HTTP header) is performed
		along with a credential sent (i.e. a cookie, certificate, and/or HTTP
		Basic authentication is performed). If the server does not give
		credentials to the origin site (through Access-Control-Allow-Credentials
		HTTP header), the resource will be tainted and its usage restricted.
	*/
	CrossOriginUseCredentials = html.Attribute("crossorigin=use-credentials")
)

// image decoding hints for the browser
const (
	// Decode the image synchronously, for atomic
	// presentation with other content.
	DecodingSync = html.Attribute("decoding=sync")

	// Decode the image asynchronously, to reduce
	// delay in presenting other content.
	DecodingAsync = html.Attribute("decoding=async")

	// Default: no preference for the decoding mode.
	// The browser decides what is best for the user.
	DecodingAuto = html.Attribute("decoding=auto")
)

// The height of the image in pixels.
type Height uint

// RenderAttr implements the attr.Renderer interface.
func (h Height) RenderAttr() []byte {
	return []byte(html.Attr("height", fmt.Sprint(h)))
}

// IsMap indicates that the image is part of a server-side map.
// If so, the coordinates where the user clicked on the
// image are sent to the server.
const IsMap = html.Attribute("ismap")

// Indicates how the browser should load the image:
const (
	// Loads the image immediately, regardless of whether or
	// not the image is currently within the visible viewport
	// (this is the default value).
	LoadingEager = html.Attribute("loading=eager")

	// Defers loading the image until it reaches a calculated
	// distance from the viewport, as defined by the browser.
	// The intent is to avoid the network and storage bandwidth
	// needed to handle the image until it's reasonably certain
	// that it will be needed. This generally improves the
	// performance of the content in most typical use cases.
	LoadingLazy = html.Attribute("loading=lazy")
)

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

// Size used for sizes.
type Size struct {
	MediaCondition css.MediaQuery
	Value          css.Size
}

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

// Source URL. Mandatory for the <img> element. On browsers supporting
// srcset, src is treated like a candidate image with a pixel density
// descriptor 1x, unless an image with this pixel density descriptor
// is already defined in srcset, or unless srcset contains w descriptors.
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

// Width is the intrinsic width of the image in pixels.
type Width uint

// RenderAttr implements the attr.Renderer interface.
func (w Width) RenderAttr() []byte {
	return []byte(html.Attr("width", fmt.Sprint(w)))
}

// Map is the partial URL (starting with #) of an image map associated
// with the element.
type Map string

// RenderAttr implements the attr.Renderer interface.
func (m Map) RenderAttr() []byte {
	return []byte(html.Attr("usemap", string(m)))
}
