//main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
)

// Tree struct
type Tree struct {
	LeftNode  *Tree
	Value     int
	RightNode *Tree
}

// Tree insert method for inserting at m position
func (tree *Tree) insert(m int) {
	if tree != nil {

		if tree.LeftNode == nil {
			tree.LeftNode = &Tree{nil, m, nil}
		} else {
			if tree.RightNode == nil {
				tree.RightNode = &Tree{nil, m, nil}
			} else {

				if tree.LeftNode != nil {

					tree.LeftNode.insert(m)
				} else {

					tree.RightNode.insert(m)
				}

			}

		}

	} else {
		tree = &Tree{nil, m, nil}
	}
}

//print method for printing a Tree
func print(tree *Tree) {
	if tree != nil {

		fmt.Println(" Value", tree.Value)
		fmt.Printf("Tree Node Left")
		print(tree.LeftNode)
		fmt.Printf("Tree Node Right")
		print(tree.RightNode)
	} else {
		fmt.Printf("Nil\n")
	}
}

// main method
func main() {
	var tree *Tree = &Tree{nil, 1, nil}
	print(tree)
	tree.insert(3)
	print(tree)
	tree.insert(5)
	print(tree)
	tree.LeftNode.insert(7)
	print(tree)

}
