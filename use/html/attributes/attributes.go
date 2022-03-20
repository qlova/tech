package attributes

// String is a string containing an HTML attribute.
type String string

// Renderer is any declarative element that can be render itself as an HTML attribute.
type Renderer interface {
	RenderAttr() []byte
}
