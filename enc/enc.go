//Package enc defines encoding interfaces and encoding abstraction.
//Use this package if you need configurable or abstractable encoding.
package enc

import "io"

//TextFormat is a text format encoding.
type TextFormat interface {
	Format

	TextEncode(interface{}, io.Writer) error
	TextDecode(interface{}, io.Reader) error
}

//Format is an encoding format.
type Format interface {
	Format() string

	Encode(interface{}, io.Writer) error
	Decode(interface{}, io.Reader) error
}

type wrapper struct {
	name       string
	wrapReader func(io.Reader) io.Reader
	wrapWriter func(io.Writer) io.Writer
	format     Format
}

func (j wrapper) Format() string { return "xml" }

func (j wrapper) TextEncode(value interface{}, writer io.Writer) error {
	return j.Encode(value, writer)
}

func (j wrapper) TextDecode(value interface{}, reader io.Reader) error {
	return j.Decode(value, reader)
}

func (j wrapper) Encode(value interface{}, writer io.Writer) error {
	if closer, ok := writer.(io.Closer); ok {
		defer closer.Close()
	}

	return j.format.Encode(value, j.wrapWriter(writer))
}

func (j wrapper) Decode(value interface{}, reader io.Reader) error {
	return j.format.Decode(value, j.wrapReader(reader))
}
