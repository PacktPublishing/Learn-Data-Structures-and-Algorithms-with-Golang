///main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing testing package
import (
	"testing"
)

// NewDeliveryPersonGraph test method
func TestNewDeliveryPersonGraph(test *testing.T) {

	var routeGraph DeliveryPersonRouteGraph

	routeGraph = NewDeliveryPersonRouteGraph()

	routeGraph = append(routeGraph,GraphEdge{})

	test.Log(routeGraph)

	if routeGraph == nil {

		test.Errorf("error in creating a routeGraph")
	}

}

// NewRouteMap test method
func TestNewRouteMap(test *testing.T) {

	var routeMap *RouteMap

	routeMap = NewRouteMap()


	test.Log(routeMap)

	if routeMap == nil {

		test.Errorf("error in creating a routeMap")
	}

}

// NewNode test method
func TestNewNode(test *testing.T) {

	var node *Node

	node = NewNode()

	test.Log(node)

	if node == nil {

		test.Errorf("error in creating a node")
	}

}
