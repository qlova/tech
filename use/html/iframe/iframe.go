/*
	Package iframe provides the HTML <iframe> element.

	The <iframe> HTML element represents a nested browsing
	context, embedding another HTML page into the current one.

	Each embedded browsing context has its own session history
	and document. The browsing context that embeds the others
	is called the *parent browsing context*. The *topmost*
	browsing context — the one with no parent — is usually
	the browser window, represented by the Window object.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/iframe
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package iframe

import (
	"fmt"
	"strings"

	"qlova.tech/use/html"
	"qlova.tech/web/tree"
)

// Tag is the html <iframe> tag.
const Tag = html.Tag("iframe")

// New returns a new <iframe> element.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}

// Feature policy.
type Feature string

// Specifies a feature policy for the <iframe>. The policy defines what
// features are available to the <iframe> based on the origin of the
// request (e.g. access to the microphone, camera, battery, web-share API, etc.).
type Allow []Feature

// RenderAttr implements the attr.Renderer interface.
func (a Allow) RenderAttr() []byte {
	var s []string
	for _, f := range a {
		s = append(s, string(f))
	}
	return []byte(html.Attr("allow", strings.Join(s, ";")))
}

// ContentSecurityPolicy enforced for the embedded resource.
type ContentSecurityPolicy string

// RenderAttr implements the attr.Renderer interface.
func (c ContentSecurityPolicy) RenderAttr() []byte {
	return []byte(html.Attr("csp", string(c)))
}

// Width of the iframe.
type Width uint

// RenderAttr implements the attr.Renderer interface.
func (w Width) RenderAttr() []byte {
	return []byte(html.Attr("width", fmt.Sprint(w)))
}

// Height of the iframe.
type Height uint

// RenderAttr implements the attr.Renderer interface.
func (h Height) RenderAttr() []byte {
	return []byte(html.Attr("height", fmt.Sprint(h)))
}

// Indicates how the browser should load the iframe:
const (
	// Loads the iframe immediately, regardless of whether or
	// not the iframe is currently within the visible viewport
	// (this is the default value).
	LoadingEager = html.Attribute("loading=eager")

	// Defers loading the iframe until it reaches a calculated
	// distance from the viewport, as defined by the browser.
	// The intent is to avoid the network and storage bandwidth
	// needed to handle the iframe until it's reasonably certain
	// that it will be needed. This generally improves the
	// performance of the content in most typical use cases.
	LoadingLazy = html.Attribute("loading=lazy")
)

// Name for the embedded browsing context. This can be used in the
// target attribute of the <a>, <form>, or <base> elements; the
// formtarget attribute of the <input> or <button> elements; or the
// windowName parameter in the window.open() method.
type Name string

// RenderAttr implements the attr.Renderer interface.
func (n Name) RenderAttr() []byte {
	return []byte(html.Attr("name", string(n)))
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

// Allowance for Sandbox.
type Allowance string

// Applies extra restrictions to the content in the frame. The value of
// the attribute can either be empty to apply all restrictions, or
// Allowances to lift particular restrictions.
type Sandbox []Allowance

const (
	// Allows for downloads to occur without a gesture from the user.
	AllowDownloadsWithoutUserActivation = Allowance("allow-downloads-without-user-activation")

	// Allows for downloads to occur with a gesture from the user.
	AllowDownloads = Allowance("allow-downloads")

	// Allows the resource to submit forms. If this keyword is not used,
	// form submission is blocked.
	AllowForms = Allowance("allow-forms")

	// Lets the resource open modal windows.
	AllowModals = Allowance("allow-modals")

	// Lets the resource lock the screen orientation.
	AllowOrientationLock = Allowance("allow-orientation-lock")

	// Lets the resource use the Pointer Lock API.
	AllowPointerLock = Allowance("allow-pointer-lock")

	// Allows popups (such as window.open(), target="_blank",
	// or showModalDialog()). If this keyword is not used, the
	// popup will silently fail to open.
	AllowPopups = Allowance("allow-popups")

	// Lets the sandboxed document open new windows without those
	// windows inheriting the sandboxing. For example, this can
	// safely sandbox an advertisement without forcing the same
	// restrictions upon the page the ad links to.
	AllowPopupsToEscapeSandbox = Allowance("allow-popups-to-escape-sandbox")

	// Lets the resource start a presentation session.
	AllowPresentation = Allowance("allow-presentation")

	// If this token is not used, the resource is treated as being
	// from a special origin that always fails the same-origin policy
	// (potentially preventing access to data storage/cookies and some
	// JavaScript APIs).
	AllowSameOrigin = Allowance("allow-same-origin")

	// Lets the resource run scripts (but not create popup windows).
	AllowScripts = Allowance("allow-scripts")

	// Lets the resource request access to the parent's storage
	// capabilities with the Storage Access API.
	AllowStorageAccessByUserActivation = Allowance("allow-storage-access-by-user-activation")

	// Lets the resource navigate the top-level browsing context
	// (the one named _top).
	AllowTopNavigation = Allowance("allow-top-navigation")

	// Lets the resource navigate the top-level browsing context, but only if initiated by a user gesture.
	AllowTopNavigationByUserActivation = Allowance("allow-top-navigation-by-user-activation")
)

// RenderAttr implements the attr.Renderer interface.
func (a Sandbox) RenderAttr() []byte {
	var s []string
	for _, f := range a {
		s = append(s, string(f))
	}
	return []byte(html.Attr("sandbox", strings.Join(s, " ")))
}

// Source URL of the page to embed. Use a value of about:blank to embed
// an empty page that conforms to the same-origin policy. Also note that
// programmatically removing an <iframe>'s src attribute (e.g. via
// Element.removeAttribute()) causes about:blank to be loaded in the frame
// in Firefox (from version 65), Chromium-based browsers, and Safari/iOS.
type Source string

// RenderAttr implements the attr.Renderer interface.
func (s Source) RenderAttr() []byte {
	return []byte(html.Attr("src", string(s)))
}

// Inline HTML to embed, overriding the src attribute. If a browser does not
// support the srcdoc attribute, it will fall back to the URL in the src attribute.
type SourceDocument html.String

// RenderAttr implements the attr.Renderer interface.
func (s SourceDocument) RenderAttr() []byte {
	return []byte(html.Attr("srcdoc", string(s)))
}
