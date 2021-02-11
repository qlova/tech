package cad

type Mesh struct {
	Vertices [][3]float32
	Normals  [][3]float32
	Indicies []uint32
}

func (mesh *Mesh) CalculateNormals() {
	to64 := func(in [3]float32) [3]float64 {
		return [3]float64{float64(in[0]), float64(in[1]), float64(in[2])}
	}

	to32 := func(in [3]float64) [3]float32 {
		return [3]float32{float32(in[0]), float32(in[1]), float32(in[2])}
	}

	mesh.Normals = make([][3]float32, len(mesh.Vertices))

	for i := 0; i < len(mesh.Indicies); i += 3 {
		var index = [3]uint32{
			mesh.Indicies[i],
			mesh.Indicies[i+1],
			mesh.Indicies[i+2],
		}

		var va, vb, vc [3]float64 = to64(mesh.Vertices[index[0]]),
			to64(mesh.Vertices[index[1]]),
			to64(mesh.Vertices[index[2]])

		e1 := subVec(vb, va)
		e2 := subVec(vc, va)
		no := cross(e1, e2)

		mesh.Normals[index[0]] = to32(addVec(to64(mesh.Normals[index[0]]), no))
		mesh.Normals[index[1]] = to32(addVec(to64(mesh.Normals[index[1]]), no))
		mesh.Normals[index[2]] = to32(addVec(to64(mesh.Normals[index[2]]), no))
	}
}
