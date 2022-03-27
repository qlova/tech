/*
	Package meta provides the HTML <meta> element.

	The <meta> HTML element represents metadata that cannot be represented
	by other HTML meta-related elements, like <base>, <link>, <script>,
	<style> or <title>.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/meta
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package meta

import (
	"bytes"
	"strconv"
	"strings"

	"qlova.tech/use/css"
	"qlova.tech/use/html"
)

// Tag is the name of the tag.
const Tag html.Tag = "meta"

// Element contains attributes to place in an HTML meta element.
// Will render as an HTML meta element.
type Element html.Attribute

// RenderHTML implements html.Renderer
func (elem Element) RenderHTML() []byte {
	return []byte("<meta " + string(elem) + ">")
}

/*
	This attribute declares the document's character encoding.
*/
const CharsetUTF8 = Element("charset=utf-8")

/*
	ContentSecurityPolicy allows page authors to define a content policy for the
	current page Content policies mostly specify allowed server origins and script
	endpoints which help guard against cross-site scripting attacks.
*/
func ContentSecurityPolicy(policy string) Element {
	return Data("content-security-policy", policy)
}

// Declares the MIME type and character encoding of the document.
const ContentTypeTextHTML = Element(html.Attribute("http-equiv=content-type") + " " + html.Attribute(`content="text/html; charset=utf-8"`))

// DefaultStyle sets the name of the default CSS style sheet set.
func DefaultStyle(name string) Element {
	return Data("default-style", name)
}

// Refresh sets the number of seconds until the page should be reloaded.
func Refresh(seconds uint) Element {
	return Data("refresh", strconv.FormatUint(uint64(seconds), 10))
}

// Redirect sets number of seconds until the page should redirect to another.
func Redirect(seconds uint, url string) Element {
	return Data("refresh", strconv.FormatUint(uint64(seconds), 10)+";url="+url)
}

// Data sets document metadata in terms of name-value pairs.
func Data(name, value string) Element {
	return Element(html.Attribute("name="+name) + " " + html.Attr("content", value))
}

// ApplicationName is the name of the application running in the web page.
func ApplicationName(name string) Element {
	return Data("application-name", name)
}

// Author is the name of the document's author.
func Author(name string) Element {
	return Data("author", name)
}

// Description a short and accurate summary of the content of the page.
// Several browsers, like Firefox and Opera, use this as the default
// description of bookmarked pages.
func Description(description string) Element {
	return Data("description", description)
}

// Generator is the identifier of the software that generated the page.
func Generator(generator string) Element {
	return Data("generator", generator)
}

// Keywords are words relevant to the page's content.
func Keywords(words ...string) Element {
	return Data("keywords", strings.Join(words, ","))
}

// controls the HTTP Referer header for to requests sent from the document.
const (
	NoReferrer              = Element("http-equiv=referrer content=no-referrer")
	NoReferrerWhenDowngrade = Element("http-equiv=referrer content=no-referrer-when-downgrade")

	ReferrerOrigin                      = Element("http-equiv=referrer content=origin")
	RefererOriginWhenCrossOrigin        = Element("http-equiv=referrer content=origin-when-cross-origin")
	RefererSameOrigin                   = Element("http-equiv=referrer content=same-origin")
	ReferrerStrictOrigin                = Element("http-equiv=referrer content=strict-origin")
	ReferrerStrictOriginWhenCrossOrigin = Element("http-equiv=referrer content=strict-origin-when-cross-origin")
	ReferrerUnsafeURL                   = Element("http-equiv=referrer content=unsafe-url")
)

// ThemeColor indicates a suggested color that user agents should use to
// customize the display of the page or of the surrounding user interface
func ThemeColor(color css.Color) Element {
	return Data("theme-color", string(color))
}

// ColorSchemeName is a name of a color scheme.
type ColorSchemeName string

const (
	// The document is unaware of color schemes and should be rendered
	// using the default color palette.
	Normal = ColorSchemeName("normal")

	// Indicates that the document only supports light mode, with a light
	// background and dark foreground colors. By specification, only dark
	// is not valid, because forcing a document to render in dark mode
	// when it isn't truly compatible with it can result in unreadable
	// content; all major browsers default to light mode if not otherwise
	// configured.
	OnlyLight = ColorSchemeName("only-light")

	Light = ColorSchemeName("light")
	Dark  = ColorSchemeName("dark")
)

// ColorScheme specifies one or more color schemes with which the document
// is compatible. The browser will use this information in tandem with the
// user's browser or device settings to determine what colors to use for
// everything from background and foregrounds to form controls and
// scrollbars. The primary use for ColorScheme is to indicate compatibility
// with—and order of preference for—light and dark color modes.
func ColorScheme(schemes ...ColorSchemeName) Element {
	var s []string
	for _, scheme := range schemes {
		s = append(s, string(scheme))
	}
	return Data("color-scheme", strings.Join(s, ","))
}

// ViewportFit value.
type ViewportFit string

const (
	// The auto value doesn't affect the initial layout viewport, and the
	// whole web page is viewable.
	Auto = ViewportFit("auto")

	// The cover value means that the viewport is scaled to fill the device
	// display. It is highly recommended to make use of the safe area inset
	// variables to ensure that important content doesn't end up outside
	// the display.
	Cover = ViewportFit("cover")

	// The contain value means that the viewport is scaled to fit the largest
	// rectangle inscribed within the display.
	Contain = ViewportFit("contain")
)

// Viewport gives hints about the size of the initial size of the viewport.
type Viewport struct {

	// Defines the pixel width of the viewport that you want the web site
	// to be rendered at.
	Width uint

	// Defines the height of the viewport. Not used by any browser.
	Height uint

	// Sets Width to the width of the device.
	DeviceWidth bool

	// Sets Height to the height of the device.
	DeviceHeight bool

	// Defines the ratio between the device width (device-width in portrait
	// mode or device-height in landscape mode) and the viewport size.
	InitialScale float64

	// Defines the maximum amount to zoom in. It must be greater or equal
	// to the minimum-scale or the behavior is undefined. Browser
	// settings can ignore this rule and iOS10+ ignores it by default.
	MaximumScale float64

	// Defines the minimum zoom level. It must be smaller or equal to the
	// maximum-scale or the behavior is undefined. Browser settings can
	// ignore this rule and iOS10+ ignores it by default.
	MinimumScale float64

	// If set to false, the user is not able to zoom in the webpage.
	// Browser settings can ignore this rule, and iOS10+ ignores it by default.
	UserScalable bool

	// ViewportFit option.
	Fit ViewportFit
}

// RenderHTML implements html.Renderer
func (v Viewport) RenderHTML() []byte {
	var buf bytes.Buffer
	buf.WriteString("<meta name=viewport content=")
	if v.Width > 0 {
		buf.WriteString("width=" + strconv.FormatUint(uint64(v.Width), 10) + ",")
	}
	if v.Height > 0 {
		buf.WriteString("height=" + strconv.FormatUint(uint64(v.Height), 10) + ",")
	}
	if v.DeviceWidth {
		buf.WriteString("width=device-width,")
	}
	if v.DeviceHeight {
		buf.WriteString("width=device-height,")
	}
	if v.InitialScale > 0 {
		buf.WriteString("initial-scale=" + strconv.FormatFloat(v.InitialScale, 'f', -1, 64) + ",")
	}
	if v.MaximumScale > 0 {
		buf.WriteString("maximum-scale=" + strconv.FormatFloat(v.MaximumScale, 'f', -1, 64) + ",")
	}
	if v.MinimumScale > 0 {
		buf.WriteString("minimum-scale=" + strconv.FormatFloat(v.MinimumScale, 'f', -1, 64) + ",")
	}
	if !v.UserScalable {
		buf.WriteString("user-scalable=no,")
	}
	if v.Fit != "" {
		buf.WriteString("viewport-fit=" + string(v.Fit) + ",")
	}
	buf.WriteString("/>")
	return buf.Bytes()
}

// Creator is the name of the creator of the document, such as an organization or
// institution. If there are more than one, several <meta> elements should be used.
func Creator(name string) Element {
	return Data("creator", name)
}

const (
	// Allows the robot to index the page (default).
	RobotsIndex = Element("name=robots content=index")

	// Requests the robot to not index the page.
	RobotsNoIndex = Element("name=robots content=noindex")

	// Allows the robot to follow the links on the page (default).
	RobotsFollow = Element("name=robots content=follow")

	// Requests the robot to not follow the links on the page.
	RobotsNoFollow = Element("name=robots content=nofollow")

	// Equivalent to RobotIndex + RobotFollow
	RobotsAll = Element("name=robots content=all")

	// Equivalent to RobotNoIndex + RobotNoFollow
	RobotsNone = Element("name=robots content=none")

	// Requests the search engine not to cache the page content.
	RobotsNoArchive = Element("name=robots content=noarchive")

	// Prevents displaying any description of the page in search engine results.
	RobotsNoSnippet = Element("name=robots content=nosnippet")

	// Requests this page not to appear as the referring page of an indexed image.
	RobotsNoImageIndex = Element("name=robots content=noimageindex")

	// Synonym of RobotsNoArchive.
	RobotsNoCache = Element("name=robots content=nocache")
)
