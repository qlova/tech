package glsl

import (
	"fmt"
	"strings"
	"unsafe"

	"qlova.tech/gpu"
	"qlova.tech/rgb"
	"qlova.tech/xy"
	"qlova.tech/xyz"
)

func float(f float32) string {
	var s = fmt.Sprintf("%#v", f)
	if !strings.Contains(s, ".") {
		return s + ".0"
	}
	return s
}

type Constants struct {
	Uniforms
}

func (u Constants) Name() string {
	return fmt.Sprintf("const_%v", u.Constants+1)
}

func (u Constants) Bool(b *bool) gpu.Bool {
	u.Constants++
	return NewValue[Bool]("%v", *b)
}

func (u Constants) Size(f *float32) gpu.Size {
	u.Constants++
	return NewValue[Size]("%v", *f)
}

func (u Constants) RGBA(c *rgb.Color) gpu.RGBA {
	u.Constants++

	var r, g, b, a string

	r = float(float32(c.Red()) / 255)
	g = float(float32(c.Green()) / 255)
	b = float(float32(c.Blue()) / 255)
	a = float(float32(c.Alpha()) / 255)

	return NewValue[RGBA]("vec4(%v, %v, %v, %v)", r, g, b, a)
}

func (u Constants) Vec2(b *xy.Vector) gpu.Vec2 {
	u.Constants++
	return NewValue[Vec2]("vec2(%v, %v)", b.X, b.Y)
}

func (u Constants) Vec3(b *xyz.Vector) gpu.Vec3 {
	u.Constants++
	return NewValue[Vec3]("vec3(%v, %v, %v)", b.X, b.Y, b.Z)
}

func (u Constants) Rot2(b *xy.Rotation) gpu.Rot2 {
	return Rot2(u.Vec3((*xyz.Vector)(unsafe.Pointer(b))).(Vec3))
}

func (u Constants) Rot3(b *xyz.Rotation) gpu.Rot3 {
	u.Constants++
	q := (*[4]float32)(unsafe.Pointer(b))
	return NewValue[Rot3]("vec4(%v, %v, %v, %v)", q[0], q[1], q[2], q[3])
}

func (u Constants) Mod2(b *xy.Transform) gpu.Mod2 {
	u.Constants++
	q := (*[3]float32)(unsafe.Pointer(b))
	return NewValue[Mod2](`gpu_mod2(%v, %v, %v, %v, %v, %v)`,
		float(q[0]), float(q[1]), float(q[2]),
		float(b.Scale), float(b.Position.X), float(b.Position.Y),
	)
}

func (u Constants) Mod3(b *xyz.Transform) gpu.Mod3 {
	u.Constants++
	q := (*[4]float32)(unsafe.Pointer(b))
	return NewValue[Mod3](`gpu_mod3(%v, %v, %v, %v, %v, %v, %v, %v)`,
		float(q[0]), float(q[1]), float(q[2]), float(q[3]),
		float(b.Scale),
		float(b.Position.X), float(b.Position.Y), float(b.Position.Z),
	)
}

func (u Constants) Mat3(b *xy.TransformationMatrix) gpu.Mat3 {
	u.Constants++
	data := b.Array()
	var s strings.Builder
	fmt.Fprint(&s, "mat3")
	first := '('
	for i := range data {
		fmt.Fprintf(&s, "%c%v", first, float(data[i]))
		first = ','
	}
	fmt.Fprint(&s, ")")
	return NewValue[Mat3](s.String())
}

func (u Constants) Mat4(b *xyz.TransformationMatrix) gpu.Mat4 {
	u.Constants++
	data := b.Array()
	var s strings.Builder
	fmt.Fprint(&s, "mat4")
	first := '('
	for i := range data {
		fmt.Fprintf(&s, "%c%v", first, float(data[i]))
		first = ','
	}
	fmt.Fprint(&s, ")")
	return NewValue[Mat4](s.String())
}
