package gpu

import "qlova.tech/mat/mat4"

var transforms [][]Transform

//Transform is a 4x4 matrix that represents the transform to apply to a mesh.
type Transform = mat4.Type

//NewTransform returns a new indentity transform that does not transform the target in any way.
func NewTransform() Transform {
	return Transform{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}
}

func pushTransform(t *Transform, buffer int) {
	if t != nil {
		if len(transforms) <= buffer {
			transforms = append(transforms, make([][]Transform, buffer-len(transforms)+1)...)
		}

		transforms[buffer] = append(transforms[buffer], *t)
	} else {
		panic("uh oh null transform")
	}
}
