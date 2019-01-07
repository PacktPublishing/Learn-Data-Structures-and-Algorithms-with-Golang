///main package has examples shown
// in Go Data Structures and algorithms book
package main


//NodeQueue type
type NodeQueue []*Node

//Len method
func (nq NodeQueue) Len() int {
	return len(nq)
}
//Less method
func (nq NodeQueue) Less(i, j int) bool {
	return nq[i].rank < nq[j].rank
}
//Swap method
func (nq NodeQueue) Swap(i, j int) {
	nq[i], nq[j] = nq[j], nq[i]
	nq[i].index = i
	nq[j].index = j
}
//Push method
func (nq *NodeQueue) Push(x interface{}) {
	var n int
	n = len(*nq)
	var no *Node
	no = x.(*Node)
	no.index = n
	*nq = append(*nq, no)
}
//Pop method
func (nq *NodeQueue) Pop() interface{} {
	var old NodeQueue
	old = *nq
	var n int
	n = len(old)
	var no *Node
	no = old[n-1]
	no.index = -1
	*nq = old[0 : n-1]
	return no
}
