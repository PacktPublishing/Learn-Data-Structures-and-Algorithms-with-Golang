///main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing json, fmt, os, ioutil and sort  package

import (
	"encoding/json"
	"fmt"
	//"io/ioutil"
	"os"
	"sort"
)

//GetOptimalPlan method
func GetOptimalPlan(nodes []GraphNode, graphEdges DeliveryPersonRouteGraph) DeliveryPersonRouteGraph {
	var graph DeliveryPersonRouteGraph
	var graph_new DeliveryPersonRouteGraph
	sort.Sort(graphEdges)

	var edge GraphEdge
	for _, edge = range graphEdges {
		graph_new = graph

		graph_new = append(graph_new, edge)

		var out bool
		out = false
		var currentGraphEdge GraphEdge
		for _, currentGraphEdge = range graph {
			if currentGraphEdge.Current == edge.Current {
				out = true
				break
			}
			if currentGraphEdge.Neighbor == edge.Neighbor {
				out = true
				break
			}
		}
		if out == true {
			continue
		}

		currentGraphEdge = edge
		var visitedNodesMap map[*GraphNode]bool
		visitedNodesMap = make(map[*GraphNode]bool)

		var routingLoop bool
		routingLoop = false
		for {
			if visitedNodesMap[currentGraphEdge.Neighbor] {
				routingLoop = true
				break
			}
			visitedNodesMap[currentGraphEdge.Current] = true

			var found bool
			found = false
			var item GraphEdge
			for _, item = range graph_new {
				if item.Current == currentGraphEdge.Neighbor {
					currentGraphEdge = item
					found = true
					break
				}
			}
			if !found {
				break
			}
		}
		if routingLoop {
			continue
		}
		graph = graph_new
	}
	return graph
}

//printDeliveryPersonRouteGraph method
func printDeliveryPersonRouteGraph(graph DeliveryPersonRouteGraph) {

	var newGraph DeliveryPersonRouteGraph
	var startId int
	startId = -1
	for {
		var done bool
		done = true
		var edge GraphEdge
		for _, edge = range graph {
			if (*edge.Current).Identifier != startId {
				continue
			}
			done = false
			newGraph = append(newGraph, edge)
			startId = (*edge.Neighbor).Identifier
		}
		if done {
			break
		}
	}
	var length float64
	length = float64(0)
	var edge GraphEdge
	for _, edge = range newGraph {
		fmt.Printf("%d -> %d\n", (*edge.Current).Identifier, (*edge.Neighbor).Identifier)
		length += edge.EdgeLength
	}
	fmt.Printf("%d edges, length %#v\n", len(newGraph), length)

	var sfile *os.File
	var err error
	sfile, err = os.Create("delivery_plan.json")
	if err != nil {
		panic(err)
	}
	var text []byte
	text, err = json.Marshal(newGraph)
	if err != nil {
		panic(err)
	}
	_, err = sfile.Write(text)
	if err != nil {
		panic(err)
	}
}

// evaluteDeliveryPersonRouteGraph method
func evaluateDeliveryPersonRouteGraph(graph DeliveryPersonRouteGraph, nodes []GraphNode) bool {
	var edge GraphEdge
	for _, edge = range graph {
		if (*edge.Neighbor).Identifier == -1 {
			panic("Edge pointing to -1!")
		}
	}
	if len(graph) < len(nodes)-1 {
		fmt.Printf("Output graph is shorter than expected: got %d, want %d\n", len(graph), len(nodes)-1)
		return false
	} else {
		fmt.Print("Output graph length is ok.\n")
		return true
	}
}
