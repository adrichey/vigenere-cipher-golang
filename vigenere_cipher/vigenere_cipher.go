package vigenere_cipher

import (
	"bytes"
	"errors"
)

const alphabet string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

type TransformerFunction func(byte, int) byte

type Cipher struct {
	encodingMap map[int]byte
	letterMap   map[byte]int
	secretKey   []byte
}

func (c *Cipher) createLetterMap() {
	characters := []byte(alphabet)
	c.letterMap = make(map[byte]int)
	for k, v := range characters {
		c.letterMap[v] = k
	}
}

func (c *Cipher) createEncodingMap() {
	characters := []byte(alphabet)
	c.encodingMap = make(map[int]byte)
	for k, v := range characters {
		c.encodingMap[k] = v
	}
}

func (c *Cipher) transformInput(input string, transformer TransformerFunction) (string, error) {
	s := bytes.ToUpper([]byte(input))
	err := c.validateInput(s)

	if err != nil {
		return "", err
	}

	skIndex := 0
	transformedText := make([]byte, len(input))

	for i, b := range s {
		if 'A' <= b && b <= 'Z' {
			transformedText[i] = transformer(b, skIndex)

			if skIndex == len(c.secretKey)-1 {
				skIndex = 0
			} else {
				skIndex++
			}
		} else {
			transformedText[i] = b
		}
	}

	return string(transformedText), nil
}

func (c *Cipher) Encode(input string) (string, error) {
	return c.transformInput(input, c.getEncodedByte)
}

func (c *Cipher) getEncodedByte(b byte, secretKeyIndex int) byte {
	shiftKey := c.secretKey[secretKeyIndex]
	shiftOffset := c.letterMap[shiftKey]

	letterIndex := c.letterMap[b]
	shiftTarget := (letterIndex + shiftOffset) % len(c.encodingMap)

	return c.encodingMap[shiftTarget]
}

func (c *Cipher) Decode(input string) (string, error) {
	return c.transformInput(input, c.getDecodedByte)
}

func (c *Cipher) getDecodedByte(b byte, _ int) byte {
	letterIndex := c.letterMap[b]
	shiftTarget := (letterIndex) % len(c.encodingMap)

	return c.encodingMap[shiftTarget]
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
	c.createEncodingMap()
	c.secretKey = bytes.ToUpper(sk)

	return c, nil
}
