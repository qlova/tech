package sum_test

import (
	"encoding/json"
	"testing"

	"qlova.tech/sum"
)

type SVGPathCommand struct {
	Line sum.Add[SVGPathCommand, struct {
		X float64
		Y float64
	}]
	Horizontal sum.Add[SVGPathCommand, float64]
	Vertical   sum.Add[SVGPathCommand, float64]
}

var SVGPathCommands = sum.Type[SVGPathCommand]{}.Sum()

func TestSVG(t *testing.T) {
	var command sum.Type[SVGPathCommand]
	command = SVGPathCommands.Horizontal.New(20)

	command.Switch(SVGPathCommand{
		sum.Case(command, func(struct {
			X float64
			Y float64
		}) {
			t.Fail()
		}),
		sum.Case(command, func(v float64) {
			if v != 20 {
				t.Fail()
			}
		}),
		sum.Case(command, func(float64) {
			t.Fail()
		}),
	})

	command = SVGPathCommands.Vertical.New(20)

	command.Switch(SVGPathCommand{
		sum.Case(command, func(struct {
			X float64
			Y float64
		}) {
			t.Fail()
		}),
		sum.Case(command, func(float64) {
			t.Fail()
		}),
		sum.Case(command, func(v float64) {
			if v != 20 {
				t.Fail()
			}
		}),
	})

	b, err := json.Marshal(command)
	if err != nil {
		t.Fatal(err)
	}

	if err := json.Unmarshal(b, &command); err != nil {
		t.Fatal(err)
	}

	command.Switch(SVGPathCommand{
		sum.Case(command, func(struct {
			X float64
			Y float64
		}) {
			t.Fail()
		}),
		sum.Case(command, func(float64) {
			t.Fail()
		}),
		sum.Case(command, func(v float64) {
			if v != 20 {
				t.Fail()
			}
		}),
	})
}
