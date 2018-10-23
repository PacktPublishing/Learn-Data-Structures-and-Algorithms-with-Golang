//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt package
import (
	"fmt"
)

// main method
func main() {
	var TwoDArray [8][8]int

	TwoDArray[3][6] = 18

	TwoDArray[7][4] = 3

	fmt.Println(TwoDArray)

}
