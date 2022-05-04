package conditional

type Type byte

const (
	Boolean Type = iota

	Number
	NumberSame
	NumberLess
	NumberMore

	String
	StringRegex

	Slice
	Map
)

type Expression struct {
	Type Type
	Args []any
}
