package gpu

import (
	"qlova.tech/rgb"
	"qlova.tech/xyz"
)

// Light is a compacted representation of a
// light, depending on its value, it may be
// a directional light, a point light, or a
// spotlight.
type Light struct {
	Position    xyz.Vector
	Direction   xyz.Vector
	Attenuation xyz.Vector
	Color       rgb.Color
}

// NewDirectionalLight returns a directional light.
func NewDirectionalLight(position, direction xyz.Vector, color rgb.Color) Light {
	return Light{
		Position:  position,
		Direction: direction,
		Color:     color,
	}
}

// NewPointLight returns a point light. Attenuation values must be non-zero.
func NewPointLight(position, attenuation xyz.Vector, color rgb.Color) Light {
	return Light{
		Position:    position,
		Attenuation: attenuation,
		Color:       color,
	}
}

// NewSpotLight returns a spotlight.
func NewSpotlight(position, direction xyz.Vector, angle float32, color rgb.Color) Light {
	return Light{
		Position:    position,
		Direction:   direction,
		Attenuation: xyz.Vector{angle, 0, 0},
		Color:       color,
	}
}
