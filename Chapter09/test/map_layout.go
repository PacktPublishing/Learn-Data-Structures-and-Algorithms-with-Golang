///main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
)

// Place class
type Place struct {

    Name string
    Latitude float64
    Longitude float64

}
// MapLayout class
type MapLayout struct {
	GraphNodes map[Place]struct{}
	Links map[Place]map[Place]struct{}
}

// NewMapLayout method
func NewMapLayout() *MapLayout {
	return &MapLayout{
		GraphNodes: make(map[Place]struct{}),
		Links: make(map[Place]map[Place]struct{}),
	}
}

// AddPlace method
func (mapLayout *MapLayout) AddPlace(place Place) bool {

	var exists bool
	if _, exists = mapLayout.GraphNodes[place]; exists {
		return true
	}
	mapLayout.GraphNodes[place] = struct{}{}
	return true
}

// Add Link
func (mapLayout *MapLayout) AddLink(place1 Place, place2 Place) {
	var exists bool
	if _, exists = mapLayout.GraphNodes[place1]; !exists {
		mapLayout.AddPlace(place1)
	}
	if _, exists = mapLayout.GraphNodes[place2]; !exists {
		mapLayout.AddPlace(place2)
	}

	if _, exists = mapLayout.Links[place1]; !exists {
		mapLayout.Links[place1] = make(map[Place]struct{})
	}
	mapLayout.Links[place1][place2] = struct{}{}

}

// PrintLinks method
func (mapLayout *MapLayout) PrintLinks() {
	var root Place
	root = Place{"Algeria",3,28}

	fmt.Printf("Printing all links adjacent to %s\n", root.Name)

	var node Place
	for node = range mapLayout.Links[root] {
		// Edge exists from u to v.
		fmt.Printf("Link: %s -> %s\n", root.Name, node.Name)
	}

  var m map[Place]struct{}
	fmt.Println("Printing all links.")
	for root, m = range mapLayout.Links {
		var vertex Place
		for vertex = range m {
			// Edge exists from u to v.
			fmt.Printf("Link: %s -> %s\n",root.Name, vertex.Name)
		}
	}
}

func main() {

	var mapLayout *MapLayout

	 mapLayout = NewMapLayout()

	 var root Place = Place{"Algeria",3,28}
	 var netherlands Place = Place{"Netherlands",5.75,52.5}

   var korea Place = Place{"Korea",124.1,-8.36}
	 var tunisia Place = Place{"Tunisia",9, 34}


	 mapLayout.AddPlace(root)
	 mapLayout.AddPlace(netherlands)
	 mapLayout.AddPlace(korea)
	 mapLayout.AddPlace(tunisia)

	 mapLayout.AddLink(root, netherlands)
	 mapLayout.AddLink(root,korea)
	 mapLayout.AddLink(root,tunisia)


	 var singapore Place = Place{"Singapore",103.8,1.36}
   var uae Place = Place{"UAE",54,24}
	 var japan Place = Place{"Japan",139.75, 35.68}

	 mapLayout.AddLink(korea, singapore)
	 mapLayout.AddLink(korea,japan)
	 mapLayout.AddLink(netherlands,uae)

	 mapLayout.PrintLinks()
}
