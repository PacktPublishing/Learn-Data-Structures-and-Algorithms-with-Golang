//main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
	"strconv"
)

// hash method
func hash(str string) int {

	var constant int

	var strint int

	constant = 42
	var err error
	strint, err = strconv.Atoi(str)

	constant = constant * strint

	fmt.Println(strint)

	fmt.Println(err)
	return constant

}

// main method
func main() {

	var str string

	str = "checkforhash"

	var hashCode int

	hashCode = hash(str)

	fmt.Println(hashCode)

}
