package glsl110

import (
	"qlova.tech/dsl"
	"qlova.tech/dsl/glsl"
)

type Source struct {
	glsl.Source
}

func (s *Source) Files() (vert, frag []byte, err error) {
	vert, frag, err = s.Source.Files()
	if err != nil {
		return nil, nil, err
	}

	return append([]byte("#version 110\n"), vert...), append([]byte("#version 110\n"), frag...), nil
}

func Compile(vertex, fragment dsl.Shader) (vert, frag []byte, err error) {
	var source Source

	v, f := source.Cores()
	vertex(v)
	fragment(f)

	return source.Files()
}
