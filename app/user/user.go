package user

import (
	"qlova.tech/app/data"
	"qlova.tech/app/show"
	"qlova.tech/app/then"
	"qlova.tech/app/user/hint"
)

type Data interface{}

type Interface = show.Node

func View(hints ...hint.Layout) func(ndoes ...show.Node) show.Layout {
	return func(nodes ...show.Node) show.Layout {
		node := show.Layout{
			Nodes: nodes,
		}
		for _, hint := range hints {
			hint.HintLayout(&node)
		}
		return node
	}
}

type literalString string

func Read(text literalString, args ...hint.String) show.String {
	node := show.String{
		Value: string(text),
	}
	for _, arg := range args {
		arg.HintString(&node)
	}
	return node
}

type Step = then.Step

type Steps = then.Steps

type Pathable interface {
	Steps | func() Steps
}

func Path[T Pathable](path T, args ...hint.Choice) show.Choice {
	var choice show.Choice
	switch path := any(path).(type) {
	case Steps:
		choice.Steps = []then.Step(path)
	case func() Steps:
		choice.Steps = path()
	default:
		panic("invalid path")
	}
	for _, arg := range args {
		arg.HintChoice(&choice)
	}
	return choice
}

type Pickable interface {
	string
}

func Pick[T Pickable](pickable *T, args ...any) show.Picker[T] {
	return show.Picker[T]{
		Value: pickable,
	}
}

func List[T any](list *[]T, fn func(*T) show.Layout, args ...any) show.Looped {

	return show.Looped{
		Value: list,
		Loops: func(s data.Sync) (data.Sync, show.Layout) {
			var zero T
			return s.With(&zero, list), fn(&zero)
		},
	}
}

type Showable interface {
	~string
}

func Show[T Showable](viewable *T, args ...any) show.Viewer[T] {
	return show.Viewer[T]{
		Value: viewable,
	}
}
