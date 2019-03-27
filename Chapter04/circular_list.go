//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing container/ring and fmt package
import (
	"container/ring"
	"fmt"
)

// main method
func main() {

	var integers []int
	integers = []int{1, 3, 5, 7}

	var circular_list *ring.Ring
	circular_list = ring.New(len(integers))

	var i int
	for i = 0; i < circular_list.Len(); i++ {
		circular_list.Value = integers[i]
		circular_list = circular_list.Next()
	}

	circular_list.Do(func(element interface{}) {
		fmt.Print(element, ",")
	})
	fmt.Println()

	for i = 0; i < circular_list.Len(); i++ {
		fmt.Print(circular_list.Value, ",")
		circular_list = circular_list.Prev()
	}
	fmt.Println()

	circular_list = circular_list.Move(2)
	circular_list.Do(func(element interface{}) {
		fmt.Print(element, ",")
	})
	fmt.Println()
}
