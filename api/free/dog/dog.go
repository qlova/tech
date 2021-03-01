//Package dog provides dog breeds and URLs to images of dogs.
package dog

import (
	"encoding/json"
	"errors"
	"io"
	"strings"

	"qlova.tech/api"
	"qlova.tech/api/rest"
)

//Protocol implements the Dog api.Protocol.
type Protocol struct {
	rest.Protocol
}

//DecodeValue implements api.Protocol.DecodeValue.
func (message Protocol) DecodeValue(reader io.Reader, value interface{}) error {
	var Message struct {
		Message json.RawMessage `json:"message"`
		Status  string          `json:"status"`
	}

	if err := message.Protocol.DecodeValue(reader, &Message); err != nil {
		return err
	}

	if Message.Status != "success" {
		return errors.New(Message.Status)
	}

	if err := json.Unmarshal(Message.Message, value); err != nil {
		return err
	}

	return nil
}

//Breeds is a list of dog breeds and their sub-breeds.
type Breeds map[Breed][]SubBreed

func (b Breeds) String() string {
	var s strings.Builder

	for breed, subBreeds := range b {
		s.WriteString(string(breed))
		s.WriteByte('\n')

		for _, subBreed := range subBreeds {
			s.WriteString(string(subBreed))
			s.WriteByte('\n')
		}
	}

	return s.String()
}

//Breed is a breed of dog.
type Breed string

//SubBreed is a subbreed of dog.
type SubBreed string

//WithSubBreed returns the breed with the given sub breed.
func (breed Breed) WithSubBreed(subbreed SubBreed) Breed {
	return breed + "/" + Breed(subbreed)
}

//API implements api.Interface
var API struct {
	rest.Interface `api:"https://dog.ceo/api"`

	Protocol Protocol

	//Breeds returns a map of dog breeds to sub-breeds.
	Breeds func() (Breeds, error) `api:"/breeds/list/all"`

	//SubBreeds returns an array of all the sub-breeds from a breed
	SubBreeds func(breed Breed) ([]SubBreed, error) `api:"/breed/%v/list"`

	//RandomImageURL returns a random image url of a dog.
	RandomImageURL func() (string, error) `api:"/breeds/image/random"`

	//RandomImageURLs returns random image urls of dogs.
	//Max number returned is 50.
	RandomImageURLs func(amount uint) ([]string, error) `api:"/breeds/image/random/%v"`

	//ImagesByBreed returns an array of all the images from a breed, e.g. hound
	ImageURLsByBreed func(breed Breed) ([]string, error) `api:"/breed/%v/images"`

	//RandomImageURLByBreed returns a random dog image from a breed, e.g. hound
	RandomImageURLByBreed func(breed Breed) ([]string, error) `api:"/breed/%v/images/random"`

	//RandomImageURLByBreed returns multiple random dog images from a breed, e.g. hound
	RandomImageURLsByBreed func(breed Breed, amount uint) ([]string, error) `api:"/breed/%v/images/random/%v"`
}

func init() {
	api.Connect(&API)
}
