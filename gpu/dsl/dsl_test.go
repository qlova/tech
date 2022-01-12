package dsl_test

import (
	"fmt"
	"testing"

	"qlova.tech/gpu"
	"qlova.tech/gpu/dsl"
	"qlova.tech/gpu/dsl/glsl/glsl460"
)

func TestMain(t *testing.T) {

	var textured gpu.Textured

	vert, frag, _ := glsl460.Compile(textured.Vertex, textured.Fragment, func(c dsl.Core) {})
	fmt.Println(string(vert))
	fmt.Println(string(frag))

}
