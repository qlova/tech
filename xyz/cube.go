package xyz

import (
	"reflect"

	"qlova.tech/xy"
)

func Cube() Vertices {
	return newSegmentedBox(1, 1, 1, 1, 1, 1)
}

type model struct {
	positions []Vector
	normals   []Vector
	uvs       []xy.Vector
	indices   []uint16
}

func (m model) Length() int {
	return len(m.indices)
}

func (m model) Buffers() []Buffer {
	return []Buffer{
		NewBuffer(m.positions),
		NewBuffer(m.normals),
		NewBuffer(m.uvs),
		NewBuffer(m.indices),
	}
}

func (m model) Indexed() (int, bool) {
	return 3, true
}

func (m model) Layout() []Pointer {
	return []Pointer{
		{
			Attribute: "position",
			Kind:      reflect.Float32,
			Buffer:    0,
			Count:     3,
		},
		{
			Attribute: "normal",
			Kind:      reflect.Float32,
			Buffer:    1,
			Count:     3,
		},
		{
			Attribute: "uv",
			Kind:      reflect.Float32,
			Buffer:    2,
			Count:     2,
		},
	}
}

//adapted from https://github.com/g3n/engine/blob/master/geometry/box.go#NewSegmentedBox
func newSegmentedBox(width, height, length float32, widthSegments, heightSegments, lengthSegments int) model {
	const x, y, z = 0, 1, 2

	var box model

	// Validate arguments
	if widthSegments <= 0 || heightSegments <= 0 || lengthSegments <= 0 {
		return model{}
	}

	// Internal function to build each of the six box planes
	buildPlane := func(u, v int, udir, vdir int, width, height, length float32, materialIndex uint) {

		offset := len(box.positions)
		gridX := widthSegments
		gridY := heightSegments
		var w int

		if (u == x && v == y) || (u == y && v == x) {
			w = z
		} else if (u == x && v == z) || (u == z && v == x) {
			w = y
			gridY = lengthSegments
		} else if (u == z && v == y) || (u == y && v == z) {
			w = x
			gridX = lengthSegments
		}

		var normal Vector
		if length > 0 {
			normal[w] = 1
		} else {
			normal[w] = -1
		}

		wHalf := width / 2
		hHalf := height / 2
		gridX1 := gridX + 1
		gridY1 := gridY + 1
		segmentWidth := width / float32(gridX)
		segmentHeight := height / float32(gridY)

		// Generate the plane vertices, normals, and uv coordinates
		for iy := 0; iy < gridY1; iy++ {
			for ix := 0; ix < gridX1; ix++ {
				var vector Vector
				vector[u] = (float32(ix)*segmentWidth - wHalf) * float32(udir)
				vector[v] = (float32(iy)*segmentHeight - hHalf) * float32(vdir)
				vector[w] = length
				box.positions = append(box.positions, vector)
				box.normals = append(box.normals, normal)
				box.uvs = append(box.uvs, xy.Vector{
					float32(ix) / float32(gridX),
					float32(1) - (float32(iy) / float32(gridY)),
				})
			}
		}

		// Generate the indices for the vertices, normals and uv coordinates
		for iy := 0; iy < gridY; iy++ {
			for ix := 0; ix < gridX; ix++ {
				a := ix + gridX1*iy
				b := ix + gridX1*(iy+1)
				c := (ix + 1) + gridX1*(iy+1)
				d := (ix + 1) + gridX1*iy
				box.indices = append(box.indices,
					uint16(a+offset), uint16(b+offset), uint16(d+offset),
					uint16(b+offset), uint16(c+offset), uint16(d+offset),
				)
			}
		}
	}

	wHalf := width / 2
	hHalf := height / 2
	lHalf := length / 2

	buildPlane(z, y, -1, -1, length, height, wHalf, 0) // px
	buildPlane(z, y, 1, -1, length, height, -wHalf, 1) // nx
	buildPlane(x, z, 1, 1, width, length, hHalf, 2)    // py
	buildPlane(x, z, 1, -1, width, length, -hHalf, 3)  // ny
	buildPlane(x, y, 1, -1, width, height, lHalf, 4)   // pz
	buildPlane(x, y, -1, -1, width, height, -lHalf, 5) // nz

	return box
}
