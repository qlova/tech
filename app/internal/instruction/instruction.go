package instruction

// Type refers to an instruction that the user can follow.
type Type int

// All available instruction types.
const (
	Edit Type = iota

	Append
	Remove

	Insert
	Delete

	Open
	Load
	Save
	Post

	Search

	Range
	Branch

	Add
	Sub
	Div
	Mul
	Mod
	Pow

	Predict
)

type Request struct {
	Type Type
	Args []any
}
