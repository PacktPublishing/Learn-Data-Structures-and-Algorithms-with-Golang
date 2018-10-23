//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt package
import (
	"fmt"
)

// main method
func main() {

	var arr = []int{5, 6, 7, 8, 9}
	var slic1 = arr[:3]
	fmt.Println("slice1", slic1)
	var slic2 = arr[1:5]

	fmt.Println("slice2", slic2)
	var slic3 = append(slic2, 12)

	fmt.Println("slice3", slic3)

}
