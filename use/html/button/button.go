/*
	Package button provides the HTML <button> element.

	The <button> HTML element is an interactive element activated
	by a user with a mouse, keyboard, finger, voice command, or
	other assistive technology. Once activated, it then performs
	a programmable action, such as submitting a form or opening
	a dialog.

	By default, HTML buttons are presented in a style resembling
	the platform the user agent runs on, but you can change
	buttons' appearance with CSS.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/button
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package button

import (
	"strings"

	"qlova.tech/use/html"
	"qlova.tech/use/html/link"
	"qlova.tech/use/js"
	"qlova.tech/web/tree"
)

// Tag is the html <button> tag.
const Tag = html.Tag("button")

// New returns an html <button> tree node.
func New(args ...any) tree.Node {
	var scripts []js.Renderer
	for _, arg := range args {
		if script, ok := arg.(js.Renderer); ok {
			scripts = append(scripts, script)
		}
	}
	args = append(args, OnClick(scripts...))

	return html.New(append(args, Tag)...)
}

func OnClick(actions ...js.Renderer) html.Attribute {
	var s []string
	for _, action := range actions {
		s = append(s, string(action.RenderJS()))
	}
	return html.Attr("onclick", strings.Join(s, ";"))
}

const (
	// This Boolean attribute specifies that the button should have
	// input focus when the page loads. Only one element in a
	// document can have this attribute.
	AutoFocus = html.Attribute("autofocus")

	// This Boolean attribute prevents the user from interacting
	// with the button: it cannot be pressed or focused.
	Disabled = html.Attribute("disabled")
)

// The <form> element to associate the button with (its form owner).
// The value of this attribute must be the id of a <form> in the
// same document. (If this attribute is not set, the <button> is
// associated with its ancestor <form> element, if any.)
//
// This attribute lets you associate <button> elements to <form>s
// anywhere in the document, not just inside a <form>. It can also
// override an ancestor <form> element.
func Form(id html.ID) html.Attribute {
	return html.Attr("form", string(id))
}

// FormAction is a URL that processes the information submitted
// by the button. Overrides the action attribute of the button's
// form owner. Does nothing if there is no form owner.
type FormAction string

/*
	If the button is a submit button (it's inside/associated with a
	<form> and doesn't have type="button"), specifies how to encode
	the form data that is submitted.
*/
const (
	FormEncodingDefault   = html.Attribute("formenctype=application/x-www-form-urlencoded")
	FormEncodingMultipart = html.Attribute("formenctype=multipart/form-data")
	FormEncodingTextPlain = html.Attribute("formenctype=text/plain")
)

/*
	If the button is a submit button (it's inside/associated with a
	<form> and doesn't have type="button"), this attribute specifies
	the HTTP method used to submit the form.
*/
const (
	/*
		The data from the form are included in the body of the HTTP
		request when sent to the server. Use when the form contains
		information that shouldn't be public, like login credentials.
	*/
	FormMethodPost = html.Attribute("formmethod=post")

	/*
		The form data are appended to the form's action URL, with a ?
		as a separator, and the resulting URL is sent to the server.
		Use this method when the form has no side effects, like search forms.
	*/
	FormMethodGet = html.Attribute("formmethod=get")
)

/*
	If the button is a submit button, this Boolean attribute specifies
	that the form is not to be validated when it is submitted. If this
	attribute is specified, it overrides the novalidate attribute
	of the button's form owner.
*/
const FormNoValidate = html.Attribute("formnovalidate")

/*
	If the button is a submit button, this attribute is an author-defined
	name or standardized, underscore-prefixed keyword indicating where to
	display the response from submitting the form. This is the name of, or
	keyword for, a browsing context (a tab, window, or <iframe>). If this
	attribute is specified, it overrides the target attribute of the
	button's form owner.
*/
func FormTarget(target link.Target) html.Attribute {
	return html.Attr("formtarget", string(target))
}

/*
	The name of the button, submitted as a pair with the button's value
	as part of the form data, when that button is used to submit the form.
*/
type Name string

// RenderAttr implements the attributes.Renderer interface.
func (n Name) RenderAttr() []byte {
	return []byte(html.Attr("name", string(n)))
}

/*
	Value associated with the button's name when it's submitted with the
	form data. This value is passed to the server in params when the form
	is submitted using this button.
*/
type Value string

// RenderAttr implements the attributes.Renderer interface.
func (v Value) RenderAttr() []byte {
	return []byte(html.Attr("value", string(v)))
}
