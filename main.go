package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"

	"github.com/adrichey/vigenere-cipher-golang/file_manager"
	"github.com/adrichey/vigenere-cipher-golang/vigenere_cipher"
)

var encode, decode, help bool
var inputFile, outputFile, secretKey string

func init() {
	flag.BoolVar(&decode, "d", false, "Specifies you would like to decode a message from an input text file")
	flag.BoolVar(&encode, "e", false, "Specifies you would like to encode a message from an input text file")

	flag.StringVar(&inputFile, "i", "", "Path to the input text file")
	flag.StringVar(&outputFile, "o", "output.txt", "Path to the the output text file")
	flag.StringVar(&secretKey, "k", "", "Your secret cipher key")

	flag.BoolVar(&help, "help", false, "Help")

	flag.Parse()
}

func main() {
	if help {
		displayHelp()
		return
	}

	err := checkForFlagErrors()

	if err != nil {
		fmt.Println("ERROR", err.Error())
		fmt.Println()
		displayHelp()
		return
	}

	vc, err := vigenere_cipher.NewCipher(secretKey)

	if err != nil {
		fmt.Println(err)
		return
	}

	fm := file_manager.New(inputFile, outputFile)

	lines, err := fm.ReadLines()

	if err != nil {
		fmt.Println(err)
		return
	}

	input := strings.Join(lines, "\n")

	var transformer func(string) (string, error)
	if encode {
		transformer = vc.Encode
	} else if decode {
		transformer = vc.Decode
	}

	output, err := transformer(input)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = fm.WriteToFile(output)

	if err != nil {
		fmt.Println(err)
		return
	}

	var successMessage string
	if encode {
		successMessage = "Encoded message exported to"
	} else if decode {
		successMessage = "Decoded message exported to"
	}

	fmt.Println(successMessage, outputFile)
}

func checkForFlagErrors() error {
	if inputFile == "" {
		return errors.New("missing -i: input file is a required flag")
	}

	if secretKey == "" {
		return errors.New("missing -k: cipher secrey key is a required flag")
	}

	if encode && decode {
		return errors.New("you cannot use the -e and -d flags together")
	}

	if !encode && !decode {
		return errors.New("missing -e or -d flag to specify whether to encode/decode")
	}

	return nil
}

func displayHelp() {
	fmt.Println("How to use this script:")
	fmt.Println("-i: the input text file to encode/decode")
	fmt.Println("-o: the text file to dump the encoded/decoded message (optional, defaults to output.txt)")
	fmt.Println("-k: the alphabetic secret key used to encrypt the message")
	fmt.Println("-e: encode the message from the input text file (incompatible with -d flag)")
	fmt.Println("-d: decode the message from the input text file (incompatible with -e flag)")
	fmt.Println()
	fmt.Println("Encoding Example:")
	fmt.Println("./main -e -i plaintext_input.txt -o encoded_output.txt -k CIPHERKEY")
	fmt.Println()
	fmt.Println("Decoding Example:")
	fmt.Println("./main -d -i encoded_input.txt -o decoded_output.txt -k CIPHERKEY")
	fmt.Println()
	fmt.Println("Please note that the secret key (-k) must only contain alphabetic characters A-Z")
	fmt.Println("Additionally, your input text file must only contain alphabetic characters A-Z, spaces, newlines, and tabs")
	fmt.Println()
}
