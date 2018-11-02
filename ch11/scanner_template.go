//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt, strings, scanner  and unicode package

import (
	"fmt"
	"strings"
	"text/scanner"
	"unicode"
)

// main method
func main() {
	const example = "%var1 var2%"

	var scan scanner.Scanner
	scan.Init(strings.NewReader(example))
	scan.Filename = "initial"

	var token rune
	for token = scan.Scan(); token != scanner.EOF; token = scan.Scan() {
		fmt.Printf("%s: %s\n", scan.Position, scan.TokenText())
	}

	fmt.Println()
	scan.Init(strings.NewReader(example))
	scan.Filename = "changed"

	scan.IsIdentRune = func(character rune, integer int) bool {
		return character == '%' && integer == 0 || unicode.IsLetter(character) || unicode.IsDigit(character) && integer > 0
	}

	for token = scan.Scan(); token != scanner.EOF; token = scan.Scan() {
		fmt.Printf("%s: %s\n", scan.Position, scan.TokenText())
	}

}
