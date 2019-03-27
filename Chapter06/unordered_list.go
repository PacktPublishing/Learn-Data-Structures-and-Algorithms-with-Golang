//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt package
import (
	"fmt"
)

//Node class
type Node struct {
	property int
	nextNode *Node
}

// UnOrderedList class
type UnOrderedList struct {
	headNode *Node
}

//AddToHead method of UnOrderedList class
func (UnOrderedList *UnOrderedList) AddToHead(property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil

	if UnOrderedList.headNode != nil {
		node.nextNode = UnOrderedList.headNode
	}

	UnOrderedList.headNode = node

}

//IterateList method iterates over UnOrderedList
func (UnOrderedList *UnOrderedList) IterateList() {

	var node *Node
	for node = UnOrderedList.headNode; node != nil; node = node.nextNode {
		fmt.Println(node.property)

	}
}

// main method
func main() {

	var unOrderedList UnOrderedList

	unOrderedList = UnOrderedList{}

	unOrderedList.AddToHead(1)
	unOrderedList.AddToHead(3)
	unOrderedList.AddToHead(5)
	unOrderedList.AddToHead(7)
	unOrderedList.IterateList()

}
