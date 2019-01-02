//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing flag and fmt packages
import (
	"fmt"
)

//Table class
type Table struct {
	rows int
	cols int
}

// New method
func New(rows int, cols int) Table {
	var t Table
	t = Table{rows, cols}
	return t
}

// main method
func main() {

	var table Table

	table = New(3, 4)

	fmt.Println(table)

}
