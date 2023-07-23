package then

type Step interface {
	step()
}

type Steps []Step

func (Steps) step() {}

type isStep = Step

type Append struct {
	isStep

	Item any
	List any
}

type Set struct {
	isStep

	Variable any
	Value    any
}
