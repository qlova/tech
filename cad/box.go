package cad

import "qlova.tech/vec/vec3"

//Box is a 3D box defined by a min and max point.
type Box [2]Point

//Min returns the minumum point of the box.
func (box Box) Min() Point { return box[0] }

//Max returns the maximum point of the box.
func (box Box) Max() Point { return box[1] }

// ExpandByPoint may expand this bounding box to include the specified point.
// Returns pointer to this updated bounding box.
func (box *Box) ExpandByPoint(point Point) {
	box[0] = vec3.Min(box[0], point)
	box[1] = vec3.Max(box[1], point)
}
