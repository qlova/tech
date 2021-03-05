//Package dog provides dog breeds and URLs to images of dogs.
package dog

import (
	"encoding/json"
	"errors"
	"io"
	"strings"

	"qlova.tech/api"
	"qlova.tech/api/rest"
	"qlova.tech/enc"
)

//Protocol implements the Dog api.Protocol.
type Protocol struct {
	enc.TextFormat
}

//Decode implements api.Protocol.Decode.
func (message Protocol) Decode(value interface{}, reader io.Reader) error {
	var Message struct {
		Message json.RawMessage `json:"message"`
		Status  string          `json:"status"`
	}

	if err := message.TextFormat.Decode(&Message, reader); err != nil {
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
	rest.API `api:"https://dog.ceo/api"`

	Protocol Protocol

	//Breeds returns a map of dog breeds to sub-breeds.
	Breeds func() (Breeds, error) `api:"/breeds/list/all"`

	//SubBreeds returns an array of all the sub-breeds from a breed
	SubBreeds func(breed Breed) ([]SubBreed, error) `api:"/breed/%v/list"`

	//RandomImageURL returns a random image url of a dog.
	RandomImageURL func() (string, error) `api:"/breeds/image/random"`

	//RandomImage returns a random image of a dog.
	RandomImage func() (rest.Image, error) `api:"/breeds/image/random"`

	//RandomImageURLs returns random image urls of dogs.
	//Max number returned is 50.
	RandomImageURLs func(amount uint) ([]string, error) `api:"/breeds/image/random/%v"`

	//ImagesByBreed returns an array of all the images from a breed, e.g. hound
	ImageURLsByBreed func(breed Breed) ([]string, error) `api:"/breed/%v/images"`

	//RandomImageURLByBreed returns a random dog image url from a breed, e.g. hound
	RandomImageURLByBreed func(breed Breed) ([]string, error) `api:"/breed/%v/images/random"`

	//RandomImageByBreed returns a random dog image from a breed, e.g. hound
	RandomImageByBreed func(breed Breed) (rest.Image, error) `api:"/breed/%v/images/random"`

	//RandomImageURLByBreed returns multiple random dog image urls from a breed, e.g. hound
	RandomImageURLsByBreed func(breed Breed, amount uint) ([]string, error) `api:"/breed/%v/images/random/%v"`
}

func init() {
	API.Protocol.TextFormat = enc.JSON
	api.Connect(&API)
}
