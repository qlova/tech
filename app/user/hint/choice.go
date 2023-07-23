package hint

import "qlova.tech/app/show"

type Choice interface {
	HintChoice(*show.Choice)
}

type Button string

func (b Button) HintChoice(choice *show.Choice) {
	choice.Hints.Button = string(b)
}

var OnClick onClick

type onClick struct{}

func (onClick) HintChoice(choice *show.Choice) {
	choice.Hints.OnClick = true
}

var OnPress onPress

type onPress struct{}

func (onPress) HintChoice(choice *show.Choice) {
	choice.Hints.OnClick = true
	choice.Hints.OnTouch = true
}
