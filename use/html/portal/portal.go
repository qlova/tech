/*
	Package portal provides the HTML <portal> element.

	The <portal> HTML element enables the embedding of another HTML
	page into the current one for the purposes of allowing smoother
	navigation into new pages.

	A <portal> is similar to an <iframe>. An <iframe> allows a
	separate browsing context to be embedded. However, the embedded
	content of a <portal> is more limited than that of an <iframe>.
	It cannot be interacted with, and therefore is not suitable for
	embedding widgets into a document. Instead, the <portal> acts
	as a preview of the content of another page. It can be navigated
	into therefore allowing for seamless transition to the embedded
	content.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/portal
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package portal

import (
	"qlova.tech/use/html"
	"qlova.tech/web/tree"
)

// Tag is the html <portal> tag.
const Tag = html.Tag("portal")

// New returns an html <portal> tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
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

// Source URL of the page to embed.
type Source string

// RenderAttr implements the attr.Renderer interface.
func (s Source) RenderAttr() []byte {
	return []byte(html.Attr("src", string(s)))
}
