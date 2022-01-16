package glsl110

import (
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
