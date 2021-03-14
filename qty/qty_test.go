package qty_test

import (
	"fmt"
	"testing"

	"qlova.tech/qty"
)

func Test_Main(t *testing.T) {

	var person struct {
		Name   string
		Height qty.Length
		Weight qty.Mass

		Size qty.Data
	}

	person.Name = "John"
	person.Height = qty.Metres(2)
	person.Weight = qty.Kilograms(80)
	person.Size = qty.Bytes(100)

	fmt.Println(person.Height.In(qty.Inches), person.Weight.In(qty.Pounds))
	fmt.Println(person.Size)
}
