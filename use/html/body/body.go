/*
	Package head provides the HTML <body> element.

	The <body> HTML element represents the content of an HTML
	document. There can be only one <body> element in a document.

	https://developer.mozilla.org/en-US/docs/Web/HTML/Element/body
	by Mozilla Contributors is licensed under CC-BY-SA 2.5
*/
package body

import (
	"qlova.tech/new/tree"
	"qlova.tech/use/html"
	"qlova.tech/use/js"
)

// Tag is the html <body> tag.
const Tag = html.Tag("body")

// New returns a <body> html tree node.
func New(args ...any) tree.Node {
	return html.New(append(args, Tag)...)
}

// OnAfterPrint function to call after the user has printed the document.
func OnAfterPrint(fn js.Function) html.Attribute {
	return html.Attr("onafterprint", string(fn))
}

// OnBeforePrint function to call before the user has printed the document.
func OnBeforePrint(fn js.Function) html.Attribute {
	return html.Attr("onbeforeprint", string(fn))
}

// OnBeforeUnload function to call before the document is about to be unloaded.
func OnBeforeUnload(fn js.Function) html.Attribute {
	return html.Attr("onbeforeunload", string(fn))
}

// OnBlur function to call when the document loses focus.
func OnBlur(fn js.Function) html.Attribute {
	return html.Attr("onblur", string(fn))
}

// OnError function to call when the document fails to load properly.
func OnError(fn js.Function) html.Attribute {
	return html.Attr("onerror", string(fn))
}

// OnFocus function to call when the document receives focus.
func OnFocus(fn js.Function) html.Attribute {
	return html.Attr("onfocus", string(fn))
}

// OnHashChange function to call when the fragment identifier part (starting with
// the hash ('#') character) of the document's current address has changed.
func OnHashChange(fn js.Function) html.Attribute {
	return html.Attr("onhashchange", string(fn))
}

// OnLanguageChange function to call when the preferred languages changed.
func OnLanguageChange(fn js.Function) html.Attribute {
	return html.Attr("onlanguagechange", string(fn))
}

// OnLoad function to call when the document has finished loading.
func OnLoad(fn js.Function) html.Attribute {
	return html.Attr("onload", string(fn))
}

// OnMessage function to call when the document has received a message.
func OnMessage(fn js.Function) html.Attribute {
	return html.Attr("onmessage", string(fn))
}

// OnOffline function to call when network communication has failed.
func OnOffline(fn js.Function) html.Attribute {
	return html.Attr("onoffline", string(fn))
}

// OnOnline function to call when network communication has been restored.
func OnOnline(fn js.Function) html.Attribute {
	return html.Attr("ononline", string(fn))
}

// OnPopState function to call when the user has navigated session history.
func OnPopState(fn js.Function) html.Attribute {
	return html.Attr("onpopstate", string(fn))
}

// OnRedo function to call when the user has moved forward in undo transaction history.
func OnRedo(fn js.Function) html.Attribute {
	return html.Attr("onredo", string(fn))
}

// OnResize function to call when the document has been resized.
func OnResize(fn js.Function) html.Attribute {
	return html.Attr("onresize", string(fn))
}

// OnStorage function to call when the storage area has changed.
func OnStorage(fn js.Function) html.Attribute {
	return html.Attr("onstorage", string(fn))
}

// OnUndo function to call when the user has moved backward in undo transaction history.
func OnUndo(fn js.Function) html.Attribute {
	return html.Attr("onundo", string(fn))
}

// OnUnload function to call when the document is going away.
func OnUnload(fn js.Function) html.Attribute {
	return html.Attr("onunload", string(fn))
}
