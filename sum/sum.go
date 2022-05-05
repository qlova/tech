/*
	Package sum provides sum types / enumerated types / tagged unions / variants / discriminated unions.

	Sum Types

		type SVGPathCommand struct {
			Line       sum.Add[SVGPathCommand, struct{X float64; Y float64}]
			Horizontal sum.Add[SVGPathCommand, float64]
			Vertical   sum.Add[SVGPathCommand, float64]
		}

		var SVGPathCommands = sum.Type[SVGPathCommand]{}.Sum()

		var command = SVGPathCommands.Horizontal.New(22)

		command.Switch(SVGPathCommands{
			sum.Switch(command, func(line struct{X float64; Y float64}) {
				// draw line
			}),
			sum.Switch(command, func(horizontal float64) {
				// move horizontally
			}),
			sum.Switch(command, func(vertical float64) {
				// move vertically
			}),
		})

	Enumerated Sum Types

		type Weekday struct {
			Monday,
			Tuesday,
			Wednesday,
			Thursday,
			Friday,
			Saturday,
			Sunday sum.Int[Weekday]
		}

		var Weekdays = sum.Int[Weekday]{}.Sum()

		var day sum.Int[Weekday]
		day = Weekdays.Monday
*/
package sum

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

/*
	Fields is a named struct that describes a sum type, each field
 	may Add a type to the sum, for example:

		type MySumType struct {
			Int    sum.Add[MySumType, Int]
			String sum.Add[MySumType, String]
		}

	The first sum.Add field will be the type of the zero value.
	Any field that is not a sum.Add, will be ignored by the sum
	package. Multiple fields may be added to the same type, they
	are distinguished by their name.
*/
type Fields any

// Type is a sum type / discriminated union / variant type
// the type parameter should refer to a named Fields struct
// with sum.Add fields that add to the sum of types.
type Type[Sum Fields] struct {
	any
	tag int // needed to distinguish fields of the same type and enable an O(1) Switch.
}

// String implements fmt.Stringer.
func (sum Type[Sum]) String() string {
	return fmt.Sprint(sum.any)
}

// Sum returns a set of fields that can be used to set the value of
// a sum type, you may wish to initialise this as a global variable
// to cache the result.
func (Type[Sum]) Sum() Sum {
	var sum Sum
	rvalue := reflect.ValueOf(&sum).Elem()
	for i := 0; i < rvalue.NumField(); i++ {
		rvalue.Field(i).Addr().Interface().(interface{ init(int) }).init(i)
	}
	return sum
}

// MarshalJSON implements json.Marshaler.
func (sum Type[Sum]) MarshalJSON() ([]byte, error) {
	var fields Sum
	var container struct { // TODO configurable container? struct tags?
		Type string      `json:"type"`
		Data interface{} `json:"data"`
	}
	container.Type = reflect.TypeOf(fields).Field(sum.tag).Name
	container.Data = sum.any
	return json.Marshal(container)
}

// MarshalJSON implements json.Unmarshaler.
func (sum *Type[Sum]) UmarshalJSON(data []byte) error {
	var container struct { // TODO configurable container? struct tags?
		Type string          `json:"type"`
		Data json.RawMessage `json:"data"`
	}
	if container.Data == nil {
		*sum = Type[Sum]{}
		return nil
	}
	if err := json.Unmarshal(data, &container); err != nil {
		return err
	}
	var fields Sum
	rtype := reflect.TypeOf(fields).Elem()
	for i := 0; i < rtype.NumField(); i++ {
		if rtype.Field(i).Name == container.Type {
			var zero = reflect.New(rtype.Field(i).Type).Elem()
			if err := json.Unmarshal(container.Data, zero.Interface()); err != nil {
				return err
			}
			sum.any = zero.Interface()
			sum.tag = i
			return nil
		}
	}
	return fmt.Errorf("unknown type %s", container.Type)
}

// Add is used to add fixed types to the Fields of a Type.
type Add[Sum Fields, T any] struct {
	tag int
	fun func(Type[Sum])
}

func (field *Add[Sum, T]) init(index int) { field.tag = index } // used by Type.Sum
func (field Add[Sum, T]) switchable(v Type[Sum]) {
	if field.fun != nil {
		field.fun(v)
	}
}

// New returns a value of Type[Sum] with a value matching the
// type this Add corresponds to.
func (field Add[Sum, T]) New(val T) Type[Sum] {
	return Type[Sum]{
		any: val,
		tag: field.tag,
	}
}

// Switch can be used to read the value of a sum type, the cases
// can should be added with Case calls. If the cases are named, then
// this is a non-exhaustive switch. If the cases are not named, then
// the switch is exhaustive. You can force a type switch to be non
// exhaustive by adding a _ struct{} field to the Fields of the sum
// type.
func (val Type[Sum]) Switch(cases Sum) {
	switcher, ok := reflect.ValueOf(cases).Field(val.tag).Interface().(interface {
		switchable(Type[Sum])
	})
	if ok {
		switcher.switchable(val)
	}
}

// Case can be passed to a switch to run a function when the associated
// type in the sum type is matched. The val passed to this function
// should refer to the value being switched on.
func Case[Sum any, T any](val Type[Sum], fn func(T)) Add[Sum, T] {
	return Add[Sum, T]{
		fun: func(val Type[Sum]) {
			v, _ := val.any.(T)
			fn(v)
		},
	}
}

/*
	ComparableFields is a named struct that describes an enumerated sum type,
	each field adds a distinct value to the enumerated set. The size and
	configuration of the underlying value is by convention defined on the
	last field in the structure.

		type Weekdays struct {
			Monday,
			Tuesday,
			Wednesday,
			Thursday,
			Friday,
			Saturday,
			Sunday sum.Int[Weekdays]
		}

	The first field is considered the zero value.
*/
type ComparableFields interface {
	comparable
}

// Int is a special kind of sum type, where each field is untyped. The
// underlying value is stored as an int.
type Int[Sum ComparableFields] struct {
	int
}

// String implements fmt.Stringer.
func (enum Int[Sum]) String() string {
	var sum Sum
	return reflect.TypeOf(sum).Field(enum.int).Name
}

// MarshalText implements encoding.TextMarshaler.
func (enum Int[Sum]) MarshalText() ([]byte, error) {
	return []byte(enum.String()), nil
}

// UmarshalText implements encoding.TextUnmarshaler.
func (enum *Int[Sum]) UnmarshalText(text []byte) error {
	var sum Sum

	rtype := reflect.TypeOf(sum)
	for i := 0; i < rtype.NumField(); i++ {
		field := rtype.Field(i)
		if field.Type != reflect.TypeOf(*enum) {
			continue
		}
		if string(text) == field.Name {
			enum.int = i
			return nil
		}
	}

	return fmt.Errorf("invalid %T value: %s", sum, text)
}

// MarshalJSON implements json.Marshaler.
func (enum Int[Sum]) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum.String())), nil
}

// MarshalJSON implements json.Unmarshaler.
func (enum Int[Sum]) UmarshalJSON(data []byte) error {
	s, err := strconv.Unquote(string(data))
	if err != nil {
		return err
	}
	return enum.UnmarshalText([]byte(s))
}

// Sum returns a set of fields that can be used to set the value of
// an enumerated sum type, you may wish to initialise this as a global
// variable to cache the result.
func (enum Int[Sum]) Sum() Sum {
	var sum Sum

	rvalue := reflect.ValueOf(&sum).Elem()
	for i := 0; i < rvalue.NumField(); i++ {
		field := rvalue.Field(i)

		if field.Type() != reflect.TypeOf(enum) {
			continue
		}

		enum.int = i
		field.Set(reflect.ValueOf(enum))
	}

	return sum
}

// Switch can be used to create an O(N) exaustive switch on this value
// by providing unnamed cases. To perform a non-exhaustive switch, use
// the builtin switch statement instead, which will likely be O(1).
func (enum Int[Sum]) Switch(fields Fields, cases Sum) {
	var zero Sum
	if fields == zero || fields != cases { //TODO disable this check with build tag?
		panic("invalid switch, fields and cases do not match")
	}
}

// Case evaulates that that the enumerated value is equal to the
// given case, if it is, the provided function is executed.
func (enum Int[Sum]) Case(val Int[Sum], fn func()) Int[Sum] {
	if enum == val {
		fn()
	}
	return val
}
