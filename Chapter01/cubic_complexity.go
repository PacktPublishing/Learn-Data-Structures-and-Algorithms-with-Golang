//main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
)

// main method
func main() {

	var k, l, m int
	var arr [10][10][10]int
	for k = 0; k < 10; k++ {

		for l = 0; l < 10; l++ {

			for m = 0; m < 10; m++ {

				arr[k][l][m] = 1

				fmt.Println("Element value ", k, l, m, " is", arr[k][l][m])
			}

		}

	}
}
