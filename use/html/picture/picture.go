/*
	Package picture provides the HTML <picture> element.

	The <picture> HTML element contains zero or more <source>
	elements and one <img> element to offer alternative
	versions of an image for different display/device scenarios.

	The browser will consider each child <source> element and
	choose the best match among them. If no matches are found—or
	the browser doesn't support the <picture> element—the URL of
	the <img> element's src attribute is selected. The selected
	image is then presented in the space occupied by the <img>
	element.

	To decide which URL to load, the user agent examines each
	<source>'s srcset, media, and type attributes to select a
	compatible image that best matches the current layout and
	capabilities of the display device.

	The <img> element serves two purposes:

		1. It describes the size and other attributes of the
		   image and its presentation.
		2. It provides a fallback in case none of the offered
		   <source> elements are able to provide a usable image.

	Common use cases for <picture>:

		* Art direction. Cropping or modifying images for different
		  media conditions (for example, loading a simpler version
		  of an image which has too many details, on smaller displays).
		* Offering alternative image formats, for cases where certain
		  formats are not supported.

		Note: For example, newer formats like AVIF or WEBP have many a
		dvantages, but might not be supported by the browser. A list of
		supported image formats can be found in: Image file type and
		format guide.

		* Saving bandwidth and speeding page load times by loading the
		  most appropriate image for the viewer's display.

	If providing higher-density versions of an image for high-DPI
	(Retina) display, use srcset on the <img> element instead. This
	lets browsers opt for lower-density versions in data-saving modes,
	and you don't have to write explicit media conditions.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/picture
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package picture

import (
	"qlova.tech/use/html"
	"qlova.tech/web/tree"
)

// Tag is the html <picture> tag.
const Tag = html.Tag("picture")

// New returns a <picture> html tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}
