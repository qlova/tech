//Package provides a core GLSL compiler.
package glsl

import (
	"bytes"
	"fmt"

	"qlova.tech/dsl"
	"qlova.tech/dsl/dslutil"
)

type Source struct {
	dslutil.State

	vertexHead bytes.Buffer
	vertexBody bytes.Buffer

	fragmentHead bytes.Buffer
	fragmentBody bytes.Buffer
}

func (p *Source) Files() (vert []byte, frag []byte, err error) {
	vert = append(vert, p.vertexHead.Bytes()...)
	vert = append(vert, `

	mat4 transpose(mat4 m) {
		return mat4(m[0][0], m[1][0], m[2][0], m[3][0],
					m[0][1], m[1][1], m[2][1], m[3][1],
					m[0][2], m[1][2], m[2][2], m[3][2],
					m[0][3], m[1][3], m[2][3], m[3][3]);
	  }
	  

	mat4 inverse(mat4 m) {
		float
			a00 = m[0][0], a01 = m[0][1], a02 = m[0][2], a03 = m[0][3],
			a10 = m[1][0], a11 = m[1][1], a12 = m[1][2], a13 = m[1][3],
			a20 = m[2][0], a21 = m[2][1], a22 = m[2][2], a23 = m[2][3],
			a30 = m[3][0], a31 = m[3][1], a32 = m[3][2], a33 = m[3][3],
	  
			b00 = a00 * a11 - a01 * a10,
			b01 = a00 * a12 - a02 * a10,
			b02 = a00 * a13 - a03 * a10,
			b03 = a01 * a12 - a02 * a11,
			b04 = a01 * a13 - a03 * a11,
			b05 = a02 * a13 - a03 * a12,
			b06 = a20 * a31 - a21 * a30,
			b07 = a20 * a32 - a22 * a30,
			b08 = a20 * a33 - a23 * a30,
			b09 = a21 * a32 - a22 * a31,
			b10 = a21 * a33 - a23 * a31,
			b11 = a22 * a33 - a23 * a32,
	  
			det = b00 * b11 - b01 * b10 + b02 * b09 + b03 * b08 - b04 * b07 + b05 * b06;
	  
		return mat4(
			a11 * b11 - a12 * b10 + a13 * b09,
			a02 * b10 - a01 * b11 - a03 * b09,
			a31 * b05 - a32 * b04 + a33 * b03,
			a22 * b04 - a21 * b05 - a23 * b03,
			a12 * b08 - a10 * b11 - a13 * b07,
			a00 * b11 - a02 * b08 + a03 * b07,
			a32 * b02 - a30 * b05 - a33 * b01,
			a20 * b05 - a22 * b02 + a23 * b01,
			a10 * b10 - a11 * b08 + a13 * b06,
			a01 * b08 - a00 * b10 - a03 * b06,
			a30 * b04 - a31 * b02 + a33 * b00,
			a21 * b02 - a20 * b04 - a23 * b00,
			a11 * b07 - a10 * b09 - a12 * b06,
			a00 * b09 - a01 * b07 + a02 * b06,
			a31 * b01 - a30 * b03 - a32 * b00,
			a20 * b03 - a21 * b01 + a22 * b00) / det;
	  }
	`...)
	vert = append(vert, "\nvoid main() {\n"...)
	vert = append(vert, p.vertexBody.Bytes()...)
	vert = append(vert, "}\n"...)

	frag = append(frag, p.fragmentHead.Bytes()...)
	frag = append(frag, "\nvoid main() {\n"...)
	frag = append(frag, "vec4 lighting_albedo = vec4(0);\n"...)
	frag = append(frag, "vec4 lighting_normal = vec4(0,0,-1,0);\n"...)
	frag = append(frag, p.fragmentBody.Bytes()...)

	//Basic lighting shader.
	frag = append(frag, `
		{
			vec3 ambient = vec3(0.2, 0.2, 0.2);
			vec3 sun = normalize(vec3(0.0, 1.0, 1.0));
			vec3 norm = normalize(lighting_normal.xyz);
			vec3 lightDir = normalize(-sun);

			float diff = max(dot(norm, lightDir), 0.0);
			vec3 diffuse = vec3(1, 1, 1)*diff;
			
			vec3 result = (ambient+diffuse)*lighting_albedo.xyz;

			gl_FragColor = vec4(result, 1.0);
		}
	`...)

	frag = append(frag, "}\n"...)

	return
}

type stage struct {
	*dslutil.State

	Head *bytes.Buffer
	Main *bytes.Buffer

	argPrefix string
	outPrefix string

	fragment bool
}

// Cores returns GLSL cores based on GLSL version 110.
// Other version of GLSL can extend these cores.
func (s *Source) Cores() (vert, frag dsl.Core) {

	vert = core(stage{
		Head:      &s.vertexHead,
		Main:      &s.vertexBody,
		State:     &s.State,
		outPrefix: "frag_",
	})

	frag = core(stage{
		Head:      &s.fragmentHead,
		Main:      &s.fragmentBody,
		State:     &s.State,
		argPrefix: "frag_",
		fragment:  true,
	})

	return
}

func (s stage) newIfElseChain() dsl.IfElseChain {
	return dsl.IfElseChain{
		ElseIf: func(condition dsl.Bool, fn func()) dsl.IfElseChain {
			s.Main.WriteString("else if (")
			s.Main.WriteString(string(condition.Value))
			s.Main.WriteString(") {\n")
			fn()
			s.Main.WriteString("}\n")
			return s.newIfElseChain()
		},
		Else: func(fn func()) {
			s.Main.WriteString("else {\n")
			fn()
			s.Main.WriteString("}\n")
		},
	}
}

func core(s stage) dsl.Core {
	var core dsl.Core

	core.Var = s.NewDefiner(s.Main, s,
		"%v %v = %v;\n")

	//TODO fix this, integer types need to be flat.
	if s.fragment {
		core.In = dslutil.Attributes(s.Head, s,
			"varying %[2]v "+s.argPrefix+"%[1]v;\n", s.argPrefix+"%[1]v")
	} else {
		core.In = dslutil.Attributes(s.Head, s,
			"attribute %[2]v "+s.argPrefix+"%[1]v;\n", s.argPrefix+"%[1]v")
		core.Out = dslutil.Attributes(s.Head, s,
			"varying %[2]v "+s.outPrefix+"%[1]v;\n", s.outPrefix+"%[1]v")
	}

	core.Uniform = s.NewUniforms(s.Head, s,
		"uniform %[2]v %[1]v;\n", "%[1]v")
	core.Get = s.NewUniforms(s.Head, s,
		"uniform %[2]v %[1]v;\n", "%[1]v")

	core.Constructor = s.NewConstructor(s, dslutil.Constructor{
		Bool:  "%v",
		Int:   "%v",
		Uint:  "%v",
		Float: "%.4f",
		Vec2:  "vec2(%v, %v)",
		Vec3:  "vec4(%v, %v, %v, 1)",
		Vec4:  "vec4(%v, %v, %v, %v)",
		RGB:   "vec4(%v, %v, %v, %v)",
	})

	core.Set = s.NewSetter(s.Main, "%v = %v;\n")

	core.Position = s.NewVec3("gl_Position")
	core.Fragment = s.NewRGB("lighting_albedo")
	core.Normal = s.NewVec3("lighting_normal")

	core.If = func(condition dsl.Bool, fn func()) dsl.IfElseChain {
		s.Main.WriteString("if (")
		s.Main.WriteString(string(condition.Value))
		s.Main.WriteString(") {\n")
		fn()
		s.Main.WriteString("}\n")

		return s.newIfElseChain()
	}

	core.Range = func(start, end dsl.Int, fn func(dsl.Int)) {
		var name = s.GetVariableName()
		fmt.Fprintf(s.Main, "for (int %v = %v; %v < %v; %v++) {\n",
			name, start.Value, name, end.Value, name)
		fn(s.NewInt(name))
		s.Main.WriteString("}\n")
	}

	core.Discard = func() {
		s.Main.WriteString("discard;\n")
	}
	core.Break = func() {
		s.Main.WriteString("break;\n")
	}
	core.Continue = func() {
		s.Main.WriteString("continue;\n")
	}
	core.While = func(condition dsl.Bool, fn func()) {
		s.Main.WriteString("while (")
		s.Main.WriteString(string(condition.Value))
		s.Main.WriteString(") {\n")
		fn()
		s.Main.WriteString("}\n")
	}

	return core
}
