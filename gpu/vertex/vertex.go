//Package vertex provides common vertex attributes.
package vertex

//Attribute is a named attribute of a vertex.
type Attribute string

// Common attributes, use these in your shaders
// for maximum compatibility.
const (
	Position Attribute = "position"
	Normal   Attribute = "normal"
	UV       Attribute = "uv"
	Color    Attribute = "color"
	Weight   Attribute = "weight"
	Joint    Attribute = "joint"
)
