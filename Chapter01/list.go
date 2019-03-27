//main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt and container list packages
import (
	"container/list"
	"fmt"
)

// main method
func main() {
	var intList list.List
	intList.PushBack(11)
	intList.PushBack(23)
	intList.PushBack(34)

	for element := intList.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value.(int))
	}
}
