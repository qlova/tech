/*
	Package table provides the HTML <table> element.

	The <table> HTML element represents tabular data — that is,
	information presented in a two-dimensional table comprised
	of rows and columns of cells containing data.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/table
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package table

import (
	"qlova.tech/use/html"
	"qlova.tech/web/tree"
)

// Tag is the html <table> tag.
const Tag = html.Tag("table")

// New returns an html <table> tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}
