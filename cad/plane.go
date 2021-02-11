package cad

type Plane struct {
	Vertices, Normals [][3]float32
	UVs               [][2]float32

	Indicies []uint32
}

func NewPlane(w, h float64, resolution int) Plane {
	segmentWidth := w / float64(resolution)
	segmentHeight := h / float64(resolution)

	gridX := resolution
	gridY := resolution
	gridX1 := gridX + 1
	gridY1 := gridY + 1

	// Create buffers
	var terrain Plane

	// Generate plane vertices, vertices normals and vertices texture mappings.
	for iy := 0; iy < gridY1; iy++ {
		y := float64(iy) * segmentHeight
		for ix := 0; ix < gridX1; ix++ {
			x := float64(ix) * segmentWidth

			terrain.Vertices = append(terrain.Vertices, [3]float32{float32(x), 0, float32(y)})
			terrain.Normals = append(terrain.Normals, [3]float32{0, 1, 0})
			terrain.UVs = append(terrain.UVs, [2]float32{float32(float64(ix) / float64(gridX)), float32((float64(iy) / float64(gridY)))})
		}
	}

	// Generate plane vertices indices for the faces
	for iy := 0; iy < gridY; iy++ {
		for ix := 0; ix < gridX; ix++ {
			a := ix + gridX1*iy
			b := ix + gridX1*(iy+1)
			c := (ix + 1) + gridX1*(iy+1)
			d := (ix + 1) + gridX1*iy

			terrain.Indicies = append(terrain.Indicies, uint32(a), uint32(b), uint32(d))
			terrain.Indicies = append(terrain.Indicies, uint32(b), uint32(c), uint32(d))
		}
	}

	return terrain
}
