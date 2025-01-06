package main

import (
	"github.com/adrichey/vigenere-cipher-golang/vigenereCipher"
)

func main() {
	vigenereCipher := vigenereCipher.NewCipher()

	vigenereCipher.PrintLetterMap()
}
