package rest

import (
	"encoding/json"
	"image"
	"net/http"
)

//Image is an image that can unmarshal itself from a json URL string.
type Image struct {
	image.Image
}

//UnmarshalJSON unmarshals the image from the given image url.
func (i *Image) UnmarshalJSON(data []byte) error {
	var url string
	if err := json.Unmarshal(data, &url); err != nil {
		return err
	}

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return err
	}

	i.Image = img

	return nil
}
