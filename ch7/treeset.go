///main package has examples shown
// in Go Data Structures and algorithms book
package main

// TreeSet type
type TreeSet struct {
	bst *BinarySearchTree
}

// InsertTreeNode method
func (treeset *TreeSet) InsertTreeNode(treeNodes ...TreeNode) {
	var treeNode TreeNode
	for _, treeNode = range treeNodes {
		treeset.bst.InsertElement(treeNode.key, treeNode.value)
	}
}

// Delete method
func (treeset *TreeSet) Delete(treeNodes ...TreeNode) {
	var treeNode TreeNode
	for _, treeNode = range treeNodes {
		treeset.bst.RemoveNode(treeNode.key)
	}
}

// Search method
func (treeset *TreeSet) Search(treeNodes ...TreeNode) bool {
	var treeNode TreeNode
	var exists bool
	for _, treeNode = range treeNodes {
		if exists = treeset.bst.SearchNode(treeNode.key); !exists {
			return false
		}
	}
	return true
}

// String method
func (treeset *TreeSet) String() {

	treeset.bst.String()

}

// main method
func main() {
	var treeset *TreeSet = &TreeSet{}

	treeset.bst = &BinarySearchTree{}

	var node1 TreeNode = TreeNode{8, 8, nil, nil}
	var node2 TreeNode = TreeNode{3, 3, nil, nil}
	var node3 TreeNode = TreeNode{10, 10, nil, nil}
	var node4 TreeNode = TreeNode{1, 1, nil, nil}
	var node5 TreeNode = TreeNode{6, 6, nil, nil}

	treeset.InsertTreeNode(node1, node2, node3, node4, node5)

	treeset.String()

}
