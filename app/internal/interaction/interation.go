package interaction

// Type refers to a kind of interaction that the user can make.
type Type byte

// All available interaction types.
const (
	Read Type = iota
	View
	Pick
	Path
	When
)

type Request struct {
	Type Type
	Args []any
}
