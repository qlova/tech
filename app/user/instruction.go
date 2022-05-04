package user

import (
	"context"

	"qlova.tech/app/internal/instruction"
)

// Postable values are values that can be posted.
type Postable interface {

	// Post processes the value and
	// produces some sort of side-effect.
	Post(ctx context.Context) error
}

func Post(data Postable) instruction.Request {
	return instruction.Request{
		Type: instruction.Post,
		Args: []any{data},
	}
}

// Edit requests a user to locally edit the given pointer with the given value.
func Edit[T any](pointer *T, value *T) instruction.Request {
	return instruction.Request{
		Type: instruction.Edit,
		Args: []any{pointer, value},
	}
}

// Append requests a user to locally append the given value to the given slice.
func Append[T any](slice *[]T, value *T) instruction.Request {
	return instruction.Request{
		Type: instruction.Append,
		Args: []any{slice, value},
	}
}

// Remove requests a user to locally remove the given value from the given slice.
func Remove[T any](slice *[]T, value *T) instruction.Request {
	return instruction.Request{
		Type: instruction.Remove,
		Args: []any{slice, value},
	}
}
