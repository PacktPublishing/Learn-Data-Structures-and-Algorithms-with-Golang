//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt, strings and scanner package
import (
	"fmt"
	"strings"
	"text/scanner"
)

// main method
func main() {

	const example = `ab	bb	cb	db
eb	fb	gb	hb
ib	jb	kb	lb
mb	nb	ob	pb`

	var (
		column   int
		row      int
		scan     scanner.Scanner
		strArray [4][4]string
	)
	scan.Init(strings.NewReader(example))
	scan.Whitespace ^= 1<<'\t' | 1<<'\n'

	var token rune
	for token = scan.Scan(); token != scanner.EOF; token = scan.Scan() {
		switch token {
		case '\n':
			row++
			column = 0
		case '\t':
			column++
		default:
			strArray[row][column] = scan.TokenText()
		}
	}

	fmt.Print(strArray)
	fmt.Println()

}
