package rest

import (
	"encoding/json"
	"io"
)

//Protocol implements a JSON api.Protocol.
type Protocol struct{}

//EncodeValue implements api.protocol.EncodeValue
func (p Protocol) EncodeValue(writer io.Writer, value interface{}) error {
	return json.NewEncoder(writer).Encode(value)
}

//DecodeValue implements api.protocol.DecodeValue
func (p Protocol) DecodeValue(reader io.Reader, value interface{}) error {
	return json.NewDecoder(reader).Decode(value)
}
