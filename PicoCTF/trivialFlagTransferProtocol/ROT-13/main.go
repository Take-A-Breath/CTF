package main

import (
	"fmt"
	"strings"
	"unicode"
)

// Writiing a function that can take an argument and encode/decode ROT-13
// Applying ROT13 to a piece of text merely requires examining its alphabetic characters and replacing each one by the letter
// 13 places further along in the alphabet, wrapping back to the beginning if necessary

func rot13(r rune) rune {
	uc := unicode.ToUpper(r)

	if uc >= 'A' && uc <= 'Z' {
		if uc >= 'M' {
			return r - 13
		} else {
			return r + 13
		}
	}
	return r
}

func main() {
	// Eventually refactor to take os.Args instead of altering script with value
	input := "VHFRQGURCEBTENZNAQUVQVGJVGU-QHRQVYVTRAPR.PURPXBHGGURCUBGBF"
	mapped := strings.Map(rot13, input)

	fmt.Println("Input: " + input)
	fmt.Println("Output: " + mapped)
}
