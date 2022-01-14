//Package glsl460 compiles DSL into GLSL version 460.
package glsl460

import (
	"qlova.tech/dsl"
	"qlova.tech/dsl/glsl"
)

type Source struct {
	glsl.Source
}

// Compile compiles the given vertex, fragment and shader functions into GLSL version 460.
func Compile(vertex, fragment dsl.Shader) (vert, frag []byte, err error) {
	var source Source

	v, f := source.Cores()
	vertex(v)
	fragment(f)

	return source.Files()
}
