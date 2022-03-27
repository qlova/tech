package js

type Function string

type String string

func (s String) RenderJS() []byte {
	return []byte(string(s))
}

type Renderer interface {
	RenderJS() []byte
}
