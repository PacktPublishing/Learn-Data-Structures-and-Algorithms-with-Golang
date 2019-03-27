//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt package
import (
	"fmt"
)

//h function which returns the product of parameters x and y
func h(x int, y int) int {

	return x * y
}

// g function which returns x and y parameters after modification
func g(l int, m int) (x int, y int) {
	x = 2 * l
	y = 4 * m
	return
}

// main method
func main() {

	fmt.Println(h(g(1, 2)))
}
