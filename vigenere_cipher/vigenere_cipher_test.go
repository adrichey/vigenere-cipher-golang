package vigenere_cipher

import (
	"testing"
)

/*
func TestInvalidNewCipher(t *testing.T) {
	testCases := []string{".", "\t", "\r", "\n", "123", "~"}

	for _, testCase := range testCases {
		vc, err := NewCipher(testCase)

		if vc != nil || err == nil {
			quotedTestCase := strconv.Quote(testCase)
			t.Fatalf(`NewCipher(%v) = %q, %v, want nil, error`, quotedTestCase, vc, err)
		}
	}
}

func TestValidNewCipher(t *testing.T) {
	testCases := []string{"sdfhbe", "blwertbwertSEFSfweEEWEFWESSSSSSSS", "JQWBFQJ"}

	for _, testCase := range testCases {
		vc, err := NewCipher(testCase)

		if vc == nil || err != nil {
			quotedTestCase := strconv.Quote(testCase)
			t.Fatalf(`NewCipher(%v) = %q, %v, want Cipher struct, nil`, quotedTestCase, vc, err)
		}
	}
}

func TestInvalidEncodeInput(t *testing.T) {
	testCases := []string{"This is invalid.", "1234"}

	for _, testCase := range testCases {
		vc, _ := NewCipher("TEST")
		encodedMessage, err := vc.Encode(testCase)

		if encodedMessage != "" || err == nil {
			quotedTestCase := strconv.Quote(testCase)
			t.Fatalf(`vc.Encode(%v) = %q, %v, want "", error`, quotedTestCase, encodedMessage, err)
		}
	}
}

func TestEncode(t *testing.T) {
	testCases := []map[string]string{
		{
			"secretKey":      "TESTKEY",
			"decodedMessage": "THIS IS A DECODED MESSAGE",
			"encodedMessage": "MLAL SW Y WIUHNIB FIKLKKC",
		},
		{
			"secretKey":      "KEYTEST",
			"decodedMessage": "THIS IS A DECODED MESSAGE",
			"encodedMessage": "DLGL MK T NIAHHWW WIQLEYX",
		},
		{
			"secretKey":      "Ozymandias",
			"decodedMessage": "WE ATTACK AT DAWN WE ATTACK WITH KINDNESS",
			"encodedMessage": "KD YFTNFS AL RZUZ WR DBTSQJ UUTU NQNVBDQE",
		},
	}

	for _, testCase := range testCases {
		vc, _ := NewCipher(testCase["secretKey"])
		encodedMessage, err := vc.Encode(testCase["decodedMessage"])

		if encodedMessage != testCase["encodedMessage"] || err != nil {
			t.Fatalf(`vc.Encode("%v") = %q, %v, want "%v", nil`, testCase["decodedMessage"], encodedMessage, err, testCase["encodedMessage"])
		}
	}
}

func TestInvalidDecodeInput(t *testing.T) {
	testCases := []string{"This is invalid.", "1234"}

	for _, testCase := range testCases {
		vc, _ := NewCipher("TEST")
		decodedMessage, err := vc.Decode(testCase)

		if decodedMessage != "" || err == nil {
			quotedTestCase := strconv.Quote(testCase)
			t.Fatalf(`vc.Decode(%v) = %q, %v, want "", error`, quotedTestCase, decodedMessage, err)
		}
	}
}
*/

func TestDecode(t *testing.T) {
	testCases := []map[string]string{
		{
			"secretKey":      "TESTKEY",
			"decodedMessage": "THIS IS A DECODED MESSAGE",
			"encodedMessage": "MLAL SW Y WIUHNIB FIKLKKC",
		},
		{
			"secretKey":      "KEYTEST",
			"decodedMessage": "THIS IS A DECODED MESSAGE",
			"encodedMessage": "DLGL MK T NIAHHWW WIQLEYX",
		},
		{
			"secretKey":      "Ozymandias",
			"decodedMessage": "WE ATTACK AT DAWN WE ATTACK WITH KINDNESS",
			"encodedMessage": "KD YFTNFS AL RZUZ WR DBTSQJ UUTU NQNVBDQE",
		},
	}

	for _, testCase := range testCases {
		vc, _ := NewCipher(testCase["secretKey"])
		decodedMessage, err := vc.Decode(testCase["encodedMessage"])

		if decodedMessage != testCase["decodedMessage"] || err != nil {
			t.Fatalf(`vc.Decode("%v") = %q, %v, want "%v", nil`, testCase["encodedMessage"], decodedMessage, err, testCase["decodedMessage"])
		}
	}
}
