//Package httpcat provides cat themed HTTP status images.
package httpcat

import (
	"image"
	"image/jpeg"
	"io"

	"qlova.tech/api"
	"qlova.tech/api/rest"
)

//Image of an HTTP status.
type Image struct {
	image.Image
}

//Decode decodes the image from the given reader.
func (i *Image) Decode(reader io.Reader) error {
	img, err := jpeg.Decode(reader)
	if err != nil {
		return err
	}

	i.Image = img

	return nil
}

//API implements api.Interface
var API struct {
	rest.API `rest:"https://http.cat"`

	//Image returns a cat-themed image for
	//the given HTTP status.
	Image func(code int) (Image, error) `rest:"/%v"`
}

func init() {
	api.Connect(&API)
}
