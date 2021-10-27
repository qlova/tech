package glsl

import "errors"

//Shader is a glsl shader.
type Shader struct {
	Vertex   string
	Fragment string
}

//CompileTo implements shader.Shader
func (s Shader) CompileTo(platform, version string) (interface{}, error) {
	if platform != "glsl" {
		return nil, errors.New("unsupported platform")
	}

	return s, nil
}
