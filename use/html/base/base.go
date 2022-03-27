/*
	Package base provides the HTML <base> element.

	The <base> HTML element specifies the base URL to use for all relative
	URLs in a document. There can be only one <base> element in a document.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/base
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package base

import (
	"qlova.tech/use/html"
	"qlova.tech/use/html/link"
	"qlova.tech/web/tree"
)

// Tag is the html <base> tag.
const Tag = html.Tag("base")

// New returns a <base> html tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag, html.Void)...)
}

// Href is the base URL to be used throughout the document for relative URLs.
// Absolute and relative URLs are allowed.
func Href(href string) html.Attribute {
	return html.Attr("href", href)
}

// Target defines a keyword or author-defined name of the default browsing context
// to show the results of navigation from <a>, <area>, or <form> elements without
// explicit target attributes.
func Target(target link.Target) html.Attribute {
	return html.Attr("target", string(target))
}
