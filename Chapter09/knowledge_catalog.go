///main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
)

// Class Type
type Class struct {
	Name string
}

// Knowledge Graph type
type KnowledgeGraph struct {
	GraphNodes map[Class]struct{}
	Links      map[Class]map[Class]struct{}
}

// NewKnowledgeGraph method
func NewKnowledgeGraph() *KnowledgeGraph {
	return &KnowledgeGraph{
		GraphNodes: make(map[Class]struct{}),
		Links:      make(map[Class]map[Class]struct{}),
	}
}

// AddClass method
func (knowledgeGraph *KnowledgeGraph) AddClass(class Class) bool {

	var exists bool
	if _, exists = knowledgeGraph.GraphNodes[class]; exists {
		return true
	}
	knowledgeGraph.GraphNodes[class] = struct{}{}
	return true
}

// Add Link
func (knowledgeGraph *KnowledgeGraph) AddLink(class1 Class, class2 Class) {
	var exists bool
	if _, exists = knowledgeGraph.GraphNodes[class1]; !exists {
		knowledgeGraph.AddClass(class1)
	}
	if _, exists = knowledgeGraph.GraphNodes[class2]; !exists {
		knowledgeGraph.AddClass(class2)
	}

	if _, exists = knowledgeGraph.Links[class1]; !exists {
		knowledgeGraph.Links[class1] = make(map[Class]struct{})
	}
	knowledgeGraph.Links[class1][class2] = struct{}{}

}

// Print Links method
func (knowledgeGraph *KnowledgeGraph) PrintLinks() {
	var car Class
	car = Class{"Car"}

	fmt.Printf("Printing all links adjacent to %s\n", car.Name)

	var node Class
	for node = range knowledgeGraph.Links[car] {
		fmt.Printf("Link: %s -> %s\n", car.Name, node.Name)
	}

	var m map[Class]struct{}
	fmt.Println("Printing all links.")
	for car, m = range knowledgeGraph.Links {
		var vertex Class
		for vertex = range m {
			fmt.Printf("Link: %s -> %s\n", car.Name, vertex.Name)
		}
	}
}

// main method
func main() {

	var knowledgeGraph *KnowledgeGraph

	knowledgeGraph = NewKnowledgeGraph()

	var car = Class{"Car"}
	var tyre = Class{"Tyre"}
	var door = Class{"Door"}
	var hood = Class{"Hood"}

	knowledgeGraph.AddClass(car)
	knowledgeGraph.AddClass(tyre)
	knowledgeGraph.AddClass(door)
	knowledgeGraph.AddClass(hood)

	knowledgeGraph.AddLink(car, tyre)
	knowledgeGraph.AddLink(car, door)
	knowledgeGraph.AddLink(car, hood)

	var tube = Class{"Tube"}
	var axle = Class{"Axle"}
	var handle = Class{"Handle"}
	var windowGlass = Class{"Window Glass"}

	knowledgeGraph.AddClass(tube)
	knowledgeGraph.AddClass(axle)
	knowledgeGraph.AddClass(handle)
	knowledgeGraph.AddClass(windowGlass)

	knowledgeGraph.AddLink(tyre, tube)
	knowledgeGraph.AddLink(tyre, axle)
	knowledgeGraph.AddLink(door, handle)
	knowledgeGraph.AddLink(door, windowGlass)

	knowledgeGraph.PrintLinks()
}
