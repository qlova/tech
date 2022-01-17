package glsl

import (
	"fmt"

	"qlova.tech/gpu"
)

type Type interface {
	gpu.Type

	setValue(string)
}

type Value struct {
	*File

	string
}

func (v Value) GPU() {}

func (v Value) String() string {
	return v.string
}

func (v Value) setValue(to string) {
	v.string = to
}

func (v Value) setVariable(to string) {
	v.string = to
}

func NewValue[T Type](format string, args ...interface{}) T {
	var v T

	for _, arg := range args {
		if f, ok := arg.(float32); ok {
			arg = float(f)
		}
	}

	if len(args) > 0 {
		v.setValue(fmt.Sprintf(format, args...))
	} else {
		v.setValue(format)
	}
	return v
}

type Bool struct{ Value }

func (b Bool) Set(value gpu.Bool) {
	fmt.Fprintf(&b.Body, "%s = %s;", b.string, value)
}

type Size struct{ Value }

func (a Size) Set(b gpu.Size) {
	fmt.Fprintf(&a.Body, "%s = %s;", a, b)
}

func (a Size) LessThan(b gpu.Size) gpu.Bool {
	return NewValue[Bool]("(%s < %s)", a, b)
}

func (a Size) MoreThan(b gpu.Size) gpu.Bool {
	return NewValue[Bool]("(%s > %s)", a, b)
}

func (a Size) Plus(b gpu.Size) gpu.Size {
	return NewValue[Size]("(%s + %s)", a, b)
}

func (a Size) ScaledBy(b gpu.Size) gpu.Size {
	return NewValue[Size]("(%s * %s)", a, b)
}

type RGBA struct{ Value }

func (a RGBA) Set(b gpu.RGBA) {
	fmt.Fprintf(&a.Body, "%s = %s;", a, b)
}

func (a RGBA) Vec3() gpu.Vec3 {
	return NewValue[Vec3]("vec3(%s.xyz)", a)
}

type Vec2 struct{ Value }

func (a Vec2) Set(b gpu.Vec2) {
	fmt.Fprintf(&a.Body, "%s = %s;", a, b)
}

func (a Vec2) ScaledBy(b gpu.Size) gpu.Vec2 {
	return NewValue[Vec2]("(%s * %s)", a, b)
}

func (a Vec2) Dot(b gpu.Vec2) gpu.Size {
	return NewValue[Size]("dot(%s, %s)", a, b)
}

func (a Vec2) Distance(b gpu.Vec2) gpu.Size {
	return NewValue[Size]("distance(%s, %s)", a, b)
}

type Vec3 struct{ Value }

func (a Vec3) Set(b gpu.Vec3) {
	fmt.Fprintf(&a.Body, "%s = %s;", a, b)
}

func (a Vec3) ScaledBy(b gpu.Size) gpu.Vec3 {
	return NewValue[Vec3]("(%s * %s)", a, b)
}

func (a Vec3) Cross(b gpu.Vec3) gpu.Vec3 {
	return NewValue[Vec3]("cross(%s, %s)", a, b)
}

func (a Vec3) Dot(b gpu.Vec3) gpu.Size {
	return NewValue[Size]("dot(%s, %s)", a, b)
}

func (a Vec3) Distance(b gpu.Vec3) gpu.Size {
	return NewValue[Size]("distance(%s, %s)", a, b)
}

type Rot2 struct{ Value }

func (a Rot2) Set(b gpu.Rot2) {
	fmt.Fprintf(&a.Body, "%s = %s;", a, b)
}

func (a Rot2) Rotate(b gpu.Vec2) gpu.Vec2 {
	return NewValue[Vec2]("gpu_rotate2D(%s, %s)", a, b)
}

func (a Rot2) ScaledBy(b gpu.Size) gpu.Rot2 {
	return NewValue[Rot2]("(%s * %s)", a, b)
}

type Rot3 struct{ Value }

func (a Rot3) Set(b gpu.Rot3) {
	fmt.Fprintf(&a.Body, "%s = %s;", a, b)
}

func (a Rot3) Rotate(b gpu.Vec3) gpu.Vec3 {
	return NewValue[Vec3]("gpu_rotate3D(%s, %s)", a, b)
}

func (a Rot3) ScaledBy(b gpu.Size) gpu.Rot3 {
	return NewValue[Rot3]("(%s * %s)", a, b)
}

type Mod2 struct{ Value }

func (a Mod2) Set(b gpu.Mod2) {
	fmt.Fprintf(&a.Body, "%s = %s;", a, b)
}

func (a Mod2) Transform(b gpu.Vec2) gpu.Vec2 {
	return NewValue[Vec2]("gpu_transform2D(%s, %s)", a, b)
}

func (a Mod2) Mat3() gpu.Mat3 {
	return NewValue[Mat3]("gpu_mat3(%s)", a)
}

type Mod3 struct{ Value }

func (a Mod3) Set(b gpu.Mod3) {
	fmt.Fprintf(&a.Body, "%s = %s;", a, b)
}

func (a Mod3) Transform(b gpu.Vec3) gpu.Vec3 {
	return NewValue[Vec3]("gpu_transform3D(%s, %s)", a, b)
}

func (a Mod3) Mat4() gpu.Mat4 {
	return NewValue[Mat4]("gpu_mat4(%s)", a)
}

type Mat3 struct{ Value }

func (a Mat3) Set(b gpu.Mat3) {
	fmt.Fprintf(&a.Body, "%s = %s;", a, b)
}

func (a Mat3) Transform(b gpu.Vec2) gpu.Vec2 {
	return NewValue[Vec2]("(%s * %s)", a, b)
}

type Mat4 struct{ Value }

func (a Mat4) Set(b gpu.Mat4) {
	fmt.Fprintf(&a.Body, "%s = %s;", a, b)
}

func (a Mat4) Transform(b gpu.Vec3) gpu.Vec3 {
	return NewValue[Vec3]("(%s * %s)", a, b)
}

type Tex2 struct{ Value }

func (a Tex2) Sample(b gpu.Vec2) gpu.RGBA {
	return NewValue[RGBA]("texture2D(%s, %s)", a, b)
}

type Tex3 struct{ Value }

func (a Tex3) Sample(b gpu.Vec3) gpu.RGBA {
	return NewValue[RGBA]("texture3D(%s, %s)", a, b)
}

type Map3 struct{ Value }

func (a Map3) Sample(b gpu.Vec3) gpu.RGBA {
	return NewValue[RGBA]("textureCube(%s, %s)", a, b)
}
