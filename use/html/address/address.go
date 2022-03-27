/*
	Package address provides the HTML <address> element.

	The <address> HTML element indicates that the enclosed HTML provides
	contact information for a person or people, or for an organization.

	The contact information provided by an <address> element's contents
	can take whatever form is appropriate for the context, and may include
	any type of contact information that is needed, such as a physical
	address, URL, email address, phone number, social media handle,
	geographic coordinates, and so forth. The <address> element should
	include the name of the person, people, or organization to which
	the contact information refers.

	<address> can be used in a variety of contexts, such as providing a
	business's contact information in the page header, or indicating the
	author of an article by including an <address> element within the
	<article>.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/address
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package address

import (
	"qlova.tech/use/html"
	"qlova.tech/web/tree"
)

// Tag is the html <address> tag.
const Tag = html.Tag("address")

// New returns an html <address> tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}
