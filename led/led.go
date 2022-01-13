//Package led provides a lighting model.
package led

import (
	"qlova.tech/rgb"
	"qlova.tech/vec/vec3"
)

// Light is a compacted representation of a
// light, depending on its value, it may be
// a directional light, a point light, or a
// spotlight.
type Light struct {
	Position    vec3.Float32
	Direction   vec3.Float32
	Attenuation vec3.Float32
	Color       rgb.Color
}

// NewDirectionalLight returns a directional light.
func NewDirectionalLight(position, direction vec3.Float32, color rgb.Color) Light {
	return Light{
		Position:  position,
		Direction: direction,
		Color:     color,
	}
}

// NewPointLight returns a point light. Attenuation values must be non-zero.
func NewPointLight(position vec3.Float32, attenuation vec3.Float32, color rgb.Color) Light {
	return Light{
		Position:    position,
		Attenuation: attenuation,
		Color:       color,
	}
}

// NewSpotLight returns a spotlight.
func NewSpotlight(position, direction vec3.Float32, angle float32, color rgb.Color) Light {
	return Light{
		Position:    position,
		Direction:   direction,
		Attenuation: vec3.Float32{angle, 0, 0},
		Color:       color,
	}
}
