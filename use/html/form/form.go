/*
	Package form provides the HTML <form> element.

	The <form> HTML element represents a document section
	containing interactive controls for submitting information.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/form
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package form

import (
	"strings"

	"qlova.tech/new/tree"
	"qlova.tech/use/html"
	"qlova.tech/use/html/link"
)

// Tag is the html <form> tag.
const Tag = html.Tag("form")

// New returns an html <form> tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}

// AcceptCharset sets the character encodings the server accepts. The
// browser uses them in the order in which they are listed. The
// default value means the same encoding as the page.
func AcceptCharset(encodings ...string) html.Attribute {
	return html.Attr("accept-charset", strings.Join(encodings, " "))
}

// AutoComplete indicates whether input elements can by default have
// their values automatically completed by the browser. autocomplete
// attributes on form elements override it on <form>.
func AutoComplete(value bool) html.Attribute {
	if value {
		return html.Attr("autocomplete", "on")
	}
	return html.Attr("autocomplete", "off")
}

/*
	Name of the form. The value must not be the empty string, and
	must be unique among the form elements in the forms collection
	that it is in, if any.
*/
type Name string

// RenderAttr implements the attributes.Renderer interface.
func (n Name) RenderAttr() []byte {
	return []byte(html.Attr("name", string(n)))
}

// Rel creates a hyperlink or annotation depending on the value,
// see the link.Type for details.
func Rel(relationship link.Type) html.Attribute {
	return html.Attr("rel", string(relationship))
}

/*
	The URL that processes the form submission. This value can be
	overridden by a formaction attribute on a <button>,
	<input type="submit">, or <input type="image"> element. This
	attribute is ignored when method="dialog" is set.
*/
type Action string

// RenderAttr implements the attributes.Renderer interface.
func (a Action) RenderAttr() []byte {
	return []byte(html.Attr("action", string(a)))
}

/*
	If the value of the method attribute is post, enctype is the
	MIME type of the form submission.
*/
const (
	EncodingDefault   = html.Attribute("enctype=application/x-www-form-urlencoded")
	EncodingMultipart = html.Attribute("enctype=multipart/form-data")
	EncodingTextPlain = html.Attribute("enctype=text/plain")
)

// The HTTP method to submit the form with.
const (
	MethodPost   = html.Attribute("method=post")
	MethodGet    = html.Attribute("method=get")
	MethodDialog = html.Attribute("method=dialog")
)

/*
	NoValidate indicates that the form shouldn't be validated when
	submitted. If this attribute is not set (and therefore the
	form is validated), it can be overridden by a formnovalidate
	attribute on a <button>, <input type="submit">, or
	<input type="image"> element belonging to the form.
*/
const NoValidate = html.Attribute("novalidate")

/*
	Indicates where to display the response after submitting the form.
	In HTML 4, this is the name/keyword for a frame. In HTML5, it
	is a name/keyword for a browsing context (for example, tab,
	window, or iframe).
*/
func Target(target link.Target) html.Attribute {
	return html.Attr("target", string(target))
}
