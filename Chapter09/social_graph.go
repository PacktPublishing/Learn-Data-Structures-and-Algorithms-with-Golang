///main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
)

// Social Graph
type SocialGraph struct {
	Size  int
	Links [][]Link
}

// Link class
type Link struct {
	Vertex1    int
	Vertex2    int
	LinkWeight int
}

// NewSocialGraph method
func NewSocialGraph(num int) *SocialGraph {
	return &SocialGraph{
		Size:  num,
		Links: make([][]Link, num),
	}
}

// AddLink method
func (socialGraph *SocialGraph) AddLink(vertex1 int, vertex2 int, weight int) {
	socialGraph.Links[vertex1] = append(socialGraph.Links[vertex1], Link{Vertex1: vertex1, Vertex2: vertex2, LinkWeight: weight})
}

// Print Links Example
func (socialGraph *SocialGraph) PrintLinks() {

	var vertex int
	vertex = 0

	fmt.Printf("Printing all links from %d\n", vertex)
	var link Link
	for _, link = range socialGraph.Links[vertex] {
		fmt.Printf("Link: %d -> %d (%d)\n", link.Vertex1, link.Vertex2, link.LinkWeight)
	}

	fmt.Println("Printing all links in graph.")
	var adjacent []Link
	for _, adjacent = range socialGraph.Links {
		for _, link = range adjacent {
			fmt.Printf("Link: %d -> %d (%d)\n", link.Vertex1, link.Vertex2, link.LinkWeight)
		}
	}
}

// main method
/*
func main() {

	var socialGraph *SocialGraph

	socialGraph = NewSocialGraph(4)

	socialGraph.AddLink(0, 1, 1)
	socialGraph.AddLink(0, 2, 1)
	socialGraph.AddLink(1, 3, 1)
	socialGraph.AddLink(2, 4, 1)

	socialGraph.PrintLinks()

}
*/
