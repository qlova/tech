package gpu

import (
	"math"

	"qlova.tech/ray"
	"qlova.tech/vec/vec3"
)

type node struct {
	children int

	mesh Mesh
}

//Model is a flattened 3D scene.
type Model []node

//NewModel creates a new model out of the given children meshes.
func NewModel(meshes ...Mesh) Model {
	var model Model
	for _, mesh := range meshes {
		model = append(model, node{
			mesh: mesh,
		})
	}
	return model
}

//Draw draws the model.
func (model Model) Draw(options DrawOptions, t *Transform) error {
	for _, node := range model {
		if err := node.mesh.Draw(options, t); err != nil {
			return err
		}
	}
	return nil
}

//Raycast checks if the specified ray intersects this model.
//If ok, the intersection point is returned.
func (model Model) Raycast(ray ray.Caster, t Transform) (point vec3.Type, ok bool) {

	var closest float32 = math.MaxFloat32

	for _, node := range model {
		if p, hit := node.mesh.Raycast(ray, t); hit {
			if dist := vec3.Distance(ray.Origin, p); dist < closest {
				closest = dist
				point = p
				ok = true
			}
		}
	}

	return
}
