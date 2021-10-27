package gpu

//Vec2 is a vector of 2 floats.
//See the vec2 package for functions that operate on Vec2.
type Vec2 struct {
	X, Y float32
}

//Vec3 is a vector of 3 floats.
//See the vec3 package for functions that operate on Vec3.
type Vec3 struct {
	X, Y, Z float32
}

//Vec4 is a vector of 4 floats.
//See the vec4 package for functions that operate on Vec4.
type Vec4 struct {
	X, Y, Z, W float32
}

type RGB struct {
	R, G, B float32
}

type RGBA struct {
	R, G, B, A float32
}

//Mat3 is a 3x3 matrix.
//See the mat3 package for functions that operate on Mat4.
type Mat3 [9]float32

//Mat4 is a 4x4 matrix.
//See the mat4 package for functions that operate on Mat4.
type Mat4 [16]float32
