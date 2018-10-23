//main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
	"sort"
)

// main method
func main() {
	var elements []int
	elements = []int{1, 3, 16, 10, 45, 31, 28, 36, 45, 75}
	var element int
	element = 36

	var i int

	i = sort.Search(len(elements), func(i int) bool { return elements[i] >= element })
	if i < len(elements) && elements[i] == element {
		fmt.Printf("found element %d at index %d in %v\n", element, i, elements)
	} else {
		fmt.Printf("element %d not found in %v\n", element, elements)
	}
}
