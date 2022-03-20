/*
	Package link provides the HTML <link> element.

	The <link> HTML element specifies relationships between the current
	document and an external resource. This element is most commonly
	used to link to stylesheets, but is also used to establish site
	icons (both "favicon" style icons and icons for the home screen
	and apps on mobile devices) among other things.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/head
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package link

import (
	"bytes"
	"strings"

	"qlova.tech/use/html"
)

// Tag is the html <link> tag.
const Tag = html.Tag("link")

// Element is a <link> element.
type Element []html.Attribute

// RenderHTML implements html.Renderer
func (elem Element) RenderHTML() []byte {
	var b bytes.Buffer
	b.WriteString("<link ")
	for _, attr := range elem {
		b.WriteString(string(attr))
		b.WriteString(" ")
	}
	b.WriteString(">")
	return b.Bytes()
}

/*
	These attributes are only used when RelPreload or RelPrefetch
	have been set on the <link> element. It specifies the type of content
	being loaded by the <link>, which is necessary for request matching,
	application of correct content security policy, and setting of
	correct Accept request header. Furthermore, RelPreload uses this
	as a signal for request prioritization.
*/
const (
	AsAudio    = html.Attribute("as=audio")
	AsDocument = html.Attribute("as=document")
	AsEmbed    = html.Attribute("as=embed")
	AsFetch    = html.Attribute("as=fetch")
	AsFont     = html.Attribute("as=font")
	AsImage    = html.Attribute("as=image")
	AsObject   = html.Attribute("as=object")
	AsScript   = html.Attribute("as=script")
	AsStyle    = html.Attribute("as=style")
	AsTrack    = html.Attribute("as=track")
	AsVideo    = html.Attribute("as=video")
	AsWorker   = html.Attribute("as=worker")
)

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

const (
	/*
		For RelStylesheet only, the disabled Boolean attribute indicates
		whether or not the described stylesheet should be loaded and applied
		to the document. If disabled is specified in the HTML when it is
		loaded, the stylesheet will not be loaded during page load. Instead,
		the stylesheet will be loaded on-demand, if and when the disabled
		attribute is changed to false or removed.

		Setting the disabled property in the DOM causes the stylesheet to be
		removed from the document's Document.styleSheets list.
	*/
	Disabled html.Attribute = " disabled"
)

// New returns a <link> html tree node.
func New(args ...html.Attribute) Element {
	return Element(args)
}

// Href specifies the URL of the linked resource. A URL can be absolute or
// relative.
func Href(href string) html.Attribute {
	return html.Attr("href", href)
}

/*
	HrefLang indicates the language of the linked resource. It is
	purely advisory. Allowed values are specified by RFC 5646: Tags for
	Identifying Languages (also known as BCP 47). Use this attribute
	only if the href attribute is present.
*/
func HrefLang(lang string) html.Attribute {
	return html.Attr("hreflang", lang)
}

/*
	ImageSizes (For RelPreload and AsImage only) is a sizes attribute
	that indicates to preload the appropriate resource used by an img
	element with corresponding values for its srcset and sizes attributes.
*/
func ImageSizes(sizes ...string) html.Attribute {
	return html.Attr("imagesizes", strings.Join(sizes, ","))
}

/*
	ImageSrcSet (For RelPreload and AsImage only) is a sourceset attribute
	that indicates to preload the appropriate resource used by an img
	element with corresponding values for its srcset and sizes attributes.
*/
func ImageSrcSet(srcset ...string) html.Attribute {
	return html.Attr("imagesrcset", strings.Join(srcset, ","))
}

/*
	Media specifies the media that the linked resource applies to.
	Its value must be a media type / media query. This attribute is mainly
	useful when linking to external stylesheets — it allows the user agent
	to pick the best adapted one for the device it runs on.
*/
func Media(media string) html.Attribute {
	return html.Attr("media", media)
}

// Type of link.
type Type string

// Link types.
const (
	TypeAlternate     = Type("alternate")
	TypeAuthor        = Type("author")
	TypeBookmark      = Type("bookmark")
	TypeCanonical     = Type("canonical")
	TypeExternal      = Type("external")
	TypeHelp          = Type("help")
	TypeIcon          = Type("icon")
	TypeLicense       = Type("license")
	TypeManifest      = Type("manifest")
	TypeMe            = Type("me")
	TypeModulePreload = Type("modulepreload")
	TypeNext          = Type("next")
	TypeNoFollow      = Type("nofollow")
	TypeNoOpener      = Type("noopener")
	TypeNoReferrer    = Type("noreferrer")
	TypePingback      = Type("pingback")
	TypePreconnect    = Type("preconnect")
	TypePrefetch      = Type("prefetch")
	TypePreload       = Type("preload")
	TypePrev          = Type("prev")
	TypeSearch        = Type("search")
	TypeShortlink     = Type("shortlink")
	TypeStylesheet    = Type("stylesheet")
	TypeTag           = Type("tag")
)

// Target specifies where to display the linked URL.
type Target string

const (
	// TargetSelf is a special Target that shows the result in the current browsing
	// context.
	TargetSelf Target = `_self`

	// TargetBlank is a special Target that shows the result in a new, unnamed
	// browsing context.
	TargetBlank Target = `_blank`

	// TargetParent is a special Target that shows the result in the parent browsing
	// context of the current one, if the current page is inside a frame. If there
	// is no parent, acts the same as TargetSelf.
	TargetParent Target = `_parent`

	// TargetTop is a special Target that shows the result in the parent browsing
	// context of the current one, if the current page is inside a frame. If there
	// is no parent, acts the same as TargetSelf.
	TargetTop Target = `_top`
)

// Rel is the relationship of the linked URL as space-separated link types.
func Rel(types ...Type) html.Attribute {
	var s []string
	for _, t := range types {
		s = append(s, string(t))
	}
	return html.Attr("rel", strings.Join(s, " "))
}

const (
	// (if RelIcon) means that the icon can be scaled to any size
	// as it is in a vector format, like image/svg+xml
	SizeAny = html.Attribute("sizes=any")
)

// Sizes defines the sizes of the icons for visual media
// contained in the resource. It must be present only if RelIcon.
func Sizes(sizes ...string) html.Attribute {
	return html.Attr("sizes", strings.Join(sizes, ","))
}

// The title attribute has special semantics on the <link> element.
// When used on a RelStylesheet it defines a default or an alternate
// stylesheet.
func Title(title string) html.Attribute {
	return html.Attr("title", title)
}

/*
	MimeType is used to define the type of the content linked to.
	The value of the attribute should be a MIME type such as text/html,
	text/css, and so on. The common use of this attribute is to define
	the type of stylesheet being referenced (such as text/css), but
	given that CSS is the only stylesheet language used on the web,
	not only is it possible to omit the type attribute, but is
	actually now recommended practice. It is also used on RelPreload
	link types, to make sure the browser only downloads file types
	that it supports.
*/
func MimeType(mimetype string) html.Attribute {
	return html.Attr("type", mimetype)
}
