//main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
)

// main method
func main() {
	var m [10]int
	var k int

	for k = 0; k < 10; k++ {
		m[k] = k * 200

		fmt.Printf("Element[%d] = %d\n", k, m[k])
	}
}
