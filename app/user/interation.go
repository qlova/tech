/*
	Package user provides a pattern for describing interactive user interfaces.
*/
package user

import (
	"net/url"
	"time"

	"qlova.tech/app/internal/conditional"
	"qlova.tech/app/internal/instruction"
	"qlova.tech/app/internal/interaction"
)

/*
	Interface is a collection of possible interactions for a user.

	The available Interactions are:

	 	- Read - interpret a string.
		- View - interpret a value.
		- Pick - select a value.
		- Give - provide personal info.
		- When - proceed when the given condition is true.

	Each Interaction has a Target, which can be decorated with a
	semantic hint, in order to assist the user in organising the
	interface.
*/
type Interface = []interaction.Request

// InterfaceRenderer is an interactive value that can render
// a user interface that a user can interact with.
type InterfaceRenderer interface {
	RenderInterface() Interface
}

// List is an iterable slice or map with
// a custom interface for their elements.
type List struct {
	Value     any
	Interface Interface
}

// Readable values.
type Readable interface {
	~*string | ~string
}

// Read requests the user to read the given text.
func Read[Text Readable](text Text) interaction.Request {
	return interaction.Request{
		Type: interaction.Read,
		Args: []any{text},
	}
}

// Viewable values.
type Viewable interface {
	~[]interaction.Request | List |
		~[]byte | ~*[]byte |
		~[]string | ~*[]string |
		~map[string]string | ~*map[string]string
}

// View requests the user to view the given data.
func View[Data Viewable](data Data) interaction.Request {
	return interaction.Request{
		Type: interaction.View,
		Args: []any{data},
	}
}

// Pickable values are values that the user is able to
// select.
type Pickable interface {
	~bool | ~string |
		~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint |
		~int8 | ~int16 | ~int32 | ~int64 | ~int |
		~float32 | ~float64 | url.URL | time.Time
}

// Pick requests the user to select a value, if any options are provided
// the value must be chosen from the given options.
func Pick[Value Pickable](value *Value, options ...Value) interaction.Request {
	return interaction.Request{
		Type: interaction.Pick,
		Args: []any{value},
	}
}

// Path requests the user to decide whether to take this path, if they decide to
// take this path, they must follow the given instructions.
func Path[Name ~string](name Name, instructions ...instruction.Request) interaction.Request {
	return interaction.Request{
		Type: interaction.Path,
		Args: []any{name, instructions},
	}
}

func When(expr conditional.Expression, request interaction.Request) interaction.Request {
	return interaction.Request{
		Type: interaction.When,
		Args: []any{expr, request},
	}
}

type SliceFunc[T any] func(value *T) Interface

func Slice[T any](slice *[]T, fn SliceFunc[T]) List {
	var value = new(T)
	return List{
		Value:     slice,
		Interface: fn(value),
	}
}
