/*
	Package video provides the HTML <video> element.

	The <video> HTML element embeds a media player which
	supports video playback into the document. You can
	use <video> for audio content as well, but the
	<audio> element may provide a more appropriate user
	experience.

	The above example shows simple usage of the <video>
	element. In a similar manner to the <img> element,
	we include a path to the media we want to display
	inside the src attribute; we can include other
	attributes to specify information such as video
	width and height, whether we want it to autoplay
	and loop, whether we want to show the browser's
	default video controls, etc.

	The content inside the opening and closing
	<video></video> tags is shown as a fallback in
	browsers that don't support the element.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/video
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package video

import (
	"fmt"

	"qlova.tech/new/tree"
	"qlova.tech/use/html"
	"qlova.tech/use/html/attributes"
)

// Tag is the html <video> tag.
const Tag = html.Tag("video")

// New returns a new <video> element.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}

const (
	/*
		A Boolean attribute; if specified, the video automatically
		begins to play back as soon as it can do so without stopping
		to finish loading the data.

		Note: Sites that automatically play audio (or videos with
		an audio track) can be an unpleasant experience for users,
		so should be avoided when possible. If you must offer
		autoplay functionality, you should make it opt-in (requiring
		a user to specifically enable it). However, this can be useful
		when creating media elements whose source will be set at a
		later time, under user control. See our autoplay guide for
		additional information about how to properly use autoplay.

		To disable video autoplay, autoplay="false" will not work;
		the video will autoplay if the attribute is there in the
		<video> tag at all. To remove autoplay, the attribute needs
		to be removed altogether.

		In some browsers (e.g. Chrome 70.0) autoplay doesn't work
		if no muted attribute is present.
	*/
	Autoplay = attributes.String("autoplay")

	// A Boolean attribute which if true indicates that the element
	// should automatically toggle picture-in-picture mode when the
	// user switches back and forth between this document and another
	// document or application.
	AutoPictureInPicture = attributes.String("autopictureinpicture")

	// If this attribute is present, the browser will offer controls
	// to allow the user to control video playback, including volume,
	// seeking, and pause/resume playback.
	Controls = attributes.String("controls")
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
	// Prevents the browser from suggesting a Picture-in-Picture
	// context menu or to request Picture-in-Picture automatically
	// in some cases.
	DisablePictureInPicture = attributes.String("disablepictureinpicture")

	// A Boolean attribute used to disable the capability of remote
	// playback in devices that are attached using wired (HDMI, DVI, etc.)
	// and wireless technologies (Miracast, Chromecast, DLNA, AirPlay, etc).
	DisableRemotePlayback = attributes.String("disableremoteplayback")
)

// Width of the video's display area
type Width uint

// RenderAttr implements the attr.Renderer interface.
func (w Width) RenderAttr() []byte {
	return []byte(html.Attr("width", fmt.Sprint(w)))
}

// Height of the video's display area,
type Height uint

// RenderAttr implements the attr.Renderer interface.
func (h Height) RenderAttr() []byte {
	return []byte(html.Attr("height", fmt.Sprint(h)))
}

const (
	// Loop if specified, the browser will automatically seek
	// back to the start upon reaching the end of the video.
	Loop = attributes.String("loop")

	// Muted indicates the default setting of the audio
	// contained in the video. If set, the audio will be
	// initially silenced. Its default value is false, meaning
	// that the audio will be played when the video is played.
	Muted = attributes.String("muted")

	// PlaysInline indicates that the video is to be played
	// "inline", that is within the element's playback area.
	// Note that the absence of this attribute does not imply
	// that the video will always be played in fullscreen.
	PlaysInline = attributes.String("playsinline")

	// Poster URL for an image to be shown while the video is
	// downloading. If this attribute isn't specified, nothing
	// is displayed until the first frame is available, then
	// the first frame is shown as the poster frame.
	Poster = attributes.String("poster")

	// Indicates that the video should not be preloaded.
	PreloadNone = attributes.String("preload=none")

	// Indicates that only video metadata (e.g. length) is fetched.
	PreloadMetadata = attributes.String("preload=metadata")

	// Indicates that the whole video file can be downloaded,
	// even if the user is not expected to use it.
	PreloadAuto = attributes.String("preload=auto")
)

// Source URL of the video to embed. This is optional; you may
// instead use the <source> element within the video block to
// specify the video to embed.
type Source string

// RenderAttr implements the attr.Renderer interface.
func (s Source) RenderAttr() []byte {
	return []byte(html.Attr("src", string(s)))
}
