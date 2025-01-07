package main

import (
	"fmt"
	"strings"

	"github.com/adrichey/vigenere-cipher-golang/file_manager"
	"github.com/adrichey/vigenere-cipher-golang/vigenere_cipher"
)

const inputFile string = "input.txt"
const outputFile string = "output.txt"
const decodedFile string = "decoded.txt"

func main() {
	vc, err := vigenere_cipher.NewCipher("Ozymandias")

	if err != nil {
		panic(err)
	}

	fm := file_manager.New(inputFile, outputFile)

	// Read plaintext from file
	lines, err := fm.ReadLines()

	if err != nil {
		panic(err)
	}

	// Encode the message from the file
	input := strings.Join(lines, "\n")
	output, err := vc.Encode(input)

	if err != nil {
		panic(err)
	}

	// Write the encoded output to a file
	err = fm.WriteToFile(output)

	if err != nil {
		panic(err)
	}

	fmt.Println("Encrypted message added to", outputFile)

	fm = file_manager.New(outputFile, decodedFile)

	// Read encoded message from file
	lines, err = fm.ReadLines()

	if err != nil {
		panic(err)
	}

	// Decode the message from the file
	input = strings.Join(lines, "\n")
	output, err = vc.Decode(input)

	if err != nil {
		panic(err)
	}

	// Write the decoded output to a file
	err = fm.WriteToFile(output)

	if err != nil {
		panic(err)
	}

	fmt.Println("Decoded message added to", decodedFile)
}
