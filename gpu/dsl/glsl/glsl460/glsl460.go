//Package glsl460 compiles DSL into GLSL version 460.
package glsl460

import (
	"qlova.tech/gpu/dsl"
	"qlova.tech/gpu/dsl/glsl"
)

type Program struct {
	glsl.Program
}

// Compile compiles the given vertex, fragment and shader functions into GLSL version 460.
func Compile(vertex, fragment, shader func(dsl.Core)) (vert, frag []byte, err error) {
	var program Program

	v, f, s := program.Cores()
	vertex(v)
	fragment(f)
	shader(s)

	return program.Shaders()
}
