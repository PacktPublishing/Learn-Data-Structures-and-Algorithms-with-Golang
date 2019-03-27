//main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt and bytes package
import (
	"fmt"
)

//factorial method
func Factor(num int) int {
	if num <= 1 {
		return 1
	}
	return num * Factor(num-1)
}

//main method
func main() {
	var num int = 5
	fmt.Println("Factorial: %d is %d", num, Factor(num))
}
