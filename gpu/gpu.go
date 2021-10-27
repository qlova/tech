package gpu

import _ "embed"

//go:embed gpu.go
var Source string

//Float is an alias to float32.
type Float = float32

//Vec2 is a vector of 2 floats.
//See the vec2 package for functions that operate on Vec2.
type Vec2 [2]Float

//Vec3 is a vector of 3 floats.
//See the vec3 package for functions that operate on Vec3.
type Vec3 [3]Float

//Vec4 is a vector of 4 floats.
//See the vec4 package for functions that operate on Vec4.
type Vec4 [4]Float

//RGB is a color represented by 3 floats.
//See the rgb package for functions that operate on RGB.
type RGB [3]Float

//RGBA is a color represented by 4 floats.
//See the rgba package for functions that operate on RGBA.
type RGBA [4]Float

//Mat3 is a 3x3 matrix.
//See the mat3 package for functions that operate on Mat4.
type Mat2 [2 * 2]Float

//Mat3 is a 3x3 matrix.
//See the mat3 package for functions that operate on Mat4.
type Mat3 [3 * 3]Float

//Mat4 is a 4x4 matrix.
//See the mat4 package for functions that operate on Mat4.
type Mat4 [4 * 4]Float

//Texture is a texture on the GPU.
type Texture Pointer

//Program is a pointer to a compiled shader on the GPU.
type Program Pointer

func (p Program) Draw(mesh Mesh) {
	//TODO
}

//Pointer is an opaque reference to a GPU memory location.
type Pointer struct {
	uint64
}
