///main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing container/heap package

import (
	"container/heap"
	"fmt"
)

//RobotPath interface
type RobotPath interface {
	GetName() string
	GetNeighbors() []RobotPath
	GetNeighborCost(to RobotPath) float64
	GetEstimatedCost(to RobotPath) float64
	GetCurrentNode() *Node
}

//NodePath class
type NodePath struct {
	current Node
}

//Node class
type Node struct {
	pather   RobotPath
	name     string
	cost     float64
	rank     float64
	previous *Node
	open     bool
	closed   bool
	index    int
	next     *Node
}

func NewNode() *Node {

   return &Node{nil,"",1,1,nil,false,false,0,nil}


}
//GetName method
func (n *NodePath) GetName() string {

	return n.current.name
}

//GetCurrentNode method
func (n *NodePath) GetCurrentNode() *Node {

	return &(n.current)
}

//GetNeighbors method
func (n *NodePath) GetNeighbors() []RobotPath {
	var neighbors []RobotPath
	neighbors = []RobotPath{}

	if n.current.previous != nil {

		neighbors = append(neighbors, n.current.previous.pather)
	}

	if n.current.next != nil {

		neighbors = append(neighbors, n.current.next.pather)
	}
	return neighbors
}

//GetNeighborCost method
func (n *NodePath) GetNeighborCost(to RobotPath) float64 {
	return 1.0
}

//GetEstimatedCost method
func (n *NodePath) GetEstimatedCost(to RobotPath) float64 {
	return 1.0
}

//NodeMap class
type NodeMap map[RobotPath]*Node

//getRobotPath method
func (nm NodeMap) GetRobotPath(p RobotPath) *Node {
	var n *Node
	var ok bool
	n, ok = nm[p]
	if !ok {
		if n == nil {
		n = &Node{
			pather: p,
		}

	  }
		n = p.GetCurrentNode()
		nm[p] = n

	}
	return n
}

//GetPath method
func GetPath(from RobotPath, to RobotPath) (path []RobotPath, distance float64, found bool) {
	var nm NodeMap
	nm = NodeMap{}
	var nq *NodeQueue
	nq = &NodeQueue{}
	heap.Init(nq)
	var fromNode *Node
	fromNode = nm.GetRobotPath(from)

	if fromNode != nil {
		fromNode.open = true
		heap.Push(nq, fromNode)
	}

	for {
		if nq.Len() == 0 {
			return
		}
		var current *Node
		current = heap.Pop(nq).(*Node)
		current.open = false
		current.closed = true

		if current == nm.GetRobotPath(to) {
			var p []RobotPath
			p = []RobotPath{}
			var curr *Node
			curr = current
			for curr != nil {
				p = append(p, curr.pather)
				curr = curr.previous
			}
			return p, current.cost, true
		}

		var neighbor RobotPath
		for _, neighbor = range current.pather.GetNeighbors() {
			var cost float64
			cost = current.cost + current.pather.GetNeighborCost(neighbor)
			var neighborNode *Node
			neighborNode = nm.GetRobotPath(neighbor)
			if cost < neighborNode.cost {
				if neighborNode.open {
					heap.Remove(nq, neighborNode.index)
				}
				neighborNode.open = false
				neighborNode.closed = false
			}

			if !neighborNode.open && !neighborNode.closed {
				neighborNode.cost = cost
				neighborNode.open = true
				neighborNode.rank = cost + neighbor.GetEstimatedCost(to)
				neighborNode.previous = current
				heap.Push(nq, neighborNode)
			}
		}
	}
}

// Main method
func main() {
	var start Node
	start = Node{}
	start.name = "start"
	start.cost = 1
	start.index = 0

	var first Node
	first = Node{}
	first.name = "first"
	first.cost = 1
	first.index = 1
	first.previous = &start

	start.next = &first

	var second Node
	second = Node{}
	second.name = "second"
	second.cost = 1
	second.index = 2
	second.previous = &first

	first.next = &second

	var goal Node
	goal = Node{}
	goal.name = "goal"
	goal.cost = 1
	goal.index = 3
	goal.previous = &second

	second.next = &goal

	var startNode NodePath
	startNode = NodePath{}
	start.pather = &startNode
	startNode.current = start

	var firstNode NodePath
	firstNode = NodePath{}
	first.pather = &firstNode
  firstNode.current = first

	var secondNode NodePath
	secondNode = NodePath{}
	second.pather = &secondNode
	secondNode.current = second

	var goalNode NodePath
	goalNode = NodePath{}
	goal.pather = &goalNode
	goalNode.current = goal

	var paths []RobotPath
	var distance float64
	var found bool

	paths, distance, found = GetPath(&startNode, &goalNode)

	//fmt.Println(path, distance, found)
  var path RobotPath
	var i int
	for i, path = range paths {

		fmt.Println(i,path.GetName())
	}

	fmt.Println("distance",distance)

	fmt.Println("found",found)

}
