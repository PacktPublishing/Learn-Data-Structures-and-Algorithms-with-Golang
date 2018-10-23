///main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
)

// main method
func main() {

	var matrix = [3][3]int{
		{1, 0, 0},
		{1, 1, 0},
		{2, 1, 1}}

	fmt.Println(matrix)
}
