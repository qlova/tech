package shader

import (
	"qlova.tech/gpu"
	"qlova.tech/gpu/driver/opengl/glsl"
)

//Colored material colors everything the given color.
type Colored struct {
	Color gpu.Vec3
}

func (c *Colored) Variables() []interface{} {
	return []interface{}{&c.Color}
}

//CompileTo implements gpu.Shader.
func (Colored) CompileTo(platform, version string) (interface{}, error) {
	return glsl.Shader{
		Vertex: `#version 460 core

		in vec4 position;

		out int gpu_DrawID;

		uniform mat4 camera;

		layout(binding=0,std430) readonly buffer transformBuffer{
			mat4 transformArray[];
		};
		#define transform transformArray[gl_DrawID]

		void main() {
			gl_Position = camera*transform*position;
			gpu_DrawID = gl_DrawID;
		}
	`,
		Fragment: `#version 460 core
		out vec4 FragColor;

		flat in int gpu_DrawID;

		layout(binding=1,std430) readonly buffer bufColor{
			vec3 Color[];
		};
		#define color Color[gpu_DrawID]

		void main() {
			FragColor = vec4(color, 1);
		}
	`,
	}, nil
}

//Textured material applies a texture.
type Textured struct {
	Texture gpu.Texture
}

func (c *Textured) Variables() []interface{} {
	return []interface{}{&c.Texture}
}

//CompileTo implements gpu.Shader.
func (Textured) CompileTo(platform, version string) (interface{}, error) {
	return glsl.Shader{
		Vertex: `#version 460 core

		in vec4 position;
		in vec2 uv;

		out int gpu_DrawID;
		out vec2 frag_uv;

		uniform mat4 camera;

		layout(binding=0,std430) readonly buffer transformBuffer{
			mat4 transformArray[];
		};
		#define transform transformArray[gl_DrawID]

		void main() {
			  
		}
	`,
		Fragment: `#version 460 core

		#extension GL_ARB_bindless_texture : require

		out vec4 FragColor;

		flat in int gpu_DrawID;
		in vec2 frag_uv;

		layout(binding=1,std430) readonly buffer textureBuffer{
			sampler2D textureArray[];
		};
		#define albedo textureArray[gpu_DrawID]

		void main() {
			vec4 c = texture(albedo, frag_uv);
			if (c.a<0.85) discard;
			FragColor = c;
		}
	`,
	}, nil
}
