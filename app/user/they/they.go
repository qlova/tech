// Package they provides functions for manipulating the user.
package they

import (
	"qlova.tech/app/then"
)

func Append[T any](item *T, list *[]T) then.Append {
	return then.Append{
		Item: item,
		List: list,
	}
}

func Remove[T any](item *T, list *[]T) then.Steps {
	return nil
}

func Set[T any](variable *T, value T) then.Set {
	return then.Set{
		Variable: variable,
		Value:    value,
	}
}
