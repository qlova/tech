/*
	Package article provides the HTML <article> element.

	The <article> HTML element represents a self-contained composition in
	a document, page, application, or site, which is intended to be
	independently distributable or reusable (e.g., in syndication).
	Examples include: a forum post, a magazine or newspaper article, or a
	blog entry, a product card, a user-submitted comment, an interactive
	widget or gadget, or any other independent item of content.

	A given document can have multiple articles in it; for example,
	on a blog that shows the text of each article one after another
	as the reader scrolls, each post would be contained in an <article>
	element, possibly with one or more <section>s within.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/article
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package article

import (
	"qlova.tech/use/html"
	"qlova.tech/web/tree"
)

// Tag is the html <article> tag.
const Tag = html.Tag("article")

// New returns an html <article> tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}
