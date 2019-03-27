//main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt and bytes package
import (
	"fmt"
)

//bubble Sorter method
func bubbleSorter(integers [11]int) {

	var num int
	num = 11
	var isSwapped bool
	isSwapped = true
	for isSwapped {
		isSwapped = false
		var i int
		for i = 1; i < num; i++ {
			if integers[i-1] > integers[i] {

				var temp = integers[i]
				integers[i] = integers[i-1]
				integers[i-1] = temp
				isSwapped = true
			}
		}
	}
	fmt.Println(integers)
}

// main method
func main() {
	var integers [11]int = [11]int{31, 13, 12, 4, 18, 16, 7, 2, 3, 0, 10}
	fmt.Println("Bubble Sorter")
	bubbleSorter(integers)

}
