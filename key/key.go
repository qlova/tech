//Package key provides simple encryption functions.
package key

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

//Code is an ASCII-based representation of a standard US QWERTY keyboard keyboard layout.
type Code int8

type State [2]uint64

const (
	Left  Code = -1
	Any   Code = 0
	Right Code = 1
	Home  Code = 2
	End   Code = 3

	Pause Code = 5

	Backspace Code = 8
	Tab       Code = 9
	Enter     Code = 10
	PageDown  Code = 12
	PageUp    Code = 13
	CapsLock  Code = 15

	Up   Code = 14
	Down Code = 15

	Shift      Code = 17
	Ctrl       Code = 18
	Super      Code = 19
	Option     Code = 19
	Command    Code = 20
	Alt        Code = 20
	ScrollLock Code = 21

	Insert      Code = 26
	Esc         Code = 27
	PrintScreen Code = 30
	Space       Code = 32

	Number Code = 48
	Letter Code = 65

	Function Code = 100
	Del      Code = 127
)

//AES is a 256bit AES key.
type AES [32]byte

//Encrypt encrypts the data and returns the resulting encrypted bytes.
func (key AES) Encrypt(data []byte) []byte {
	c, err := aes.NewCipher(key[:])
	if err != nil {
		panic(err.Error())
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		panic(err.Error())
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	return gcm.Seal(nonce, nonce, data, nil)
}

//Decrypt returns the corresponding plaintext of data or nil.
func (key AES) Decrypt(data []byte) []byte {
	c, err := aes.NewCipher(key[:])
	if err != nil {
		panic(err.Error())
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		panic(err.Error())
	}

	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return nil
	}

	nonce, ciphertext := []byte(data[:nonceSize]), []byte(data[nonceSize:])
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil
	}

	return plaintext
}
