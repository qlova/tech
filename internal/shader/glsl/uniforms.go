package glsl

import (
	"fmt"

	"qlova.tech/gpu"
	"qlova.tech/rgb"
	"qlova.tech/xy"
	"qlova.tech/xyz"
)

type Uniform struct {
	Name    string
	Pointer interface{}
}

type Uniforms struct {
	*File
}

func (u Uniforms) Name() string {
	return fmt.Sprintf("uniform_%v", len(u.Uniforms)+1)
}

func uniform[T Type](u Uniforms, kind, array string, ptr interface{}) T {
	name := u.Name()
	fmt.Fprintf(&u.Head, "uniform %s %s%s;\n", kind, name, array)
	u.Uniforms = append(u.Uniforms, Uniform{Name: name, Pointer: ptr})
	return NewValue[T](name)
}

func (u Uniforms) Bool(b *bool) gpu.Bool {
	return uniform[Bool](u, "bool", "", b)
}

func (u Uniforms) Size(b *float32) gpu.Size {
	return uniform[Size](u, "float", "", b)
}

func (u Uniforms) RGBA(b *rgb.Color) gpu.RGBA {
	return uniform[RGBA](u, "vec4", "", b)
}

func (u Uniforms) Vec2(b *xy.Vector) gpu.Vec2 {
	return uniform[Vec2](u, "vec2", "", b)
}

func (u Uniforms) Vec3(b *xyz.Vector) gpu.Vec3 {
	return uniform[Vec3](u, "vec3", "", b)
}

func (u Uniforms) Rot2(b *xy.Rotation) gpu.Rot2 {
	return uniform[Rot2](u, "vec3", "", b)
}

func (u Uniforms) Rot3(b *xyz.Rotation) gpu.Rot3 {
	return uniform[Rot3](u, "vec4", "", b)
}

func (u Uniforms) Mod2(b *xy.Transform) gpu.Mod2 {
	return uniform[Mod2](u, "vec3", "[2]", b)
}

func (u Uniforms) Mod3(b *xyz.Transform) gpu.Mod3 {
	return uniform[Mod3](u, "vec4", "[2]", b)
}

func (u Uniforms) Mat3(b *xy.TransformationMatrix) gpu.Mat3 {
	return uniform[Mat3](u, "mat3", "", b)
}

func (u Uniforms) Mat4(b *xyz.TransformationMatrix) gpu.Mat4 {
	return uniform[Mat4](u, "mat4", "", b)
}

func (u Uniforms) Tex2(b *gpu.Texture) gpu.Tex2 {
	return uniform[Tex2](u, "sampler2D", "", b)
}

func (u Uniforms) Tex3(b *gpu.Texture) gpu.Tex3 {
	return uniform[Tex3](u, "sampler3D", "", b)
}

func (u Uniforms) Map3(b *gpu.Texture) gpu.Map3 {
	return uniform[Map3](u, "samplerCube", "", b)
}
