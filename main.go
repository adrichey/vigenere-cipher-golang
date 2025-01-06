package main

import (
	"fmt"

	"github.com/adrichey/vigenere-cipher-golang/vigenereCipher"
)

func main() {
	vigenereCipher, err := vigenereCipher.NewCipher("Ozymandias")

	if err != nil {
		panic(err)
	}

	// TODO: Fetch this input from an input file eventually
	input := "We attack at dawn We attack with kindness"
	output, err := vigenereCipher.Encode(input)

	if err != nil {
		panic(err)
	}

	outputTest := "KD YFTNFS AL RZUZ WR DBTSQJ UUTU NQNVBDQE" // TODO - Turn this into a test
	fmt.Println(outputTest)
	fmt.Println(output)
	fmt.Println(output == outputTest)
}
