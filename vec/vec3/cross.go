package vec3

//Cross returns the cross product of two vectors, x and y.
func Cross(x, y Type) Type {
	return Type{
		x[Y]*y[Z] - x[Z]*y[Y],
		x[Z]*y[X] - x[X]*y[Z],
		x[X]*y[Y] - x[Y]*y[X],
	}
}
