//main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
)

// IComposite interface
type IComposite interface {
	perform()
}

// Leaflet struct
type Leaflet struct {
	name string
}

// Leaflet class method perform
func (leaf *Leaflet) perform() {

	fmt.Println("Leaflet " + leaf.name)
}

// Branch struct
type Branch struct {
	leafs    []Leaflet
	name     string
	branches []Branch
}

// Branch class method perform
func (branch *Branch) perform() {

	fmt.Println("Branch: " + branch.name)
	for _, leaf := range branch.leafs {
		leaf.perform()
	}

	for _, branch := range branch.branches {
		branch.perform()
	}
}

// Branch class method add leaflet
func (branch *Branch) add(leaf Leaflet) {
	branch.leafs = append(branch.leafs, leaf)

}

//Branch class method addBranch branch
func (branch *Branch) addBranch(newBranch Branch) {

	branch.branches = append(branch.branches, newBranch)
}

//Branch class  method getLeaflets
func (branch *Branch) getLeaflets() []Leaflet {
	return branch.leafs
}

// main method
func main() {

	var branch = &Branch{name: "branch 1"}

	var leaf1 = Leaflet{name: "leaf 1"}
	var leaf2 = Leaflet{name: "leaf 2"}

	var branch2 = Branch{name: "branch 2"}

	branch.add(leaf1)
	branch.add(leaf2)
	branch.addBranch(branch2)

	branch.perform()

}
