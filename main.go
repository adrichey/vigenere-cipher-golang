package main

import (
	"fmt"

	"github.com/adrichey/vigenere-cipher-golang/vigenereCipher"
)

func main() {
	vigenereCipher, err := vigenereCipher.NewCipher("THISISTHEKEY")

	if err != nil {
		panic(err)
	}

	// TODO: Fetch this input from an input file eventually
	input := "We attack at dawn We attack with kindness"
	output, err := vigenereCipher.Encode(input)

	if err != nil {
		panic(err)
	}

	fmt.Println(output)
}
