///main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt and sync package
import (
	"fmt"
	"sync"
)

// TreeNode class
type TreeNode struct {
	key       int
	value     int
	leftNode  *TreeNode //left
	rightNode *TreeNode //right
}

// BinarySearchTree class
type BinarySearchTree struct {
	rootNode *TreeNode
	lock     sync.RWMutex
}

// InsertElement method
func (tree *BinarySearchTree) InsertElement(key int, value int) {
	tree.lock.Lock()
	defer tree.lock.Unlock()

	var treeNode *TreeNode
	treeNode = &TreeNode{key, value, nil, nil}
	if tree.rootNode == nil {
		tree.rootNode = treeNode
	} else {
		insertTreeNode(tree.rootNode, treeNode)
	}
}

//  insertTreeNode method
func insertTreeNode(rootNode *TreeNode, newTreeNode *TreeNode) {
	if newTreeNode.key < rootNode.key {
		if rootNode.leftNode == nil {
			rootNode.leftNode = newTreeNode
		} else {
			insertTreeNode(rootNode.leftNode, newTreeNode)
		}
	} else {
		if rootNode.rightNode == nil {
			rootNode.rightNode = newTreeNode
		} else {
			insertTreeNode(rootNode.rightNode, newTreeNode)
		}
	}
}

// InOrderTraverseTree method
func (tree *BinarySearchTree) InOrderTraverseTree(function func(int)) {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	inOrderTraverseTree(tree.rootNode, function)
}

//  inOrderTraverseTree method
func inOrderTraverseTree(treeNode *TreeNode, function func(int)) {
	if treeNode != nil {
		inOrderTraverseTree(treeNode.leftNode, function)
		function(treeNode.value)
		inOrderTraverseTree(treeNode.rightNode, function)
	}
}

// PreOrderTraverse method
func (tree *BinarySearchTree) PreOrderTraverseTree(function func(int)) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	preOrderTraverseTree(tree.rootNode, function)
}

//  preOrderTraverseTree method
func preOrderTraverseTree(treeNode *TreeNode, function func(int)) {
	if treeNode != nil {
		function(treeNode.value)
		preOrderTraverseTree(treeNode.leftNode, function)
		preOrderTraverseTree(treeNode.rightNode, function)
	}
}

// PostOrderTraverseTree method
func (tree *BinarySearchTree) PostOrderTraverseTree(function func(int)) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	postOrderTraverseTree(tree.rootNode, function)
}

// postOrderTraverseTree method
func postOrderTraverseTree(treeNode *TreeNode, function func(int)) {
	if treeNode != nil {
		postOrderTraverseTree(treeNode.leftNode, function)
		postOrderTraverseTree(treeNode.rightNode, function)
		function(treeNode.value)
	}
}

// MinNode method
func (tree *BinarySearchTree) MinNode() *int {
	tree.lock.RLock()
	defer tree.lock.RUnlock()

	var treeNode *TreeNode
	treeNode = tree.rootNode
	if treeNode == nil {
		return (*int)(nil)
	}
	for {
		if treeNode.leftNode == nil {
			return &treeNode.value
		}
		treeNode = treeNode.leftNode
	}
}

// MaxNode method
func (tree *BinarySearchTree) MaxNode() *int {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	var treeNode *TreeNode
	treeNode = tree.rootNode
	if treeNode == nil {
		return (*int)(nil)
	}
	for {
		if treeNode.rightNode == nil {
			return &treeNode.value
		}
		treeNode = treeNode.rightNode
	}
}

// SearchNode method
func (tree *BinarySearchTree) SearchNode(key int) bool {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	return searchNode(tree.rootNode, key)
}

//  searchNode method
func searchNode(treeNode *TreeNode, key int) bool {
	if treeNode == nil {
		return false
	}
	if key < treeNode.key {
		return searchNode(treeNode.leftNode, key)
	}
	if key > treeNode.key {
		return searchNode(treeNode.rightNode, key)
	}
	return true
}

// RemoveNode method
func (tree *BinarySearchTree) RemoveNode(key int) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	removeNode(tree.rootNode, key)
}

//  removeNode method
func removeNode(treeNode *TreeNode, key int) *TreeNode {
	if treeNode == nil {
		return nil
	}
	if key < treeNode.key {
		treeNode.leftNode = removeNode(treeNode.leftNode, key)
		return treeNode
	}
	if key > treeNode.key {
		treeNode.rightNode = removeNode(treeNode.rightNode, key)
		return treeNode
	}
	if treeNode.leftNode == nil && treeNode.rightNode == nil {
		treeNode = nil
		return nil
	}
	if treeNode.leftNode == nil {
		treeNode = treeNode.rightNode
		return treeNode
	}
	if treeNode.rightNode == nil {
		treeNode = treeNode.leftNode
		return treeNode
	}
	var leftmostrightNode *TreeNode
	leftmostrightNode = treeNode.rightNode
	for {
		if leftmostrightNode != nil && leftmostrightNode.leftNode != nil {
			leftmostrightNode = leftmostrightNode.leftNode
		} else {
			break
		}
	}
	treeNode.key, treeNode.value = leftmostrightNode.key, leftmostrightNode.value
	treeNode.rightNode = removeNode(treeNode.rightNode, treeNode.key)
	return treeNode
}

// String method
func (tree *BinarySearchTree) String() {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	fmt.Println("************************************************")
	stringify(tree.rootNode, 0)
	fmt.Println("************************************************")
}

// stringify method
func stringify(treeNode *TreeNode, level int) {
	if treeNode != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "***> "
		level++
		stringify(treeNode.leftNode, level)
		fmt.Printf(format+"%d\n", treeNode.key)
		stringify(treeNode.rightNode, level)
	}
}

// prints the binary search tree
func print(tree *BinarySearchTree) {
	if tree != nil {

		fmt.Println(" Value", tree.rootNode.value)
		fmt.Printf("Root Tree Node")
		printTreeNode(tree.rootNode)
	} else {
		fmt.Printf("Nil\n")
	}
}

//printTreeNode method
func printTreeNode(treeNode *TreeNode) {
	if treeNode != nil {
		fmt.Println(" Value", treeNode.value)
		fmt.Printf("TreeNode Left")
		printTreeNode(treeNode.leftNode)
		fmt.Printf("TreeNode Right")
		printTreeNode(treeNode.rightNode)
	} else {
		fmt.Printf("Nil\n")
	}

}
