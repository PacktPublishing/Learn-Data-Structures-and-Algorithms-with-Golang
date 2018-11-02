//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt package
import (
	"fmt"
)
// check method
func Check(value int) int {
	fmt.Println("Checking value")
	if value > 0 {
		return 12
	}
	return 3
}
// method C
func C(value int) int {
	return -10
}
