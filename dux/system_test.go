package dux_test

import (
	"fmt"
	"testing"

	"qlova.tech/dux"
)

type State struct {
	dux.Component

	Thing int
}

type InputManager struct {
	dux.System

	*Renderer
	*State
}

func (system *InputManager) Update() {
	system.Thing++
}

type Renderer struct {
	dux.System

	*InputManager
	*State
}

func (system *Renderer) Update() {
	fmt.Println(system.Thing)
}

func TestX(t *testing.T) {
	var world = dux.New(
		new(InputManager),
		new(Renderer),
	)

	world.Update()
	world.Update()

	t.Fail()
}
