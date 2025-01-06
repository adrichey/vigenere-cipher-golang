package vigenereCipher

import (
	"errors"
)

type Cipher struct {
	letterMap map[string]int
	secretKey []byte
}

func (c *Cipher) createLetterMap() {
	characters := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	c.letterMap = make(map[string]int)
	for k, v := range characters {
		letter := string(v)
		c.letterMap[letter] = k
	}
}

func (c *Cipher) Encode(input string) (string, error) {
	err := c.validateInput([]byte(input))

	if err != nil {
		return "", err
	}

	// TODO - Implement convert text with cipher here
	return "ENCODED TEXT WILL APPEAR HERE", nil
}

func (c *Cipher) validateInput(s []byte) error {
	for _, b := range s {
		if !('a' <= b && b <= 'z') &&
			!('A' <= b && b <= 'Z') &&
			b != '\r' &&
			b != '\n' &&
			b != '\t' &&
			b != ' ' {
			return errors.New("input is invalid. It must only contain characters A-Z and spaces.")
		}
	}

	return nil
}

func NewCipher(secretKey string) (*Cipher, error) {
	// Validate the secretKey to make sure we can use it with this cipher
	sk := []byte(secretKey)
	for _, b := range sk {
		if !('a' <= b && b <= 'z') && !('A' <= b && b <= 'Z') {
			return nil, errors.New("secretKey is invalid while creating new Cipher instance. It must only contain characters A-Z.")
		}
	}

	c := &Cipher{}
	c.createLetterMap()
	c.secretKey = sk

	return c, nil
}
