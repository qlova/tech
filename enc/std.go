package enc

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strconv"

	"encoding/ascii85"
	"encoding/asn1"
	"encoding/base32"
	"encoding/base64"
	"encoding/gob"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"encoding/xml"
)

//JSON TextFormat.
var JSON TextFormat = WithJSON(json.NewEncoder, json.NewDecoder)

//WithJSON returns a JSON TextFormat with the given encoder and decoder.
func WithJSON(encoder func(io.Writer) *json.Encoder, decoder func(io.Reader) *json.Decoder) TextFormat {
	return jsonFormat{encoder, decoder}
}

type jsonFormat struct {
	Encoder func(io.Writer) *json.Encoder
	Decoder func(io.Reader) *json.Decoder
}

func (j jsonFormat) Format() string { return "json" }

func (j jsonFormat) TextEncode(value interface{}, writer io.Writer) error {
	return j.Encoder(writer).Encode(value)
}

func (j jsonFormat) TextDecode(value interface{}, reader io.Reader) error {
	return j.Decoder(reader).Decode(value)
}

func (j jsonFormat) Encode(value interface{}, writer io.Writer) error {
	return j.Encoder(writer).Encode(value)
}

func (j jsonFormat) Decode(value interface{}, reader io.Reader) error {
	return j.Decoder(reader).Decode(value)
}

//XML TextFormat.
var XML TextFormat = WithXML(xml.NewEncoder, xml.NewDecoder)

//WithXML returns a XML TextFormat with the given encoder and decoder.
func WithXML(encoder func(io.Writer) *xml.Encoder, decoder func(io.Reader) *xml.Decoder) TextFormat {
	return xmlFormat{encoder, decoder}
}

type xmlFormat struct {
	Encoder func(io.Writer) *xml.Encoder
	Decoder func(io.Reader) *xml.Decoder
}

func (j xmlFormat) Format() string { return "xml" }

func (j xmlFormat) TextEncode(value interface{}, writer io.Writer) error {
	return j.Encoder(writer).Encode(value)
}

func (j xmlFormat) TextDecode(value interface{}, reader io.Reader) error {
	return j.Decoder(reader).Decode(value)
}

func (j xmlFormat) Encode(value interface{}, writer io.Writer) error {
	return j.Encoder(writer).Encode(value)
}

func (j xmlFormat) Decode(value interface{}, reader io.Reader) error {
	return j.Decoder(reader).Decode(value)
}

//ASN1 TextFormat.
var ASN1 TextFormat = asn1Format{}

type asn1Format struct{}

func (j asn1Format) Format() string { return "asn.1" }

func (j asn1Format) TextEncode(value interface{}, writer io.Writer) error {
	return j.Encode(value, writer)
}

func (j asn1Format) TextDecode(value interface{}, reader io.Reader) error {
	return j.Decode(value, reader)
}

func (j asn1Format) Encode(value interface{}, writer io.Writer) error {
	b, err := asn1.Marshal(value)
	if err != nil {
		return err
	}

	fmt.Fprint(writer, strconv.Itoa(len(b)), ' ')

	_, err = writer.Write(b)
	if err != nil {
		return err
	}

	return nil
}

func (j asn1Format) Decode(value interface{}, reader io.Reader) error {
	var bytes string
	var single = []byte{0}
	for {
		if _, err := reader.Read(single); err != nil {
			return err
		}
		if single[0] == ' ' {
			break
		}
		bytes += string(rune(single[0]))
	}

	length, err := strconv.Atoi(bytes)
	if err != nil {
		return err
	}

	var b = make([]byte, length)

	_, err = reader.Read(b)
	if err != nil {
		return err
	}

	_, err = asn1.Unmarshal(b, value)
	if err != nil {
		return err
	}

	return nil

}

//Base64 returns a standard base64 encoding of the given format.
func Base64(f Format) TextFormat {
	return WithBase64(base64.StdEncoding, f)
}

//WithBase64 returns a base64 encoding of the given kind to wrapping the
//given format.
func WithBase64(kind *base64.Encoding, f Format) TextFormat {
	return wrapper{
		name: "base64(" + f.Format() + ")",
		wrapReader: func(r io.Reader) io.Reader {
			return base64.NewDecoder(kind, r)
		},
		wrapWriter: func(r io.Writer) io.Writer {
			return base64.NewEncoder(kind, r)
		},
		format: f,
	}
}

//Base32 returns a standard base32 encoding of the given format.
func Base32(f Format) TextFormat {
	return WithBase64(base64.StdEncoding, f)
}

//WithBase32 returns a base32 encoding of the given kind to wrapping the
//given format.
func WithBase32(kind *base32.Encoding, f Format) TextFormat {
	return wrapper{
		name: "base32(" + f.Format() + ")",
		wrapReader: func(r io.Reader) io.Reader {
			return base32.NewDecoder(kind, r)
		},
		wrapWriter: func(r io.Writer) io.Writer {
			return base32.NewEncoder(kind, r)
		},
		format: f,
	}
}

//ASCII85 returns a ascii85 encoding of the given format.
func ASCII85(f Format) TextFormat {
	return wrapper{
		name:       "ascii85(" + f.Format() + ")",
		wrapReader: ascii85.NewDecoder,
		wrapWriter: func(r io.Writer) io.Writer {
			return ascii85.NewEncoder(r)
		},
		format: f,
	}
}

//Gob Format.
var Gob Format = gobFormat{}

type gobFormat struct{}

func (j gobFormat) Format() string { return "gob" }

func (j gobFormat) Encode(value interface{}, writer io.Writer) error {
	return gob.NewEncoder(writer).Encode(value)
}

func (j gobFormat) Decode(value interface{}, reader io.Reader) error {
	return gob.NewDecoder(reader).Decode(value)
}

//Hex returns a hex encoding of the given format.
func Hex(f Format) TextFormat {
	return wrapper{
		name:       "hex(" + f.Format() + ")",
		wrapReader: hex.NewDecoder,
		wrapWriter: hex.NewEncoder,
		format:     f,
	}
}

//Pem returns a pem encoding where the body format is
//encoded to base64.
func Pem(f Format) TextFormat {
	return pemFormat{f}
}

type pemFormat struct {
	internal Format
}

func (j pemFormat) Format() string { return "pem(" + j.internal.Format() + ")" }

func (j pemFormat) TextEncode(value interface{}, writer io.Writer) error {
	return j.Encode(value, writer)
}

func (j pemFormat) TextDecode(value interface{}, reader io.Reader) error {
	return j.Decode(value, reader)
}

func (j pemFormat) Encode(value interface{}, writer io.Writer) error {
	var buffer bytes.Buffer
	if err := j.internal.Encode(value, &buffer); err != nil {
		return err
	}

	fmt.Fprint(writer, strconv.Itoa(buffer.Len()), ' ')

	return pem.Encode(writer, &pem.Block{
		Bytes: buffer.Bytes(),
	})
}

func (j pemFormat) Decode(value interface{}, reader io.Reader) error {
	var bytes string
	var single = []byte{0}
	for {
		if _, err := reader.Read(single); err != nil {
			return err
		}
		if single[0] == ' ' {
			break
		}
		bytes += string(rune(single[0]))
	}

	length, err := strconv.Atoi(bytes)
	if err != nil {
		return err
	}

	var b = make([]byte, length)

	_, err = reader.Read(b)
	if err != nil {
		return err
	}

	block, _ := pem.Decode(b)
	if block == nil {
		return errors.New("missing pem block")
	}

	return nil
}
