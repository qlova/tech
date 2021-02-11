package gpu

import (
	"fmt"
	"testing"
)

type Test struct {
	A, B int32
}

func (t Test) CompileTo(string, string) (interface{}, error) { return nil, nil }

func (t Test) Variables() []interface{} { return nil }

type BigTest [100]int32

func (t BigTest) CompileTo(string, string) (interface{}, error) { return nil, nil }
func (t BigTest) Variables() []interface{}                      { return nil }

func TestSize(t *testing.T) {
	var v Material

	v.SetShader(&Test{1, 3})

	fmt.Println(v.shader)

	v.SetShader(&BigTest{1, 3})

	fmt.Println(v.shader)
}
