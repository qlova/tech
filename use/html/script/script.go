/*
	Package script provides the HTML <script> element.

	The <script> HTML element is used to embed executable code or
	data; this is typically used to embed or refer to JavaScript
	code. The <script> element can also be used with other languages,
	such as WebGL's GLSL shader programming language and JSON.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/script
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package script

import (
	"qlova.tech/use/html"
	"qlova.tech/web/tree"
)

// Tag is the html <script> tag.
const Tag = html.Tag("script")

// New returns a <script> html tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}

const (
	/*
		For classic scripts, if the async attribute is present, then the
		classic script will be fetched in parallel to parsing and
		evaluated as soon as it is available.

		For module scripts, if the async attribute is present then the
		scripts and all their dependencies will be executed in the defer
		queue, therefore they will get fetched in parallel to parsing
		and evaluated as soon as they are available.

		This attribute allows the elimination of parser-blocking
		JavaScript where the browser would have to load and evaluate
		scripts before continuing to parse. defer has a similar effect
		in this case.

		This is a boolean attribute: the presence of a boolean attribute
		on an element represents the true value, and the absence of the
		attribute represents the false value.

		See Browser compatibility for notes on browser support. See also
		Async scripts for asm.js.
	*/
	Async = html.Attribute("async")

	/*
		This Boolean attribute is set to indicate to a browser that the
		script is meant to be executed after the document has been parsed,
		but before firing DOMContentLoaded.

		Scripts with the defer attribute will prevent the DOMContentLoaded
		event from firing until the script has loaded and finished evaluating.

		Warning: This attribute must not be used if the src attribute is
		absent (i.e. for inline scripts), in this case it would have no effect.

		The defer attribute has no effect on module scripts — they defer by
		default.

		Scripts with the defer attribute will execute in the order in which
		they appear in the document.

		This attribute allows the elimination of parser-blocking JavaScript
		where the browser would have to load and evaluate scripts before
		continuing to parse. async has a similar effect in this case.
	*/
	Defer = html.Attribute("defer")

	/*
		This Boolean attribute is set to indicate that the script should not
		be executed in browsers that support ES2015 modules — in effect, this
		can be used to serve fallback scripts to older browsers that do not
		support modular JavaScript code.
	*/
	NoModule = html.Attribute("nomodule")
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

// Integrity contains inline metadata that a user agent can use to verify
// that a fetched resource has been delivered free of unexpected
// manipulation. See Subresource Integrity.
type Integrity string

// RenderAttr implements the attributes.Renderer interface.
func (i Integrity) RenderAttr() []byte {
	return []byte(html.Attr("integrity", string(i)))
}

// A cryptographic nonce (number used once) to allow scripts in a script-src
// Content-Security-Policy. The server must generate a unique nonce value each
// time it transmits a policy. It is critical to provide a nonce that cannot be
// guessed as bypassing a resource's policy is otherwise trivial.
type Nonce string

// RenderAttr implements the attributes.Renderer interface.
func (n Nonce) RenderAttr() []byte {
	return []byte(html.Attr("nonce", string(n)))
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

// Source specifies the URI of an external script; this can be used as an
// alternative to embedding a script directly within a document.
type Source string

// RenderAttr implements the attr.Renderer interface.
func (s Source) RenderAttr() []byte {
	return []byte(html.Attr("src", string(s)))
}

// MimeType  indicates the type of script represented.
func MimeType(mimetype string) html.Attribute {
	return html.Attr("type", mimetype)
}
