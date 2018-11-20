///main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing heap and fmt packages

import (
	"container/heap"
	"fmt"
)

// VertexQueue class
type VertexQueue struct {
	vertices  []GraphVertex
	vertexMap map[GraphVertex]int
	prMap     map[GraphVertex]int
}

func (queue *VertexQueue) Len() int { return len(queue.vertices) }
func (queue *VertexQueue) Less(i int, j int) bool {
	return queue.prMap[queue.vertices[i]] < queue.prMap[queue.vertices[j]]
}
func (queue *VertexQueue) Swap(i int, j int) {
	queue.vertices[i], queue.vertices[j] = queue.vertices[j], queue.vertices[i]
	queue.vertexMap[queue.vertices[i]] = i
	queue.vertexMap[queue.vertices[j]] = j
}

// Push method
func (queue *VertexQueue) Push(x interface{}) {
	var n int
	n = len(queue.vertices)
	var vertex GraphVertex
	vertex = x.(GraphVertex)
	queue.vertexMap[vertex] = n
	queue.vertices = append(queue.vertices, vertex)
}

// Pop method
func (queue *VertexQueue) Pop() interface{} {
	var old []GraphVertex
	old = queue.vertices
	var n int
	n = len(old)
	var vertex GraphVertex
	vertex = old[n-1]
	queue.vertexMap[vertex] = -1
	queue.vertices = old[0 : n-1]
	return vertex
}

// update method
func (queue *VertexQueue) update(vertex GraphVertex, priority int) {
	queue.prMap[vertex] = priority
	heap.Fix(queue, queue.vertexMap[vertex])
}

// addWithPriority method
func (queue *VertexQueue) addWithPriority(vertex GraphVertex, priority int) {
	heap.Push(queue, vertex)
	queue.update(vertex, priority)
}

const (
	Infinity      = int(^uint(0) >> 1)
	Uninitialized = -1
)

// DirGraph Interface
type DirGraph interface {
	GraphVertices() []GraphVertex
	NodeNeighbors(vertex GraphVertex) []GraphVertex
	Weight(p GraphVertex, q GraphVertex) int
}

type GraphVertex int

// Graph class
type Graph struct {
	identifiers map[string]GraphVertex
	vertexNames map[GraphVertex]string
	graphEdges  map[GraphVertex]map[GraphVertex]int
}

// CreateGraph Method
func CreateGraph(vertIdentifiers map[string]GraphVertex) Graph {
	var graph Graph
	graph = Graph{identifiers: vertIdentifiers}
	graph.vertexNames = make(map[GraphVertex]string)
	var k string
	var v GraphVertex
	for k, v = range vertIdentifiers {
		graph.vertexNames[v] = k
	}
	graph.graphEdges = make(map[GraphVertex]map[GraphVertex]int)
	return graph
}

//SetEdge method
func (graph Graph) SetEdge(u string, v string, w int) {
	var ok bool
	if _, ok = graph.graphEdges[graph.identifiers[u]]; !ok {
		graph.graphEdges[graph.identifiers[u]] = make(map[GraphVertex]int)
	}
	graph.graphEdges[graph.identifiers[u]][graph.identifiers[v]] = w
}

//SetPath method
func (graph Graph) SetPath(vertex GraphVertex, prev map[GraphVertex]GraphVertex) (str string) {
	str = graph.vertexNames[vertex]
	for prev[vertex] >= 0 {
		vertex = prev[vertex]
		str = graph.vertexNames[vertex] + str
	}
	return str
}

// GraphVertices method
func (graph Graph) GraphVertices() (vertices []GraphVertex) {
	var vertex GraphVertex
	for _, vertex = range graph.identifiers {
		vertices = append(vertices, vertex)
	}
	return vertices
}

//NodeNeighbors method
func (graph Graph) NodeNeighbors(u GraphVertex) (vertices []GraphVertex) {
	var vertex GraphVertex
	for vertex = range graph.graphEdges[u] {
		vertices = append(vertices, vertex)
	}
	return vertices
}

//GetWeight method
func (graph Graph) Weight(u GraphVertex, v GraphVertex) int { return graph.graphEdges[u][v] }

//Dijkstra Algo method
func DijkstraAlgo(graph Graph, source GraphVertex) (dist map[GraphVertex]int, prev map[GraphVertex]GraphVertex) {
	dist = make(map[GraphVertex]int)
	prev = make(map[GraphVertex]GraphVertex)
	var sourceVertex GraphVertex
	sourceVertex = source
	dist[sourceVertex] = 0
	var queue *VertexQueue
	queue = &VertexQueue{[]GraphVertex{}, make(map[GraphVertex]int), make(map[GraphVertex]int)}
	var vertex GraphVertex
	for _, vertex = range graph.GraphVertices() {
		if vertex != sourceVertex {
			dist[vertex] = Infinity
		}
		prev[vertex] = Uninitialized
		queue.addWithPriority(vertex, dist[vertex])
	}
	for len(queue.vertices) != 0 {
		var u GraphVertex
		u = heap.Pop(queue).(GraphVertex)
		var v GraphVertex
		for _, v = range graph.NodeNeighbors(u) {
			var alt int
			alt = dist[u] + graph.Weight(u, v)
			if alt < dist[v] {
				dist[v] = alt
				prev[v] = u
				queue.update(v, alt)
			}
		}
	}
	return dist, prev
}

// main method
func main() {
	var graph Graph
	graph = CreateGraph(map[string]GraphVertex{
		"j": 11,
		"k": 21,
		"l": 31,
		"m": 41,
		"n": 51,
		"o": 61,
	})
	graph.SetEdge("j", "k", 71)
	graph.SetEdge("j", "l", 91)
	graph.SetEdge("j", "o", 114)
	graph.SetEdge("k", "l", 110)
	graph.SetEdge("k", "m", 115)
	graph.SetEdge("l", "m", 111)
	graph.SetEdge("l", "o", 12)
	graph.SetEdge("m", "n", 16)
	graph.SetEdge("n", "o", 19)

	var dist map[GraphVertex]int
	var prev map[GraphVertex]GraphVertex

	dist, prev = DijkstraAlgo(graph, graph.identifiers["k"])
	fmt.Printf("Distance  %s: %d, Path: %s\n", "n", dist[graph.identifiers["n"]], graph.SetPath(graph.identifiers["n"], prev))
	fmt.Printf("Distance  %s: %d, Path: %s\n", "m", dist[graph.identifiers["m"]], graph.SetPath(graph.identifiers["m"], prev))
}
