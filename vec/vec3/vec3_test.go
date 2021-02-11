package vec3

import (
	"fmt"
	"testing"

	"qlova.org/should"
)

func TestVec3(t *testing.T) {
	a := Type{1, 2, 3}
	b := Type{4, -5, 6}

	should.Be("12")(fmt.Sprint(Dot(a, b))).Test(t)

	a = Type{3, -3, 1}
	b = Type{4, 9, 2}

	should.Be(Type{-15, -2, 39})(Cross(a, b)).Test(t)
}
