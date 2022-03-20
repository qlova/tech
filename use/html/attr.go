package html

import (
	"strconv"
	"strings"

	"qlova.tech/use/css"
)

// AccessKey provides a hint for generating a keyboard shortcut for
// the current element. This attribute consists of a space-separated
// list of characters. The browser should use the first one that
// exists on the computer keyboard layout.
func AccessKey(keys ...rune) Attribute {
	var s []string
	for _, key := range keys {
		s = append(s, string(key))
	}
	return Attr("accesskey", strings.Join(s, " "))
}

// Controls whether and how text input is automatically capitalized
// as it is entered/edited by the user.
const (
	AutoCapitalizeOff        = Attribute("autocapitalize=off")
	AutoCapitalizeOn         = Attribute("autocapitalize=on")
	AutoCapitalizeNone       = Attribute("autocapitalize=none")
	AutoCapitalizeSentences  = Attribute("autocapitalize=sentences")
	AutoCapitalizeWords      = Attribute("autocapitalize=words")
	AutoCapitalizeCharacters = Attribute("autocapitalize=characters")
)

// AutoFocus indicates that an element is to be focused on page load,
// or as soon as the <dialog> it is part of is displayed. This
// attribute is a boolean, initially false.
const AutoFocus = Attribute("autofocus")

// Class for an element.
type Class string

// RenderAttr implements the attributes.Renderer interface.
func (c Class) RenderAttr() []byte {
	return []byte(Attr("class", string(c)))
}

type ContentEditable bool

// RenderAttr implements the attributes.Renderer interface.
func (c ContentEditable) RenderAttr() []byte {
	return []byte(Attr("contenteditable", strconv.FormatBool(bool(c))))
}

// Data allows proprietary information to be exchanged between the
// HTML and its DOM representation that may be used by scripts. All
// such custom data are available via the HTMLElement interface of t
// he element the attribute is set on. The HTMLElement.dataset
// property gives access to them.
func Data(key, value string) Attribute {
	return Attr("data-"+key, value)
}

// An enumerated attribute indicating the directionality of the element's text.
const (
	DirectionRightToLeft = Attribute("dir=rtl")
	DirectionLeftToRight = Attribute("dir=ltr")
	DirectionAuto        = Attribute("dir=auto")
)

// Draggable indicates indicating whether the element can be dragged, using
// the Drag and Drop API
type Draggable bool

// RenderAttr implements the attributes.Renderer interface.
func (d Draggable) RenderAttr() []byte {
	return []byte(Attr("draggable", strconv.FormatBool(bool(d))))
}

// EnterKeyHint hints what action label (or icon) to present
// for the enter key on virtual keyboards.
type EnterKeyHint string

// RenderAttr implements the attributes.Renderer interface.
func (d EnterKeyHint) RenderAttr() []byte {
	return []byte(Attr("enterkeyhint", string(d)))
}

// Hidden indicates that the element is not yet, or is no longer,
// relevant. For example, it can be used to hide elements of the page
// that can't be used until the login process has been completed. The
// browser won't render such elements. This attribute must not be
// used to hide content that could legitimately be shown.
const Hidden = Attribute("hidden")

// ID defines a unique identifier (ID) which must be unique in the
// whole document. Its purpose is to identify the element when
// linking (using a fragment identifier), scripting, or styling (with CSS).
type ID string

// InputMode provides a hint to browsers as to the type of virtual
// keyboard configuration to use when editing this element or its contents.
// Used primarily on <input> elements, but is usable on any element while
// in contenteditable mode.
type InputMode string

// RenderAttr implements the attributes.Renderer interface.
func (i InputMode) RenderAttr() []byte {
	return []byte(Attr("inputmode", string(i)))
}

// Is allows you to specify that a standard HTML element should behave
// like a registered custom built-in element (see Using custom elements
// for more details).
func Is(elem string) Attribute {
	return Attr("is", elem)
}

// ItemID is the unique, global identifier of an item.
type ItemID string

// RenderAttr implements the attributes.Renderer interface.
func (i ItemID) RenderAttr() []byte {
	return []byte(Attr("itemid", string(i)))
}

// ItemProperty is used to add properties to an item. Every
// HTML element may have an itemprop attribute specified,
// where an ItemProperty consists of a name and value pair.
type ItemProperty string

// RenderAttr implements the attributes.Renderer interface.
func (i ItemProperty) RenderAttr() []byte {
	return []byte(Attr("itemprop", string(i)))
}

// ItemReferences that are not descendants of an element with
// the itemscope attribute can be associated with the item using
// an ItemReference.
type ItemReferences []ID

// RenderAttr implements the attributes.Renderer interface.
func (i ItemReferences) RenderAttr() []byte {
	var s []string
	for _, id := range i {
		s = append(s, string(id))
	}
	return []byte(Attr("itemref", strings.Join(s, " ")))
}

// ItemScope (usually) works along with itemtype to specify that
// the HTML contained in a block is about a particular item.
// itemscope creates the Item and defines the scope of the
// itemtype associated with it. itemtype is a valid URL of a
// vocabulary (such as schema.org) that describes the item and
// its properties context.
type ItemScope string

// RenderAttr implements the attributes.Renderer interface.
func (i ItemScope) RenderAttr() []byte {
	return []byte(Attr("itemscope", string(i)))
}

// ItemType Specifies the URL of the vocabulary that will be used
// to define itemprops (item properties) in the data structure.
// itemscope is used to set the scope of where in the data
// structure the vocabulary set by itemtype will be active.
type ItemType string

// RenderAttr implements the attributes.Renderer interface.
func (i ItemType) RenderAttr() []byte {
	return []byte(Attr("itemtype", string(i)))
}

// Language helps define the language of an element: the language
// that non-editable elements are in, or the language that editable
// elements should be written in by the user. The attribute contains
// one "language tag" (made of hyphen-separated "language subtags")
// in the format defined in RFC 5646: Tags for Identifying Languages
// (also known as BCP 47). xml:lang has priority over it.
type Language string

// RenderAttr implements the attributes.Renderer interface.
func (l Language) RenderAttr() []byte {
	return []byte(Attr("lang", string(l)))
}

// Nonce is a cryptographic nonce ("number used once") which can be
// used by Content Security Policy to determine whether or not a given
// fetch will be allowed to proceed.
type Nonce string

// RenderAttr implements the attributes.Renderer interface.
func (n Nonce) RenderAttr() []byte {
	return []byte(Attr("nonce", string(n)))
}

// Parts allows CSS to select and style specific elements in
// a shadow tree via the ::part pseudo-element.
func Parts(parts ...string) Attribute {
	return Attr("part", strings.Join(parts, " "))
}

// Slot Assigns a slot in a shadow DOM shadow tree to an element:
// An element with a slot attribute is assigned to the slot created
// by the <slot> element whose name attribute's value matches that
// slot attribute's value.
func Slot(name string) Attribute {
	return Attr("slot", name)
}

// Spellcheck defines whether the element may be checked for spelling errors.
type Spellcheck bool

// RenderAttr implements the attributes.Renderer interface.
func (s Spellcheck) RenderAttr() []byte {
	return []byte(Attr("spellcheck", strconv.FormatBool(bool(s))))
}

// Style contains CSS styling declarations to be applied to the element.
// Note that it is recommended for styles to be defined in a separate
// file or files. This attribute and the <style> element have mainly
// the purpose of allowing for quick styling, for example for testing purposes.
func Style(rules css.String) Attribute {
	return Attr("style", string(rules))
}

/*
	TabIndex indicating if the element can take input focus (is focusable),
	if it should participate to sequential keyboard navigation, and if so,
	at what position. It can take several values:

		- a negative value means that the element should be focusable,
		  but should not be reachable via sequential keyboard navigation;
		- 0 means that the element should be focusable and reachable via
		  sequential keyboard navigation, but its relative order is
		  defined by the platform convention;
		- a positive value means that the element should be focusable
		  and reachable via sequential keyboard navigation; the order
		  in which the elements are focused is the increasing value of
		  the tabindex. If several elements share the same tabindex,
		  their relative order follows their relative positions in
		  the document.
*/
type TabIndex int

// RenderAttr implements the attributes.Renderer interface.
func (t TabIndex) RenderAttr() []byte {
	return []byte(Attr("tabindex", strconv.Itoa(int(t))))
}

// Title contains a text representing advisory information related to
// the element it belongs to. Such information can typically, but not
// necessarily, be presented to the user as a tooltip.
type Title string

// RenderAttr implements the attributes.Renderer interface.
func (t Title) RenderAttr() []byte {
	return []byte(Attr("title", string(t)))
}

// Translate that is used to specify whether an element's attribute
// values and the values of its Text node children are to be translated
// when the page is localized, or whether to leave them unchanged.
type Translate bool

// RenderAttr implements the attributes.Renderer interface.
func (t Translate) RenderAttr() []byte {
	if t {
		return []byte(Attr("translate", "yes"))
	}
	return []byte(Attr("translate", "no"))
}
