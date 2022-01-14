package cad

import (
	"math"

	float "qlova.tech/xyz/f32"
)

type Sphere struct {
	pos    Point
	radius float32
}

func NewSphere(pos Point, radius float32) Sphere {
	return Sphere{
		pos:    pos,
		radius: radius,
	}
}

func (sphere Sphere) Mesh(precision int) Mesh {
	var radius = sphere.radius
	var stacks = float32(precision)
	var slices = float32(precision)

	var positions [][3]float32
	var cells []uint32
	var normals [][3]float32

	// keeps track of the index of the next vertex that we create.
	var index = 0

	/*
	   First of all, we create all the faces that are NOT adjacent to the
	   bottom(0,-R,0) and top(0,+R,0) vertices of the sphere.
	   (it's easier this way, because for the bottom and top vertices, we need to add triangle faces.
	   But for the faces between, we need to add quad faces. )
	*/

	// loop through the stacks.
	for i := 1; i < int(stacks); i++ {

		var u = float32(i) / stacks
		var phi = u * math.Pi

		var stackBaseIndex = (len(cells) / 3) / 2

		// loop through the slices.
		for j := 0; j < int(slices); j++ {

			var v = float32(j) / slices
			var theta = v * (math.Pi * 2)

			var R = radius
			// use spherical coordinates to calculate the positions.
			var x = float.Cos(theta) * float.Sin(phi)
			var y = float.Cos(phi)
			var z = float.Sin(theta) * float.Sin(phi)

			positions = append(positions, [3]float32{R * x, R * y, R * z})
			normals = append(normals, [3]float32{x, y, z})

			if (i + 1) != int(stacks) { // for the last stack, we don't need to add faces.

				var i1, i2, i3, i4 uint32

				if (j + 1) == int(slices) {
					// for the last vertex in the slice, we need to wrap around to create the face.
					i1 = uint32(index)
					i2 = uint32(stackBaseIndex)
					i3 = uint32(index + int(slices))
					i4 = uint32(stackBaseIndex + int(slices))

				} else {
					// use the indices from the current slice, and indices from the next slice, to create the face.
					i1 = uint32(index)
					i2 = uint32(index + 1)
					i3 = uint32(index + int(slices))
					i4 = uint32(index + int(slices) + 1)
				}

				// add quad face
				cells = append(cells, i1, i2, i3)
				cells = append(cells, i4, i3, i2)
			}

			index++
		}
	}

	/*
	   Next, we finish the sphere by adding the faces that are adjacent to the top and bottom vertices.
	*/

	var topIndex = index
	index++
	positions = append(positions, [3]float32{0.0, radius, 0.0})
	normals = append(normals, [3]float32{0, 1, 0})

	var bottomIndex = index
	index++
	positions = append(positions, [3]float32{0, -radius, 0})
	normals = append(normals, [3]float32{0, -1, 0})

	for i := 0; i < int(slices); i++ {

		var i1 = uint32(topIndex)
		var i2 = uint32(i + 0)
		var i3 = uint32(i+1) % uint32(slices)
		cells = append(cells, i3, i2, i1)

		i1 = uint32(bottomIndex)
		i2 = uint32(bottomIndex-1) - uint32(slices) + uint32(i+0)
		i3 = uint32(bottomIndex-1) - uint32(slices) + uint32((i+1))%uint32(slices)
		cells = append(cells, i1, i2, i3)

	}

	return Mesh{Vertices: positions, Indicies: cells, Normals: normals}
}
