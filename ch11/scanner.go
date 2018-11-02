//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt, strings and scanner package
import (
	"fmt"
	"strings"
	"text/scanner"
)

//main method
func main() {
	const example = `
// code to be scanned.
if b < 19 {
	textToBeParsed = str
}`

	var scannerInstance scanner.Scanner
	scannerInstance.Init(strings.NewReader(example))
	scannerInstance.Filename = "example"

	var tok rune
	for tok = scannerInstance.Scan(); tok != scanner.EOF; tok = scannerInstance.Scan() {
		fmt.Printf("%s: %s\n", scannerInstance.Position, scannerInstance.TokenText())
	}

}
