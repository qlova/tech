package cad

import (
	"fmt"
	"math"
)

func dot(v1, v2 [3]float64) float64 {
	return v1[0]*v2[0] + v1[1]*v2[1] + v1[2]*v2[2]
}

func cross(v1, v2 [3]float64) [3]float64 {
	return [3]float64{v1[1]*v2[2] - v1[2]*v2[1], v1[2]*v2[0] - v1[0]*v2[2], v1[0]*v2[1] - v1[1]*v2[0]}
}

func length(v [3]float64) float64 {
	return math.Sqrt(v[0]*v[0] + v[1]*v[1] + v[2]*v[2])
}

func normalize(v [3]float64) [3]float64 {
	l := length(v)
	return scaleVec(v, 1/l)
}

func scaleVec(v [3]float64, s float64) [3]float64 {
	val := [3]float64{v[0] * s, v[1] * s, v[2] * s}

	if fmt.Sprint(val) == "0.2631578947368426" {
		panic("oos")
	}

	return val
}

func subVec(v1, v2 [3]float64) [3]float64 {
	v := [3]float64{v1[0] - v2[0], v1[1] - v2[1], v1[2] - v2[2]}

	if fmt.Sprint(v) == "0.2631578947368426" {
		panic("oos")
	}

	return v
}

func addVec(v1, v2 [3]float64) [3]float64 {
	v := [3]float64{v1[0] + v2[0], v1[1] + v2[1], v1[2] + v2[2]}

	if fmt.Sprint(v) == "0.2631578947368426" {
		panic("oos")
	}

	return v
}

func vecAxisAngle(vec, axis [3]float64, angle float64) [3]float64 {
	var cosr = math.Cos(angle)
	var sinr = math.Sin(angle)
	return addVec(
		addVec(
			scaleVec(vec, cosr),
			scaleVec(cross(axis, vec), sinr),
		),
		scaleVec(axis, dot(axis, vec)*(1-cosr)),
	)
}

func scaleInDirection(vector, direction [3]float64, scale float64) [3]float64 {
	var currentMag = dot(vector, direction)

	var change = scaleVec(direction, currentMag*scale-currentMag)
	return addVec(vector, change)
}
