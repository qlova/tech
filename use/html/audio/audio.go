/*
	Package audio provides the HTML <audio> element.

	The <audio> HTML element is used to embed sound content in
	documents. It may contain one or more audio sources,
	represented using the src attribute or the <source> element:
	the browser will choose the most suitable one. It can also
	be the destination for streamed media, using a MediaStream.

	The above example shows simple usage of the <audio> element.
	In a similar manner to the <img> element, we include a path
	to the media we want to embed inside the src attribute; we
	can include other attributes to specify information such as
	whether we want it to autoplay and loop, whether we want to
	show the browser's default audio controls, etc.

	The content inside the opening and closing <audio></audio>
    tags is shown as a fallback in browsers that don't support
	the element.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/audio
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package audio

import (
	"qlova.tech/use/html"
	"qlova.tech/use/html/attributes"
	"qlova.tech/web/tree"
)

// Tag is the HTML <audio> element.
const Tag = html.Tag("audio")

// New returns an HTML <audio> tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
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

const (
	// Autoplay if specified, the audio will automatically begin
	// playback as soon as it can do so, without waiting for the
	// entire audio file to finish downloading.
	Autoplay = attributes.String("autoplay")

	// DisableRemotePlayback is used to disable the capability of
	// remote playback in devices that are attached using wired
	// (HDMI, DVI, etc.) and wireless technologies (Miracast,
	// Chromecast, DLNA, AirPlay, etc). See this proposed
	// specification for more information.
	DisableRemotePlayback = attributes.String("disableremoteplayback")

	// Loop if specified, the audio player will automatically
	// seek back to the start upon reaching the end of the audio.
	Loop = attributes.String("loop")

	// Muted indicates whether the audio will be initially
	// silenced. Its default value is false.
	Muted = attributes.String("muted")

	// Indicates that the audio should not be preloaded.
	PreloadNone = attributes.String("preload=none")

	// Indicates that only audio metadata (e.g. length) is fetched.
	PreloadMetadata = attributes.String("preload=metadata")

	// Indicates that the whole audio file can be downloaded,
	// even if the user is not expected to use it.
	PreloadAuto = attributes.String("preload=auto")
)

// Source URL of the audio to embed. This is subject to HTTP
// access controls. This is optional; you may instead use the
// <source> element within the audio block to specify the
// audio to embed.
type Source string

// RenderAttr implements the attr.Renderer interface.
func (s Source) RenderAttr() []byte {
	return []byte(html.Attr("src", string(s)))
}
