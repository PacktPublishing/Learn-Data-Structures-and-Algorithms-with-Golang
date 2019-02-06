///main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt and list package
import (
	"container/list"
	"fmt"
)

// main method
func main() {
	var linkedList *list.List
	linkedList = list.New()
	var element *list.Element
	element = linkedList.PushBack(14)

	var frontElement *list.Element
	frontElement = linkedList.PushFront(1)
	linkedList.InsertBefore(6, element)
	linkedList.InsertAfter(5, frontElement)

	var currElement *list.Element
	for currElement = linkedList.Front(); currElement != nil; currElement = currElement.Next() {
		fmt.Println(currElement.Value)
	}

}
