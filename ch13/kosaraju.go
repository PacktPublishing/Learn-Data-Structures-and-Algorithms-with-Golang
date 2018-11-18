///main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing bufio, flag,fmt, os, strconv and strings  package
import (
	"bufio"
	"flag"
	"fmt"
	"os"
	//"sort"
	"strconv"
	"strings"
)

// GraphNode Class
type GraphNode struct {
	ID         int
	visited    bool
	graphEdges []int
	rEdges     []int
}

//DirGraph Class
type DirGraph struct {
	graphNodes map[int]*GraphNode
}

// CreateGraph method
func CreateGraph() *DirGraph {
	var graph DirGraph
	graph.graphNodes = make(map[int]*GraphNode)
	return &graph
}

// CreateGraphNode
func CreateGraphNode() *GraphNode {
	var node GraphNode
	node.ID = -1
	return &node
}

// addGraphEdge method
func (graph *DirGraph) addGraphEdge(tail int, head int) {
	var ok bool
	if _, ok = graph.graphNodes[tail]; !ok {
		panic("No node for edge tail")
	}

	if _, ok = graph.graphNodes[head]; !ok {
		panic("No node for edge head")
	}
	graph.graphNodes[tail].graphEdges = append(graph.graphNodes[tail].graphEdges, head)
	graph.graphNodes[head].rEdges = append(graph.graphNodes[head].rEdges, tail)
}

// addGraphNode method
func (graph *DirGraph) addGraphNode(name int) bool {
	var ok bool
	if _, ok = graph.graphNodes[name]; !ok {
		var node *GraphNode
		node = CreateGraphNode()
		graph.graphNodes[name] = node
		return true
	}
	return false
}

// reset method
func (graph *DirGraph) reset() {
	var node *GraphNode
	for _, node = range graph.graphNodes {
		node.visited = false
	}
}

// show method
func (graph *DirGraph) show() {
	var k int
	var node *GraphNode
	for k, node = range graph.graphNodes {
		fmt.Printf("Node %d (ID#: %d):\nEdges: %v\nBackwards Edges: %v\n\n", k, node.ID, node.graphEdges, node.rEdges)
	}
}

// setFinishingOrder method
func (graph *DirGraph) setFinishingOrder() []*GraphNode {
	graph.reset()
	var nodeArray []*GraphNode
	nodeArray = make([]*GraphNode, 0, len(graph.graphNodes))
	var node *GraphNode
	for _, node = range graph.graphNodes {
		if node.visited == false {
			AssignFinishingLabel(node, graph, &nodeArray)
		}
	}
	return nodeArray
}

//Assign Finishing Label
func AssignFinishingLabel(node *GraphNode, graph *DirGraph, nodeArray *[]*GraphNode) {
	node.visited = true
	var neighborLabel int
	for _, neighborLabel = range node.rEdges {
		if graph.graphNodes[neighborLabel].visited == false {
			AssignFinishingLabel(graph.graphNodes[neighborLabel], graph, nodeArray)
		}
	}
	(*nodeArray) = append(*nodeArray, node)
}

//Mark Strong Connections
func MarkStrongConnections(node *GraphNode, graph *DirGraph, id int) {
	node.visited = true
	node.ID = id
	var neighborLabel int
	for _, neighborLabel = range node.graphEdges {
		if graph.graphNodes[neighborLabel].visited == false {
			MarkStrongConnections(graph.graphNodes[neighborLabel], graph, id)
		}
	}
}

//create Strong Connections
func (graph *DirGraph) createStrongConnections() {
	var l int
	l = 0
	var nodes []*GraphNode
	nodes = graph.setFinishingOrder()
	graph.reset()
	var i int
	for i = len(nodes) - 1; i >= 0; i-- {
		var node *GraphNode
		node = nodes[i]
		if node.visited == false {
			var m int
			m = l
			l++
			MarkStrongConnections(node, graph, m)
		}
	}
}

// Generate Top Five Strong Connections
func (graph *DirGraph) showTopFiveStrongConnections() {
	var aMap map[int]int
	aMap = make(map[int]int)
	var node *GraphNode
	for _, node = range graph.graphNodes {
		if node.ID == -1 {
			panic(fmt.Sprintf("Error - Expected ID to be set (> -1), actual value is %d.", node.ID))
		}
		aMap[node.ID]++
	}

	var id int
  var j int
	j = 0
	for id, _ = range aMap {
		fmt.Print(id, " ")
		j = j +1
		if j > 4 {
			break
		}
	}
	fmt.Println()
}

func main() {
	var graph *DirGraph
	graph = CreateGraph()

	flag.Parse()

	if len(flag.Args()) < 1 {
		panic("Enter the graph edges list file name")
	}
	var file *os.File
	var err error
	file, err = os.Open(flag.Args()[0])
	if err != nil {
		panic(err)
	}
	var scanner *bufio.Scanner
	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		var line string
		line = scanner.Text()
		var fields []string
		fields = strings.Fields(line)
		if len(fields) != 2 {
			panic(fmt.Sprintf("Bad line : %v", fields))
		}
		var node1 int
		var err1 error
		node1, err1 = strconv.Atoi(fields[0])
		if err1 != nil {
			panic(err1)
		}
		var node2 int
		var err2 error
		node2, err2 = strconv.Atoi(fields[1])
		if err2 != nil {
			panic(err2)
		}
		graph.addGraphNode(node1)
		graph.addGraphNode(node2)
		graph.addGraphEdge(node1, node2)
	}
	graph.createStrongConnections()
	graph.show()
	graph.showTopFiveStrongConnections()
}
