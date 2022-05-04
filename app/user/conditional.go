package user

import "qlova.tech/app/internal/conditional"

func StringHasRunes[String ~*string](s String) conditional.Expression {
	return conditional.Expression{
		Type: conditional.String,
		Args: []any{s},
	}
}

func SliceHasElements[T any](s *[]T) conditional.Expression {
	return conditional.Expression{
		Type: conditional.Slice,
		Args: []any{s},
	}
}
