//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt and strings packages

import (
	"fmt"
	"strings"
)

// CheckPalindrome method
func CheckPalindrome(str string) string {
	var mid int
	mid = len(str) / 2
	var last int
	last = len(str) - 1
	var i int
	for i = 0; i < mid; i++ {
		if str[i] != str[last-i] {
			return "is not a Palimdrome."
		}
	}
	return "is a Palindrome"
}

// main method
func main() {

	var word string
	word = "madam"
	word = strings.ToLower(word)
	fmt.Println(word, CheckPalindrome(word))

	word = "structure"
	word = strings.ToLower(word)
	fmt.Println(word, CheckPalindrome(word))
}
