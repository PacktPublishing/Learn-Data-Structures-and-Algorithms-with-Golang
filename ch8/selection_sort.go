//main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
)

// Selection Sorter method
func SelectionSorter(elements []int) {

	var i int
	for i = 0; i < len(elements)-1; i++ {
		var min int
		min = i
		var j int
		for j = i + 1; j <= len(elements)-1; j++ {
			if elements[j] < elements[min] {
				min = j
			}
		}
		swap(elements, i, min)
	}
}

// swap method
func swap(elements []int, i int, j int) {
	var temp int
	temp = elements[j]
	elements[j] = elements[i]
	elements[i] = temp
}

//main method
func main() {
	var elements []int
	elements = []int{11, 4, 18, 6, 19, 21, 71, 13, 15, 2}
	fmt.Println("Before Sorting ", elements)
	SelectionSorter(elements)
	fmt.Println("After Sorting", elements)
}
