/*
 * Convert hex to base64
 *
 * The string:
 * 49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d
 *
 * Should produce:
 * SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t
 *
 * So go ahead and make that happen. You'll need to use this code for the rest
 * of the exercises.
 *
 * Cryptopals Rule Always operate on raw bytes, never on encoded strings. Only
 * use hex and base64 for pretty-printing.
 */

package main

import (
	"bytes"
	"fmt"
	mcc "github.com/dplbsd/mcclib"
)

// Symbols used in base64.
const test_ascii = "I'm killing your brain like a poisonous mushroom"
const test = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
const result = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

func main() {
	fmt.Printf("Test string: %q len(test): %d\n", test, len(test))
	fmt.Printf("Correct result: %q len(result): %d\n", result, len(result))

	base64 := mcc.Hex2base64(mcc.String2Hex(test))
	fmt.Printf("Our result: %s\n", base64)

	if cmp := bytes.Compare([]byte(result), base64); cmp != 0 {
		fmt.Println("Your code is wrong!")
	} else {
		fmt.Println("Your code is correct! :)")
	}
	return
}
