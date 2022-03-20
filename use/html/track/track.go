/*
	Package track provides the HTML <track> element.

	The <track> HTML element is used as a child of the media
	elements, <audio> and <video>. It lets you specify timed
	text tracks (or time-based data), for example to
	automatically handle subtitles. The tracks are formatted
	in WebVTT format (.vtt files) — Web Video Text Tracks.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/track
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package track

import (
	"qlova.tech/new/tree"
	"qlova.tech/use/html"
)

// Tag is the html <track> tag.
const Tag = html.Tag("track")

// New returns a new <track> element.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag, html.Void)...)
}

// Default indicates that the track should be enabled unless the
// user's preferences indicate that another track is more appropriate.
// This may only be used on one track element per media element.
const Default = html.Attribute("default")

/*
	How the text track is meant to be used. If omitted the default kind
	is subtitles. If the attribute contains an invalid value, it will
	use metadata (Versions of Chrome earlier than 52 treated an invalid
	value as subtitles).
*/
const (
	/*
		- Subtitles provide translation of content that cannot be
		  understood by the viewer. For example speech or text that
		  is not English in an English language film.

		- Subtitles may contain additional content, usually extra
		  background information. For example the text at the
		  beginning of the Star Wars films, or the date, time, and
		  location of a scene.
	*/
	KindSubtitles = html.Attribute("kind=subtitles")

	/*
		- Closed captions provide a transcription and possibly a
		  translation of audio.

		- It may include important non-verbal information such as
		  music cues or sound effects. It may indicate the cue's
		  source (e.g. music, text, character).

		- Suitable for users who are deaf or when the sound is muted.
	*/
	KindCaptions = html.Attribute("kind=captions")

	/*
		- Textual description of the video content.

		- Suitable for users who are blind or where the video cannot
		  be seen.
	*/
	KindDescriptions = html.Attribute("kind=descriptions")

	// Chapter titles are intended to be used when the user is
	// navigating the media resource.
	KindChapters = html.Attribute("kind=chapters")

	// Tracks used by scripts. Not visible to the user.
	KindMetadata = html.Attribute("kind=metadata")
)

// Label is a user-readable title of the text track which is
// used by the browser when listing available text tracks.
type Label string

// RenderAttr implements the attr.Renderer interface.
func (l Label) RenderAttr() []byte {
	return []byte(html.Attr("label", string(l)))
}

// Source address of the track (.vtt file). Must be a valid URL.
// This attribute must be specified and its URL value must have
// the same origin as the document — unless the <audio> or <video>
// parent element of the track element has a crossorigin attribute.
type Source string

// RenderAttr implements the attr.Renderer interface.
func (s Source) RenderAttr() []byte {
	return []byte(html.Attr("src", string(s)))
}

// SourceLanguage of the track text data. It must be a valid BCP 47
// language tag. If the kind attribute is set to subtitles, then
// SourceLanguage must be defined.
type SourceLanguage string

// RenderAttr implements the attr.Renderer interface.
func (s SourceLanguage) RenderAttr() []byte {
	return []byte(html.Attr("srclang", string(s)))
}
