package http

import (
	"io"
	"net/http"

	"qlova.tech/uri"
)

type opener string

func (o opener) Open(location uri.String) (io.ReadCloser, error) {
	resp, err := http.Get(string(o) + location.String())
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func init() {
	uri.RegisterOpenerFor("http", opener("http"))
	uri.RegisterOpenerFor("https", opener("https"))
}
