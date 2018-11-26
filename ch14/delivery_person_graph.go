///main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing math,json, rand and os packages

import (
	"math"
)

type Label int
type Category int

// GraphNode class
type GraphNode struct {
	Identifier int
	X          int
	Y          int
	Label
	Category
}

// GraphEdge class
type GraphEdge struct {
	Current    *GraphNode
	Neighbor   *GraphNode
	EdgeLength float64
}

type DeliveryPersonRouteGraph []GraphEdge

// Len method
func (graph DeliveryPersonRouteGraph) Len() int {
	return len(graph)
}

// Swap method
func (graph DeliveryPersonRouteGraph) Swap(i int, j int) {
	graph[i], graph[j] = graph[j], graph[i]
}

// Less method
func (graph DeliveryPersonRouteGraph) Less(i, j int) bool {
	return graph[i].EdgeLength < graph[j].EdgeLength
}

// GetDistance method
func GetDistance(u GraphNode, v GraphNode) float64 {
	return math.Sqrt(math.Pow(float64(u.X-v.X), 2) + math.Pow(float64(u.Y-v.Y), 2))
}
