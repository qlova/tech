//Package ray provides raycasting primitives.
package ray

import (
	"qlova.tech/cad"
	"qlova.tech/mat/mat4"
	"qlova.tech/vec/vec3"
)

//Caster is used to raycast geometry.
type Caster struct {
	Origin, Direction vec3.Type
}

func (ray Caster) at(t float32) vec3.Type {
	return vec3.Add(ray.Direction.Times(t), ray.Origin)
}

//Transform this ray by the given matrix.
func (ray *Caster) Transform(m *mat4.Type) {
	ray.Direction.Add(ray.Origin)
	ray.Direction.Transform(m)
	ray.Origin.Transform(m)
	ray.Direction.Sub(ray.Origin)
	ray.Direction.Normalize()
}

//Box tests for an intersection between the ray and the box.
//If ok, the returned point is the point of intersection.
func (ray Caster) Box(box cad.Box) (point vec3.Type, ok bool) {
	// http://www.scratchapixel.com/lessons/3d-basic-lessons/lesson-7-intersecting-simple-shapes/ray-box-intersection/

	var tmin, tmax, tymin, tymax, tzmin, tzmax float32

	invdirx := 1 / ray.Direction.X()
	invdiry := 1 / ray.Direction.Y()
	invdirz := 1 / ray.Direction.Z()

	var origin = ray.Origin

	if invdirx >= 0 {
		tmin = (box.Min().X() - origin.X()) * invdirx
		tmax = (box.Max().X() - origin.X()) * invdirx
	} else {
		tmin = (box.Max().X() - origin.X()) * invdirx
		tmax = (box.Min().X() - origin.X()) * invdirx
	}

	if invdiry >= 0 {
		tymin = (box.Min().Y() - origin.Y()) * invdiry
		tymax = (box.Max().Y() - origin.Y()) * invdiry
	} else {
		tymin = (box.Max().Y() - origin.Y()) * invdiry
		tymax = (box.Min().Y() - origin.Y()) * invdiry
	}

	if (tmin > tymax) || (tymin > tmax) {
		return
	}

	// These lines also handle the case where tmin or tmax is NaN
	// (result of 0 * Infinity). x !== x returns true if x is NaN

	if tymin > tmin || tmin != tmin {
		tmin = tymin
	}

	if tymax < tmax || tmax != tmax {
		tmax = tymax
	}

	if invdirz >= 0 {
		tzmin = (box.Min().Z() - origin.Z()) * invdirz
		tzmax = (box.Max().Z() - origin.Z()) * invdirz
	} else {
		tzmin = (box.Max().Z() - origin.Z()) * invdirz
		tzmax = (box.Min().Z() - origin.Z()) * invdirz
	}

	if (tmin > tzmax) || (tzmin > tmax) {
		return
	}

	if tzmin > tmin || tmin != tmin {
		tmin = tzmin
	}

	if tzmax < tmax || tmax != tmax {
		tmax = tzmax
	}

	//return point closest to the ray (positive side)

	if tmax < 0 {
		return
	}

	if tmin >= 0 {
		return ray.at(tmin), true
	}
	return ray.at(tmax), true
}

//Triangle tests for an intersection between the ray and the triangle.
//If ok, the returned point is the point of intersection.
func (ray Caster) Triangle(a, b, c vec3.Type) (point vec3.Type, ok bool) {
	var edge1 vec3.Type = vec3.Sub(b, a)
	var edge2 vec3.Type = vec3.Sub(c, a)
	var normal vec3.Type = vec3.Cross(edge1, edge2)

	// Solve Q + t*D = b1*E1 + b2*E2 (Q = kDiff, D = ray direction,
	// E1 = kEdge1, E2 = kEdge2, N = Cross(E1,E2)) by
	//   |Dot(D,N)|*b1 = sign(Dot(D,N))*Dot(D,Cross(Q,E2))
	//   |Dot(D,N)|*b2 = sign(Dot(D,N))*Dot(D,Cross(E1,Q))
	//   |Dot(D,N)|*t = -sign(Dot(D,N))*Dot(Q,N)
	DdN := vec3.Dot(ray.Direction, normal)
	var sign float32

	if DdN > 0 {
		sign = 1
	} else if DdN < 0 {
		sign = -1
		DdN = -DdN
	} else {
		return
	}

	var diff = vec3.Sub(ray.Origin, a)

	DdQxE2 := sign * vec3.Dot(ray.Direction, vec3.Cross(diff, edge2))

	// b1 < 0, no intersection
	if DdQxE2 < 0 {
		return
	}

	DdE1xQ := sign * vec3.Dot(ray.Direction, vec3.Cross(edge1, diff))
	// b2 < 0, no intersection
	if DdE1xQ < 0 {
		return
	}

	// b1+b2 > 1, no intersection
	if DdQxE2+DdE1xQ > DdN {
		return
	}

	// Line intersects triangle, check if ray does.
	QdN := -sign * vec3.Dot(diff, normal)

	// t < 0, no intersection
	if QdN < 0 {
		return
	}

	// Ray intersects triangle.
	return ray.at(QdN / DdN), true
}
