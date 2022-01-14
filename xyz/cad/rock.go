package cad

import (
	"math/rand"

	float "qlova.tech/xyz/f32"
	"qlova.tech/xyz/vec3"

	"github.com/ojrac/opensimplex-go"
)

type Rock struct {
	Seed int64

	NoiseScale,
	NoiseStrength,
	ScrapeMinDist,
	ScrapeStrength,
	ScrapeRadius float32

	ScrapeCount int

	Scale vec3.Float32

	Random func() float32

	adjacentVertices []map[uint32]struct{}
}

func (rock Rock) getNeighbours(positions [][3]float32, cells []uint32) []map[uint32]struct{} {
	/*
	   adjacentVertices[i] contains a set containing all the indices of the neighbours of the vertex with
	   index i.
	   A set is used because it makes the algorithm more convenient.
	*/
	var adjacentVertices = make([]map[uint32]struct{}, len(positions))

	// go through all faces.
	for iCell := 0; iCell < len(cells); iCell += 3 {

		var cellPositions = [3]uint32{cells[iCell], cells[iCell+1], cells[iCell+2]}

		wrap := func(i int) int {
			if i < 0 {
				return len(cellPositions) + i
			}
			return i % len(cellPositions)
		}

		// go through all the points of the face.
		for iPosition := 0; iPosition < len(cellPositions); iPosition++ {

			// the neighbours of this points are the previous and next points(in the array)
			var cur = cellPositions[wrap(iPosition+0)]
			var prev = cellPositions[wrap(iPosition-1)]
			var next = cellPositions[wrap(iPosition+1)]

			// create set on the fly if necessary.
			if adjacentVertices[cur] == nil {
				adjacentVertices[cur] = make(map[uint32]struct{})
			}

			// add adjacent vertices.
			adjacentVertices[cur][prev] = struct{}{}
			adjacentVertices[cur][next] = struct{}{}

		}
	}

	return adjacentVertices
}

func (rock Rock) scrape(positionIndex int, positions [][3]float32, cells []uint32, normals [][3]float32,
	adjacentVertices []map[uint32]struct{}, strength float32, radius float32) {

	var traversed = make([]bool, len(positions))
	var centerPosition = positions[positionIndex]

	// to scrape, we simply project all vertices that are close to `centerPosition`
	// onto a plane. The equation of this plane is given by dot(n, r-r0) = 0,
	// where n is the plane normal, r0 is a point on the plane(in our case we set this to be the projected center),
	// and r is some arbitrary point on the plane.
	var n = normals[positionIndex]

	var r0 = centerPosition
	vec3.Add(r0, vec3.Float32(n).Times(-strength))

	var stack []int
	stack = append(stack, positionIndex)

	/*
		Projects the point `p` onto the plane defined by the normal `n` and the point `r0`
	*/
	project := func(n, r0, p vec3.Float32) vec3.Float32 {
		// For an explanation of the math, see http://math.stackexchange.com/a/100766

		var t = vec3.Dot(n, vec3.Sub(r0, p)) / vec3.Dot(n, n)

		var projectedP = vec3.Add(p, n.Times(t))

		return projectedP
	}

	/*
	 We use a simple flood-fill algorithm to make sure that we scrape all vertices around the center.
	 This will be fast, since all vertices have knowledge about their neighbours.
	*/
	for len(stack) > 0 {

		var topIndex = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if traversed[topIndex] {
			continue // already traversed; look at next element in stack.
		}
		traversed[topIndex] = true

		var topPosition = positions[topIndex]
		// project onto plane.
		var p = topPosition

		var projectedP = project(n, r0, p)

		if vec3.Distance(projectedP, r0) < radius {
			positions[topIndex] = projectedP
			normals[topIndex] = n
		} else {
			continue
		}

		var neighbourIndices = adjacentVertices[topIndex]
		for i := range neighbourIndices {
			stack = append(stack, int(i))
		}

	}

}

func (rock Rock) Mesh() Mesh {
	if rock.Random == nil {
		rock.Random = rand.Float32
	}
	if rock.Scale == (vec3.Float32{}) {
		rock.Scale = vec3.Float32{1, 1, 1}
	}

	var noise = opensimplex.NewNormalized(rock.Seed)

	var sphere = NewSphere(vec3.Float32{}, 1).Mesh(20)

	var positions = sphere.Vertices
	var cells = sphere.Indicies
	var normals = sphere.Normals

	if rock.adjacentVertices == nil {
		rock.adjacentVertices = rock.getNeighbours(positions, cells)
	}

	/*
	   randomly generate positions at which to scrape.
	*/
	var scrapeIndices []int

	for i := 0; i < rock.ScrapeCount; i++ {

		var attempts = 0

		// find random position which is not too close to the other positions.
		for {

			var randIndex = int(float.Floor(float32(len(positions)) * rock.Random()))
			var p = positions[randIndex]

			var tooClose = false
			// check that it is not too close to the other vertices.
			for j := 0; j < len(scrapeIndices); j++ {

				var q = positions[scrapeIndices[j]]

				if vec3.Distance(p, q) < rock.ScrapeMinDist {
					tooClose = true
					break
				}
			}
			attempts++

			// if we have done too many attempts, we let it pass regardless.
			// otherwise, we risk an endless loop.
			if tooClose && attempts < 100 {
				continue
			} else {
				scrapeIndices = append(scrapeIndices, randIndex)
				break
			}
		}
	}

	// now we scrape at all the selected positions.
	for i := 0; i < len(scrapeIndices); i++ {
		rock.scrape(
			scrapeIndices[i], positions, cells, normals,
			rock.adjacentVertices, rock.ScrapeStrength, rock.ScrapeRadius)
	}

	/*
	   Finally, we apply a Perlin noise to slighty distort the mesh,
	    and then we scale the mesh.
	*/
	for i := 0; i < len(positions); i++ {
		var p = positions[i]

		var noise = rock.NoiseStrength * float32(noise.Eval3(
			float64(rock.NoiseScale*p[0]),
			float64(rock.NoiseScale*p[1]),
			float64(rock.NoiseScale*p[2])))

		positions[i][0] += noise
		positions[i][1] += noise
		positions[i][2] += noise

		positions[i][0] *= rock.Scale[0]
		positions[i][1] *= rock.Scale[1]
		positions[i][2] *= rock.Scale[2]
	}

	// of course, we must recompute the normals.
	var mesh Mesh

	mesh.Vertices = positions
	mesh.Indicies = cells
	mesh.CalculateNormals()

	return mesh
}
