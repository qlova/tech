package dsl_test

import (
	"fmt"
	"testing"

	"qlova.tech/gpu"
	"qlova.tech/gpu/driver/opengl/glsl/glsl460"
)

func TestMain(t *testing.T) {

	var textured gpu.Textured

	b, core, _ := glsl460.Compile(textured.Vertex, nil)
	fmt.Println(string(b))

	b, _, _ = glsl460.Compile(textured.Fragment, core)
	fmt.Println(string(b))

}
