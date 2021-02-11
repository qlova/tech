//Package key provides simple encryption functions.
package key

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
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
