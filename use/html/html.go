package html

import (
	"fmt"
	"strconv"
	"strings"

	"html"

	"qlova.tech/rgb"
	"qlova.tech/use/html/attributes"
	"qlova.tech/web/tree"
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
		case rgb.Color:
			args[i] = Attr("style", fmt.Sprintf("background-color: #%06x;", v))
		}
	}
	return append(args, Tag("html"))
}

func get[T any](node tree.Node) T {
	var empty T
	for _, arg := range node {
		if v, ok := arg.(T); ok {
			return v
		}
	}
	return empty
}

func renderAttributes(nodes []any, s *strings.Builder) {
	for _, arg := range nodes {
		if v, ok := arg.(attributes.Renderer); ok {
			s.WriteByte(' ')
			s.Write(v.RenderAttr())
		}
		if v, ok := arg.([]any); ok {
			renderAttributes(v, s)
		}
	}
}

func renderChildren(nodes []any, s *strings.Builder) {
	raw := get[String](nodes)
	if raw != "" {
		s.WriteString(string(raw))
	}

	for _, arg := range nodes {
		if v, ok := arg.(tree.Node); ok {
			s.WriteString(string(Render(v)))
		}
		if v, ok := arg.(Renderer); ok {
			s.Write(v.RenderHTML())
		}
		if v, ok := arg.([]any); ok {
			renderChildren(v, s)
		}
	}
}

func Render(html tree.Node) String {
	var s strings.Builder

	tag := string(get[Tag](html))
	if tag == "html" {
		s.WriteString("<!DOCTYPE html>")
	}

	var isVoid = get[void](html) == Void

	if tag == "" {
		return ""
	}

	s.WriteByte('<')
	s.WriteString(string(tag))
	renderAttributes(html, &s)
	s.WriteByte('>')

	if !isVoid {
		renderChildren(html, &s)
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
