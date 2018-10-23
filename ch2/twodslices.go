//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt package
import (
	"fmt"
)

// main method
func main() {

	var rows int
	var cols int

	rows = 7
	cols = 9
	var twodslices = make([][]int, rows)

	var i int

	for i = range twodslices {

		twodslices[i] = make([]int, cols)
	}

	fmt.Println(twodslices)
}
