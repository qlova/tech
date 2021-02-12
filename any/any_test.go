package any

import (
	"fmt"
	"testing"
)

type Test struct {
	A, B int32
}

type BigTest [100]int32

func TestSize(t *testing.T) {
	var v Size5

	v.Set(&Test{1, 3})

	fmt.Println(v.Value)

	v.Set(&BigTest{1, 3})

	fmt.Println(v.Value)
}
