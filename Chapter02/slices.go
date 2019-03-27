//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt package
import (
	"fmt"
)

//twiceValue method given slice of int type
func twiceValue(slice []int) {

	var i int
	var value int

	for i, value = range slice {

		slice[i] = 2 * value

	}

}

// main method
func main() {

	var slice = []int{1, 3, 5, 6}
	twiceValue(slice)

	var i int

	for i = 0; i < len(slice); i++ {

		fmt.Println("new slice value", slice[i])
	}
}
