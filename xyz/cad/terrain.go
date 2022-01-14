package cad

type Terrain struct {
	Vertices, Normals [][3]float32
	UVs               [][2]float32

	Indicies []uint32
}

func NewTerrain(w, h float64, resolution int, heightmap func(x, y float64) float64) Terrain {
	segmentWidth := w / float64(resolution)
	segmentHeight := h / float64(resolution)

	const overscan = 4

	//Create the terrain slightly larger so that normals are correct on the edge.
	//width := w + segmentWidth*overscan*2
	//height := h + segmentHeight*overscan*2

	//widthHalf := width / 2
	//heightHalf := height / 2
	gridX := resolution + overscan*2
	gridY := resolution + overscan*2
	gridX1 := gridX + 1
	gridY1 := gridY + 1

	// Create buffers
	var terrain Terrain

	// Generate plane vertices, vertices normals and vertices texture mappings.
	for iy := 0; iy < gridY1; iy++ {
		y := float64(iy)*segmentHeight - overscan*segmentHeight
		for ix := 0; ix < gridX1; ix++ {
			x := float64(ix)*segmentWidth - overscan*segmentWidth

			terrain.Vertices = append(terrain.Vertices, [3]float32{float32(x), float32(heightmap(x, y)), float32(y)})
			terrain.Normals = append(terrain.Normals, [3]float32{0, 0, 0})
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

	to64 := func(in [3]float32) [3]float64 {
		return [3]float64{float64(in[0]), float64(in[1]), float64(in[2])}
	}

	to32 := func(in [3]float64) [3]float32 {
		return [3]float32{float32(in[0]), float32(in[1]), float32(in[2])}
	}

	//calculate the normals
	for i := 0; i < len(terrain.Indicies); i += 3 {
		var index = [3]uint32{
			terrain.Indicies[i],
			terrain.Indicies[i+1],
			terrain.Indicies[i+2],
		}

		var va, vb, vc [3]float64 = to64(terrain.Vertices[index[0]]),
			to64(terrain.Vertices[index[1]]),
			to64(terrain.Vertices[index[2]])

		e1 := subVec(vb, va)
		e2 := subVec(vc, va)
		no := cross(e1, e2)

		terrain.Normals[index[0]] = to32(addVec(to64(terrain.Normals[index[0]]), no))
		terrain.Normals[index[1]] = to32(addVec(to64(terrain.Normals[index[1]]), no))
		terrain.Normals[index[2]] = to32(addVec(to64(terrain.Normals[index[2]]), no))
	}

	terrain.Indicies = nil

	// Generate plane vertices indices for the faces
	for iy := overscan; iy < gridY-overscan; iy++ {
		for ix := overscan; ix < gridX-overscan; ix++ {
			a := ix + gridX1*iy
			b := ix + gridX1*(iy+1)
			c := (ix + 1) + gridX1*(iy+1)
			d := (ix + 1) + gridX1*iy

			terrain.Indicies = append(terrain.Indicies, uint32(a), uint32(b), uint32(d))
			terrain.Indicies = append(terrain.Indicies, uint32(b), uint32(c), uint32(d))
		}
	}

	for i := range terrain.Normals {
		terrain.Normals[i] = to32(normalize(to64(terrain.Normals[i])))
	}

	return terrain
}
