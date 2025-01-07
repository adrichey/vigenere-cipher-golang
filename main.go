package main

import (
	"fmt"
	"strings"

	"github.com/adrichey/vigenere-cipher-golang/vigenere_cipher"
)

func main() {
	vc, err := vigenere_cipher.NewCipher("Ozymandias")

	if err != nil {
		panic(err)
	}

	// TODO: Fetch this input from an input file eventually
	input := "We attack at dawn We attack with kindness"
	output, err := vc.Encode(input)

	if err != nil {
		panic(err)
	}

	outputTest := "KD YFTNFS AL RZUZ WR DBTSQJ UUTU NQNVBDQE" // TODO - Turn this into a test
	fmt.Println(outputTest)
	fmt.Println(output)
	fmt.Println(output == outputTest)

	decodedMessage, err := vc.Decode(outputTest)
	fmt.Println(decodedMessage)
	fmt.Println(strings.ToUpper(input))
	fmt.Println(decodedMessage == strings.ToUpper(input))
}
