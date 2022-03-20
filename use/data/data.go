package data

func Type[T any]() *T {
	var data T
	return &data
}

type Patch []Mutation

type Mutation struct{}

func Set[T any](field *T, value T) Mutation {
	return Mutation{}
}

func Replace[T any](field *T, value *T) Mutation {
	return Mutation{}
}
