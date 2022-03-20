package html

import (
	"strconv"
	"strings"

	"html"

	"qlova.tech/new/node"
	"qlova.tech/new/tree"
	"qlova.tech/use/html/attributes"
)

// Renderer is any declarative element that can be render itself as HTML.
type Renderer interface {
	RenderHTML() []byte
}

// Tag is an HTML tag.
type Tag string

type void string

// Void marks an element to omit a closing tag.
const Void void = "void"

// String is a string containing HTML.
type String string

// RenderAttr implements the attributes.Renderer interface.
func (id ID) RenderAttr() []byte {
	return []byte(Attr("id", string(id)))
}

// Attribute is an HTML attribute.
type Attribute = attributes.String

// Attr is shorthand for Attribute{Name: name, Value: value}
func Attr(name, value string) Attribute {
	return Attribute(name + "=" + strconv.Quote(value))
}

// Escape escapes special characters like "<" to become "&lt;".
// It escapes only five such characters: <, >, &, ' and ".
// Unescape(Escape(s)) == s always holds, but the converse
// isn't always true.
func Escape(text string) String {
	return String(html.EscapeString(text))
}

// Unescape unescapes entities like "&lt;" to become "<". It
// unescapes a larger range of entities than Escape escapes.
// For example, "&aacute;" unescapes to "á", as does "&#225;"
// and "&#xE1;". Unescape(Escape(s)) == s always holds, but
// the converse isn't always true.
func Unescape(text string) String {
	return String(html.UnescapeString(text))
}

func New(args ...any) tree.Node {
	for i, arg := range args {
		switch v := arg.(type) {
		case string:
			args[i] = String(v)
		}
	}
	return append(args, Tag("html"))
}

func Render(html tree.Node) String {
	var s strings.Builder

	tag := string(node.Get[Tag](html))
	if tag == "html" {
		s.WriteString("<!DOCTYPE html>")
	}

	var isVoid = node.Get[void](html) == Void

	if tag == "" {
		return ""
	}

	s.WriteByte('<')
	s.WriteString(string(tag))

	for _, arg := range html {
		if v, ok := arg.(Attribute); ok {
			s.WriteByte(' ')
			s.WriteString(string(v))
		}
		if v, ok := arg.(attributes.Renderer); ok {
			s.WriteByte(' ')
			s.Write(v.RenderAttr())
		}
	}

	s.WriteByte('>')

	if !isVoid {
		raw := node.Get[String](html)
		if raw != "" {
			s.WriteString(string(raw))
		}

		for _, arg := range html {
			if v, ok := arg.(tree.Node); ok {
				s.WriteString(string(Render(v)))
			}
			if v, ok := arg.(Renderer); ok {
				s.Write(v.RenderHTML())
			}
		}

		s.WriteByte('<')
		s.WriteByte('/')
		s.WriteString(string(tag))
		s.WriteByte('>')
	}

	return String(s.String())
}

// Main returns a <main> html tree node.
// This tag does not have its own package like
// the other tags because main is a reserved
// package name in Go.
func Main(args ...any) tree.Node {
	return New(append(args, Tag("main"))...)
}
