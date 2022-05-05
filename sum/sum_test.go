package sum_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"qlova.tech/sum"
)

type Triplet struct {
	One   sum.Add[Triplet, One]
	Two   sum.Add[Triplet, Two]
	Three sum.Add[Triplet, Three]
}

var Triplets = sum.Type[Triplet]{}.Sum()

type One struct{}

func (one One) String() string {
	return fmt.Sprint("ONE")
}

type Two struct{}

func (two Two) String() string {
	return fmt.Sprint("TWO")
}

type Three struct{}

func (three Three) String() string {
	return fmt.Sprint("THREE")
}

func TestTriplet(t *testing.T) {
	var triplet sum.Type[Triplet]

	triplet.Switch(Triplet{
		sum.Case(triplet, func(one One) {}),
		sum.Case(triplet, func(two Two) { t.Fatal() }),
		sum.Case(triplet, func(three Three) { t.Fatal() }),
	})

	triplet = Triplets.One.New(One{})

	triplet.Switch(Triplet{
		sum.Case(triplet, func(one One) {}),
		sum.Case(triplet, func(two Two) { t.Fatal() }),
		sum.Case(triplet, func(three Three) { t.Fatal() }),
	})

	triplet = Triplets.Two.New(Two{})

	triplet.Switch(Triplet{
		sum.Case(triplet, func(one One) { t.Fatal() }),
		sum.Case(triplet, func(two Two) {}),
		sum.Case(triplet, func(three Three) { t.Fatal() }),
	})

	triplet = Triplets.Three.New(Three{})

	triplet.Switch(Triplet{
		sum.Case(triplet, func(one One) { t.Fatal() }),
		sum.Case(triplet, func(two Two) { t.Fatal() }),
		sum.Case(triplet, func(three Three) {}),
	})

	b, err := json.Marshal(triplet)
	if err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(b, &triplet); err != nil {
		t.Fatal(err)
	}

	triplet.Switch(Triplet{
		sum.Case(triplet, func(one One) { t.Fatal() }),
		sum.Case(triplet, func(two Two) { t.Fatal() }),
		sum.Case(triplet, func(three Three) {}),
	})
}

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

func TestWeekday(t *testing.T) {
	var day sum.Int[Weekday]

	day = Weekdays.Friday
	if day.String() != "Friday" {
		t.Fatal()
	}

	if day != Weekdays.Friday {
		t.Fatal()
	}

	day.Switch(Weekdays, Weekday{
		day.Case(Weekdays.Monday, func() { t.Fatal() }),
		day.Case(Weekdays.Tuesday, func() { t.Fatal() }),
		day.Case(Weekdays.Wednesday, func() { t.Fatal() }),
		day.Case(Weekdays.Thursday, func() { t.Fatal() }),
		day.Case(Weekdays.Friday, func() {}),
		day.Case(Weekdays.Saturday, func() { t.Fatal() }),
		day.Case(Weekdays.Sunday, func() { t.Fatal() }),
	})

	if err := day.UnmarshalText([]byte("Saturday")); err != nil {
		t.Fatal(err)
	}
	if day.String() != "Saturday" {
		t.Fatal()
	}
}
