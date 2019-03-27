//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt package
import (
	"fmt"
)

//Set class
type Set struct {
	integerMap map[int]bool
}

//create the map of integer and bool
func (set *Set) New() {
	set.integerMap = make(map[int]bool)
}

// adds the element to the set
func (set *Set) AddElement(element int) {
	if !set.ContainsElement(element) {
		set.integerMap[element] = true
	}
}

//deletes the element from the set
func (set *Set) DeleteElement(element int) {

	delete(set.integerMap, element)
}

//checks if element is in the set
func (set *Set) ContainsElement(element int) bool {
	var exists bool
	_, exists = set.integerMap[element]

	return exists
}

//Intersect method returns the set which intersects with anotherSet
func (set *Set) Intersect(anotherSet *Set) *Set {
	var intersectSet = &Set{}
	intersectSet.New()
	var value int
	for value, _ = range set.integerMap {
		if anotherSet.ContainsElement(value) {

			intersectSet.AddElement(value)
		}
	}
	return intersectSet
}

//Union method returns the set which is union of the set with anotherSet
func (set *Set) Union(anotherSet *Set) *Set {
	var unionSet = &Set{}
	unionSet.New()
	var value int
	for value, _ = range set.integerMap {
		unionSet.AddElement(value)
	}

	for value, _ = range anotherSet.integerMap {
		unionSet.AddElement(value)
	}

	return unionSet
}

// main method
func main() {

	var set *Set
	set = &Set{}

	set.New()

	set.AddElement(1)
	set.AddElement(2)

	fmt.Println("initial set", set)
	fmt.Println(set.ContainsElement(1))

	var anotherSet *Set
	anotherSet = &Set{}

	anotherSet.New()
	anotherSet.AddElement(2)
	anotherSet.AddElement(4)
	anotherSet.AddElement(5)

	fmt.Println("another set", set)

	fmt.Println("intersection of sets ", set.Intersect(anotherSet))

	fmt.Println("union of sets ", set.Union(anotherSet))

}
