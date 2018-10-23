//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt package
import (
	"fmt"
)

// main method
func main() {

	var arr = [5]int{1, 2, 4, 5, 6}

	var i int
	for i = 0; i < len(arr); i++ {

		fmt.Println("printing elements ", arr[i])

	}

	var value int
	for i, value = range arr {

		fmt.Println(" range ", value)

	}

	for _, value = range arr {

		fmt.Println("blank range", value)

	}

}
