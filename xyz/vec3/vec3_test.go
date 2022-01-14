package vec3

import (
	"fmt"
	"testing"

	"qlova.org/should"
)

func TestVec3(t *testing.T) {
	a := Float32{1, 2, 3}
	b := Float32{4, -5, 6}

	should.Be("12")(fmt.Sprint(Dot(a, b))).Test(t)

	a = Float32{3, -3, 1}
	b = Float32{4, 9, 2}

	should.Be(Float32{-15, -2, 39})(Cross(a, b)).Test(t)
}
