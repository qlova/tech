package glsl

import (
	"fmt"

	"qlova.tech/gpu"
)

type Variables struct {
	*File
}

func (u Variables) Name() string {
	return fmt.Sprintf("var_%v", u.Variables+1)
}

type Setter[T gpu.Type] struct {
	*File
	Name string
}

func (s Setter[T]) Set(value T) {
	fmt.Fprintf(&s.Body, "%s = %v;\n", s.Name, value)
}

func variable[T Type, V gpu.Type](u Variables, kind, array string, val V) T {
	u.Variables++
	name := u.Name()
	fmt.Fprintf(&u.Head, "%s %s%s = %v;\n", kind, name, array, val)
	u.Variables = u.Variables + 1
	return NewValue[T](name)
}

func (v Variables) Bool(value gpu.Bool) gpu.Variable[gpu.Bool] {
	return variable[Bool](v, "bool", "", value)
}

func (v Variables) Size(value gpu.Size) gpu.Variable[gpu.Size] {
	return variable[Size](v, "float", "", value)
}

func (v Variables) RGBA(value gpu.RGBA) gpu.Variable[gpu.RGBA] {
	return variable[RGBA](v, "vec4", "", value)
}

func (v Variables) Vec2(value gpu.Vec2) gpu.Variable[gpu.Vec2] {
	return variable[Vec2](v, "vec2", "", value)
}

func (v Variables) Vec3(value gpu.Vec3) gpu.Variable[gpu.Vec3] {
	return variable[Vec3](v, "vec3", "", value)
}

func (v Variables) Rot2(value gpu.Rot2) gpu.Variable[gpu.Rot2] {
	return variable[Rot2](v, "vec3", "", value)
}

func (v Variables) Rot3(value gpu.Rot3) gpu.Variable[gpu.Rot3] {
	return variable[Rot3](v, "vec4", "", value)
}

func (v Variables) Mod2(value gpu.Mod2) gpu.Variable[gpu.Mod2] {
	return variable[Mod2](v, "vec3", "[2]", value)
}

func (v Variables) Mod3(value gpu.Mod3) gpu.Variable[gpu.Mod3] {
	return variable[Mod3](v, "vec4", "[2]", value)
}

func (v Variables) Mat3(value gpu.Mat3) gpu.Variable[gpu.Mat3] {
	return variable[Mat3](v, "mat3", "", value)
}

func (v Variables) Mat4(value gpu.Mat4) gpu.Variable[gpu.Mat4] {
	return variable[Mat4](v, "mat4", "", value)
}
