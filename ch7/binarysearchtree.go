///main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
  "sync"
)


// TreeNode a single node that composes the tree
type TreeNode struct {
    key   int
    value int
    leftNode  *TreeNode //left
    rightNode *TreeNode //right
}

// BinarySearchTree the binary search tree
type BinarySearchTree struct {
    rootNode *TreeNode
    lock sync.RWMutex
}

// InsertElement inserts the element with key and value in a binary search  tree
func (tree *BinarySearchTree) InsertElement(key int, value int) {
    tree.lock.Lock()
    defer tree.lock.Unlock()

    var treeNode *TreeNode
     treeNode= &TreeNode{key, value, nil, nil}
    if tree.rootNode == nil {
        tree.rootNode = treeNode
    } else {
        insertTreeNode(tree.rootNode, treeNode)
    }
}

// internal function insertTreeNode to find the right place for a tree node in a binary search tree
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

// InOrderTraverseTree visits all Tree nodes with in-order traversing
func (tree *BinarySearchTree) InOrderTraverseTree(function func(int)) {
    tree.lock.RLock()
    defer tree.lock.RUnlock()
    inOrderTraverseTree(tree.rootNode, function)
}

// internal recursive function inOrderTraverseTree to traverse in order
func inOrderTraverseTree(treeNode *TreeNode, function func(int)) {
    if treeNode != nil {
        inOrderTraverseTree(treeNode.leftNode, function)
        function(treeNode.value)
        inOrderTraverseTree(treeNode.rightNode, function)
    }
}

// PreOrderTraverse visits all Tree Nodes with pre-order traversing
func (tree *BinarySearchTree) PreOrderTraverseTree(function func(int)) {
    tree.lock.Lock()
    defer tree.lock.Unlock()
    preOrderTraverseTree(tree.rootNode, function)
}

// internal recursive function preOrderTraverseTree to traverse pre order
func preOrderTraverseTree(treeNode *TreeNode, function func(int)) {
    if treeNode != nil {
        function(treeNode.value)
        preOrderTraverseTree(treeNode.leftNode, function)
        preOrderTraverseTree(treeNode.rightNode, function)
    }
}

// PostOrderTraverseTree visits all nodes with post-order traversing
func (tree *BinarySearchTree) PostOrderTraverseTree(function func(int)) {
    tree.lock.Lock()
    defer tree.lock.Unlock()
    postOrderTraverseTree(tree.rootNode, function)
}

// internal recursive function postOrderTraverseTree to traverse post order
func postOrderTraverseTree(treeNode *TreeNode, function func(int)) {
    if treeNode != nil {
        postOrderTraverseTree(treeNode.leftNode, function)
        postOrderTraverseTree(treeNode.rightNode, function)
        function(treeNode.value)
    }
}

// MinNode returns the vale of the Node  with min value stored in the tree
func (tree *BinarySearchTree) MinNode() *int {
    tree.lock.RLock()
    defer tree.lock.RUnlock()

    var treeNode *TreeNode
    treeNode = tree.rootNode
    if treeNode == nil {
			  //nil instead of 0
        return (*int)(nil)
    }
    for {
        if treeNode.leftNode == nil {
            return &treeNode.value
        }
        treeNode = treeNode.leftNode
    }
}

// MaxNode returns the value of the Node with max value stored in the tree
func (tree *BinarySearchTree) MaxNode() *int {
    tree.lock.RLock()
    defer tree.lock.RUnlock()
    var treeNode *TreeNode
    treeNode = tree.rootNode
    if treeNode == nil {
			  //nil instead of 0
        return (*int)(nil)
    }
    for {
        if treeNode.rightNode == nil {
            return &treeNode.value
        }
        treeNode = treeNode.rightNode
    }
}

// SearchNode returns true if the key exists in the tree
func (tree *BinarySearchTree) SearchNode(key int) bool {
    tree.lock.RLock()
    defer tree.lock.RUnlock()
    return searchNode(tree.rootNode, key)
}

// internal recursive function searchNode to search a Node with key in the tree
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

// RemoveNode removes the Item with key `key` from the tree
func (tree *BinarySearchTree) RemoveNode(key int) {
    tree.lock.Lock()
    defer tree.lock.Unlock()
    removeNode(tree.rootNode, key)
}

// internal recursive function removeNode to remove an element with key
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
    // key == node.key
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
        //find smallest value on the right side
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

// String prints a visual representation of the tree
func (tree *BinarySearchTree) String() {
    tree.lock.Lock()
    defer tree.lock.Unlock()
    fmt.Println("************************************************")
    stringify(tree.rootNode, 0)
    fmt.Println("************************************************")
}

// internal recursive function to print a tree
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

    fmt.Println(" Value",tree.rootNode.value)
    fmt.Printf("Root Tree Node")
    printTreeNode(tree.rootNode)
    //fmt.Printf("Tree Node Right")
    //print(tree.rootNode.rightNode)
  } else {
    fmt.Printf("Nil\n")
  }
}
//prints the treeNode
func printTreeNode(treeNode *TreeNode){
    if(treeNode != nil) {
			fmt.Println(" Value",treeNode.value)
			fmt.Printf("TreeNode Left")
			printTreeNode(treeNode.leftNode)
			fmt.Printf("TreeNode Right")
			printTreeNode(treeNode.rightNode)
		} else {
			fmt.Printf("Nil\n")
		}

}
