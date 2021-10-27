package gpu

import (
	"go/ast"
	"go/types"
)

//Renderer is a shader written in Go.
type Renderer interface {
	SourceProgram() (string, *Program)
	Vertex() Vec4
	Fragment() RGBA
}

//Shader is a description of a program to run on the GPU.
//this struct is generally only important to drivers so
//that they can compile the program's shaders.
type Shader struct {
	Variables types.Struct //aka uniforms & buffers

	VertexShader struct {
		Attributes types.Struct
		Main       *ast.BlockStmt
	}
	FragmentShader struct {
		Attributes types.Struct
		Main       *ast.BlockStmt
	}
}

func Compile(r Renderer) {
	panic("not implemented") //TODO send to driver.
}
