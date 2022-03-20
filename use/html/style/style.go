/*
	Package style provides the HTML <style> element.

	The <style> HTML element contains style information for a document,
	or part of a document. It contains CSS, which is applied to the
	contents of the document containing the <style> element.

	The <style> element must be included inside the <head> of the document.
	In general, it is better to put your styles in external stylesheets
	and apply them using <link> elements.

	If you include multiple <style> and <link> elements in your document,
	they will be applied to the DOM in the order they are included in
	the document — make sure you include them in the correct order,
	to avoid unexpected cascade issues.

	In the same manner as <link> elements, <style> elements can include
	media attributes that contain media queries, allowing you to
	selectively apply internal stylesheets to your document depending
	on media features such as viewport width.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/style
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package style

import (
	"qlova.tech/new/tree"
	"qlova.tech/use/css"
	"qlova.tech/use/html"
)

// Tag is the html <style> tag.
const Tag = html.Tag("style")

// New returns a <style> html tree node.
// converting any literal strings into
// css.String.
func New(args ...any) tree.Node {
	for i, arg := range args {
		if v, ok := arg.(string); ok {
			args[i] = css.String(v)
		}
	}

	return html.New(append(args, Tag)...)
}

// Media defines which media the style should be applied to.
func Media(query css.MediaQuery) html.Attribute {
	return html.Attr("media", string(query))
}

// Nonce is a a cryptographic nonce (number used once) used to allow inline
// styles in a style-src Content-Security-Policy. The server must generate
// a unique nonce value each time it transmits a policy. It is critical to
// provide a nonce that cannot be guessed as bypassing a resource's policy
// is otherwise trivial.
func Nonce(nonce string) html.Attribute {
	return html.Attr("nonce", nonce)
}

// Title can be used to specify alternative style sheet sets.
func Title(title string) html.Attribute {
	return html.Attr("title", title)
}
