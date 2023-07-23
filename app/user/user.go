package user

import (
	"qlova.tech/app/show"
	"qlova.tech/app/then"
	"qlova.tech/app/user/hint"
)

type Data interface{}

type Interface = show.Node

func View(hints ...hint.View) func(ndoes ...show.Node) show.View {
	return func(nodes ...show.Node) show.View {
		node := show.View{
			Nodes: nodes,
		}
		for _, hint := range hints {
			hint.HintView(&node)
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

func List[T any](list *[]T, fn func(*T) Interface, args ...any) show.Node {
	return nil
}

type Showable interface {
	~*string
}

func Show[T Showable](viewable T, args ...any) show.Node {
	return nil
}
