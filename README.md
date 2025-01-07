# Vigenère cipher decoder and encoder

You can find a great write-up about how this cipher works on [Wikipedia](https://en.wikipedia.org/wiki/Vigen%C3%A8re_cipher).

## How to use this script:

Flags
- `-i`: The input text file to encode/decode
- `-o`: The text file to dump the encoded/decoded message (optional, defaults to output.txt)
- `-k`: The alphabetic secret key used to encrypt the message
- `-e`: Encode the message from the input text file (incompatible with `-d` flag)
- `-d`: Decode the message from the input text file (incompatible with `-e` flag)

Please note that the secret key (-k) must only contain alphabetic characters A-Z.

Additionally, your input text file must only contain alphabetic characters A-Z, spaces, newlines, and tabs.

### Encoding Example:
`./main -e -i example_input.txt -o encoded_output.txt -k CIPHERKEY`

You can find an example input file that you can use to encode in [example_input.txt](example_input.txt).

### Decoding Example:
`./main -d -i encoded_input.txt -o decoded_output.txt -k CIPHERKEY`
