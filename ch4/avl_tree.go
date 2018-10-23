///main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"encoding/json"
	"fmt"
)

// Key interface
type Key interface {
	LessThan(Key) bool
	EqualTo(Key) bool
}

// TreeNode class
type TreeNode struct {
	KeyValue     Key
	BalanceValue int
	LinkedNodes  [2]*TreeNode
}

//opposite to nodeValue
func opposite(nodeValue int) int {
	return 1 - nodeValue
}

// single rotation
func singleRotation(rootNode *TreeNode, nodeValue int) *TreeNode {

	var saveNode *TreeNode
	saveNode = rootNode.LinkedNodes[opposite(nodeValue)]
	rootNode.LinkedNodes[opposite(nodeValue)] = saveNode.LinkedNodes[nodeValue]
	saveNode.LinkedNodes[nodeValue] = rootNode
	return saveNode
}

// double rotation
func doubleRotation(rootNode *TreeNode, nodeValue int) *TreeNode {

	var saveNode *TreeNode
	saveNode = rootNode.LinkedNodes[opposite(nodeValue)].LinkedNodes[nodeValue]

	rootNode.LinkedNodes[opposite(nodeValue)].LinkedNodes[nodeValue] = saveNode.LinkedNodes[opposite(nodeValue)]
	saveNode.LinkedNodes[opposite(nodeValue)] = rootNode.LinkedNodes[opposite(nodeValue)]
	rootNode.LinkedNodes[opposite(nodeValue)] = saveNode

	saveNode = rootNode.LinkedNodes[opposite(nodeValue)]
	rootNode.LinkedNodes[opposite(nodeValue)] = saveNode.LinkedNodes[nodeValue]
	saveNode.LinkedNodes[nodeValue] = rootNode
	return saveNode
}

// adjust balance factors after double rotation
func adjustBalance(rootNode *TreeNode, nodeValue int, balanceValue int) {

	var node *TreeNode
	node = rootNode.LinkedNodes[nodeValue]
	var oppNode *TreeNode
	oppNode = node.LinkedNodes[opposite(balanceValue)]
	switch oppNode.BalanceValue {
	case 0:
		rootNode.BalanceValue = 0
		node.BalanceValue = 0
	case balanceValue:
		rootNode.BalanceValue = -balanceValue
		node.BalanceValue = 0
	default:
		rootNode.BalanceValue = 0
		node.BalanceValue = balanceValue
	}
	oppNode.BalanceValue = 0
}

// balance factor of the tree is changed by single and double rotation
func BalanceTree(rootNode *TreeNode, nodeValue int) *TreeNode {
	var node *TreeNode
	node = rootNode.LinkedNodes[nodeValue]
	var balance int
	balance = 2*nodeValue - 1
	if node.BalanceValue == balance {
		rootNode.BalanceValue = 0
		node.BalanceValue = 0
		return singleRotation(rootNode, opposite(nodeValue))
	}
	adjustBalance(rootNode, nodeValue, balance)
	return doubleRotation(rootNode, opposite(nodeValue))
}

//inserts RootNode with key value
func insertRNode(rootNode *TreeNode, key Key) (*TreeNode, bool) {
	if rootNode == nil {
		return &TreeNode{KeyValue: key}, false
	}
	var dir int
	dir = 0
	if rootNode.KeyValue.LessThan(key) {
		dir = 1
	}
	var done bool
	rootNode.LinkedNodes[dir], done = insertRNode(rootNode.LinkedNodes[dir], key)
	if done {
		return rootNode, true
	}
	rootNode.BalanceValue = rootNode.BalanceValue + (2*dir - 1)
	switch rootNode.BalanceValue {
	case 0:
		return rootNode, true
	case 1, -1:
		return rootNode, false
	}
	return BalanceTree(rootNode, dir), true
}

// Insert a node into the AVL tree.
func InsertNode(treeNode **TreeNode, key Key) {
	*treeNode, _ = insertRNode(*treeNode, key)
}

// RemoveNode removes an element from an AVL tree.
func RemoveNode(treeNode **TreeNode, key Key) {
	*treeNode, _ = removeRNode(*treeNode, key)
}

func removeBalance(rootNode *TreeNode, nodeValue int) (*TreeNode, bool) {
	var node *TreeNode
	node = rootNode.LinkedNodes[opposite(nodeValue)]
	var balance int
	balance = 2*nodeValue - 1
	switch node.BalanceValue {
	case -balance:
		rootNode.BalanceValue = 0
		node.BalanceValue = 0
		return singleRotation(rootNode, nodeValue), false
	case balance:
		adjustBalance(rootNode, opposite(nodeValue), -balance)
		return doubleRotation(rootNode, nodeValue), false
	}
	rootNode.BalanceValue = -balance
	node.BalanceValue = balance
	return singleRotation(rootNode, nodeValue), true
}

func removeRNode(rootNode *TreeNode, key Key) (*TreeNode, bool) {
	if rootNode == nil {
		return nil, false
	}
	if rootNode.KeyValue.EqualTo(key) {
		switch {
		case rootNode.LinkedNodes[0] == nil:
			return rootNode.LinkedNodes[1], false
		case rootNode.LinkedNodes[1] == nil:
			return rootNode.LinkedNodes[0], false
		}
		var heirNode *TreeNode
		heirNode = rootNode.LinkedNodes[0]
		for heirNode.LinkedNodes[1] != nil {
			heirNode = heirNode.LinkedNodes[1]
		}
		rootNode.KeyValue = heirNode.KeyValue
		key = heirNode.KeyValue
	}
	var dir int
	dir = 0
	if rootNode.KeyValue.LessThan(key) {
		dir = 1
	}
	var done bool
	rootNode.LinkedNodes[dir], done = removeRNode(rootNode.LinkedNodes[dir], key)
	if done {
		return rootNode, true
	}
	rootNode.BalanceValue = rootNode.BalanceValue + (1 - 2*dir)
	switch rootNode.BalanceValue {
	case 1, -1:
		return rootNode, true
	case 0:
		return rootNode, false
	}
	return removeBalance(rootNode, dir)
}

type integerKey int

func (k integerKey) LessThan(k1 Key) bool { return k < k1.(integerKey) }
func (k integerKey) EqualTo(k1 Key) bool  { return k == k1.(integerKey) }

//main method
func main() {
	var tree *TreeNode
	fmt.Println("Empty Tree:")
	var avl []byte
	avl, _ = json.MarshalIndent(tree, "", "   ")
	fmt.Println(string(avl))

	fmt.Println("\nInsert Tree:")
	InsertNode(&tree, integerKey(5))
	InsertNode(&tree, integerKey(3))
	InsertNode(&tree, integerKey(8))
	InsertNode(&tree, integerKey(7))
	InsertNode(&tree, integerKey(6))
	InsertNode(&tree, integerKey(10))
	avl, _ = json.MarshalIndent(tree, "", "   ")
	fmt.Println(string(avl))

	fmt.Println("\nRemove Tree:")
	RemoveNode(&tree, integerKey(3))
	RemoveNode(&tree, integerKey(7))
	avl, _ = json.MarshalIndent(tree, "", "   ")
	fmt.Println(string(avl))
}
