package shader

import "qlova.tech/gpu/internal/gputype"

// Internal types available to shader core functions
// running on the GPU.
type (
	Bool  = gputype.Bool
	Int   = gputype.Int
	Uint  = gputype.Uint
	Float = gputype.Float
	Vec2  = gputype.Vec2
	Vec3  = gputype.Vec3
	Vec4  = gputype.Vec4

	Mat3 = gputype.Mat3
	Mat4 = gputype.Mat4

	RGB  = gputype.RGB
	RGBA = gputype.RGBA

	Sampler = gputype.Sampler
)
