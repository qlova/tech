//Package catfacts provides random facts about cats, dogs, horses & snails.
package catfacts

import (
	"time"

	"qlova.tech/api"
	"qlova.tech/api/rest"
)

//Fact about an animal.
type Fact struct {
	Type Animal `json:"type"`
	ID   string `json:"_id"`
	Text string `json:"text"`

	User string `json:"user"`

	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`

	Status struct {
		Verified  bool `json:"verified"`
		SentCount int  `json:"sentCount"`
	} `json:"status"`

	Deleted bool `json:"deleted"`

	V int `json:"__v"`

	Source string `json:"source"`
	Used   bool   `json:"used"`
}

//Animal type for random animal facts.
type Animal string

//Animal types.
const (
	Cat   Animal = "cat"
	Dog   Animal = "dog"
	Snail Animal = "snail"
	Horse Animal = "horse"
)

//API that returns cat facts.
var API struct {
	rest.Interface `api:"https://cat-fact.herokuapp.com"`

	//Random is shorthand for RandomAnimal(Cat).
	Random func() (Fact, error) `api:"/facts/random"`

	//RandomAnimal returns a random fact about the animal you provided.
	RandomAnimal func(animal Animal) (Fact, error) `api:"/facts/random?animal_type=%v"`

	//RandomSlice is shorthand for RandomAnimalSlice(Cat, amount).
	RandomSlice func(amount int) ([]Fact, error) `api:"/facts/random?amount=%v"`

	//RandomCat returns a random fact about cats.
	RandomAnimalSlice func(animal Animal) (Fact, error) `api:"/facts/random?animal_type=%v&amount=%v"`

	//Get returns the fact with the given ID.
	Get func(id string) (Fact, error) `api:"/facts/%v"`
}

func init() {
	api.Connect(&API)
}
