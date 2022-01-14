package dsl_test

import (
	"fmt"
	"testing"

	"qlova.tech/dsl/glsl/glsl460"
	"qlova.tech/gpu"
)

func TestMain(t *testing.T) {

	var textured gpu.Textured

	vert, frag, _ := glsl460.Compile(textured.Vertex, textured.Fragment)
	fmt.Println(string(vert))
	fmt.Println(string(frag))

}
