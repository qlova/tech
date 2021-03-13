package file

import (
	"io"
	"os"
	"strings"

	"qlova.tech/uri"
)

type opener struct{}

func (opener) Open(location uri.String) (io.ReadCloser, error) {
	path := strings.TrimPrefix(location.String(), "file://")

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func init() {
	uri.RegisterOpenerFor("file", opener{})
}
