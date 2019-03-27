///main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
)

// Node struct
type Node struct {
	nextNode *Node
	property rune
}

// create List method
func CreateLinkedList() *Node {

	var headNode *Node
	headNode = &Node{nil, 'a'}

	var currNode *Node
	currNode = headNode

	var i rune
	for i = 'b'; i <= 'z'; i++ {
		var node *Node
		node = &Node{nil, i}
		currNode.nextNode = node
		currNode = node
	}

	return headNode
}

//Stringify  List method
func StringifyList(nodeList *Node) {

	var currNode *Node
	currNode = nodeList
	for {
		fmt.Printf("%c", currNode.property)

		if currNode.nextNode != nil {
			currNode = currNode.nextNode
		} else {
			break
		}
	}
	fmt.Println("")
}

// Reverse List method
func ReverseLinkedList(nodeList *Node) *Node {

	var currNode *Node
	currNode = nodeList
	var topNode *Node = nil
	for {
		if currNode == nil {
			break
		}
		var tempNode *Node
		tempNode = currNode.nextNode
		currNode.nextNode = topNode
		topNode = currNode
		currNode = tempNode
	}

	return topNode
}

// main method
func main() {

	var linkedList = CreateLinkedList()
	StringifyList(linkedList)
	StringifyList(ReverseLinkedList(linkedList))
}
