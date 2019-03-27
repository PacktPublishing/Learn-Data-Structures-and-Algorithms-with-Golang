///main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
)

// Name type
type Name string
// Social Graph class
type SocialGraph struct {
	GraphNodes map[Name]struct{}
	Links map[Name]map[Name]struct{}
}

// NewSocialGraph method
func NewSocialGraph() *SocialGraph {
	return &SocialGraph{
		GraphNodes: make(map[Name]struct{}),
		Links: make(map[Name]map[Name]struct{}),
	}
}

// AddEntity method
func (socialGraph *SocialGraph) AddEntity(name Name) bool {

	var exists bool
	if _, exists = socialGraph.GraphNodes[name]; exists {
		return true
	}
	socialGraph.GraphNodes[name] = struct{}{}
	return true
}

// Add Link
func (socialGraph *SocialGraph) AddLink(name1 Name, name2 Name) {
	var exists bool
	if _, exists = socialGraph.GraphNodes[name1]; !exists {
		socialGraph.AddEntity(name1)
	}
	if _, exists = socialGraph.GraphNodes[name2]; !exists {
		socialGraph.AddEntity(name2)
	}

	if _, exists = socialGraph.Links[name1]; !exists {
		socialGraph.Links[name1] = make(map[Name]struct{})
	}
	socialGraph.Links[name1][name2] = struct{}{}

}

func (socialGraph *SocialGraph) PrintLinks() {
	var root Name
	root = Name("Root")

	fmt.Printf("Printing all links adjacent to %d\n", root)

	var node Name
	for node = range socialGraph.Links[root] {
		fmt.Printf("Link: %d -> %d\n", root, node)
	}

  var m map[Name]struct{}
	fmt.Println("Printing all links.")
	for root, m = range socialGraph.Links {
		var vertex Name
		for vertex = range m {
			fmt.Printf("Link: %d -> %d\n",root, vertex)
		}
	}
}

// main method
func main() {

	var socialGraph *SocialGraph

	 socialGraph = NewSocialGraph()

	 var root Name = Name("Root")
	 var john Name = Name("John Smith")
   var per Name = Name("Per Jambeck")
	 var cynthia Name = Name("Cynthia Gibas")


	 socialGraph.AddEntity(root)
	 socialGraph.AddEntity(john)
	 socialGraph.AddEntity(per)
	 socialGraph.AddEntity(cynthia)

	 socialGraph.AddLink(root, john)
	 socialGraph.AddLink(root,per)
	 socialGraph.AddLink(root,cynthia)

   var mayo Name = Name("Mayo Smith")
	 var lorrie Name = Name("Lorrie Jambeck")
	 var ellie Name = Name("Ellie Vlocksen")

	 socialGraph.AddLink(john, mayo)
	 socialGraph.AddLink(john,lorrie)
	 socialGraph.AddLink(per,ellie)

	 socialGraph.PrintLinks()
}
