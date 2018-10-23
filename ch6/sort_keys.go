///main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt and sort package
import (
	"fmt"
	"sort"
)

// A couple of type definitions to make the units clear.
type Mass float64
type Miles float64

// A Thing defines the properties of a  object.
type Thing struct {
	name          string
	mass          Mass
	distance      Miles
	meltingpoint  int
	freezingpoint int
}

// By is the type of a "less" function that defines the ordering of its Thing arguments.
type ByFactor func(Thing1 *Thing, Thing2 *Thing) bool

// Sort is a method on the function type, By, that sorts the argument slice according to the function.
func (byFactor ByFactor) Sort(Things []Thing) {
	var sortedThings *ThingSorter
	sortedThings = &ThingSorter{
		Things:   Things,
		byFactor: byFactor, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(sortedThings)
}

// ThingSorter joins a By function and a slice of Things to be sorted.
type ThingSorter struct {
	Things   []Thing
	byFactor func(Thing1 *Thing, Thing2 *Thing) bool // Closure used in the Less method.
}

// Len is part of sort.Interface.
func (ThingSorter *ThingSorter) Len() int {
	return len(ThingSorter.Things)
}

// Swap is part of sort.Interface.
func (ThingSorter *ThingSorter) Swap(i int, j int) {
	ThingSorter.Things[i], ThingSorter.Things[j] = ThingSorter.Things[j], ThingSorter.Things[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (ThingSorter *ThingSorter) Less(i int, j int) bool {
	return ThingSorter.byFactor(&ThingSorter.Things[i], &ThingSorter.Things[j])
}

// SortKeys demonstrates a technique for sorting a struct type using programmable sort criteria.
func main() {
	var Things = []Thing{
		{"IronRod", 0.055, 0.4, 3000, -180},
		{"SteelChair", 0.815, 0.7, 4000, -209},
		{"CopperBowl", 1.0, 1.0, 60, -30},
		{"BrassPot", 0.107, 1.5, 10000, -456},
	}
	// Closures that order the Thing structure.

	var name func(*Thing, *Thing) bool
	name = func(Thing1 *Thing, Thing2 *Thing) bool {
		return Thing1.name < Thing2.name
	}
	var mass func(*Thing, *Thing) bool
	mass = func(Thing1 *Thing, Thing2 *Thing) bool {
		return Thing1.mass < Thing2.mass
	}
	var distance func(*Thing, *Thing) bool
	distance = func(Thing1 *Thing, Thing2 *Thing) bool {
		return Thing1.distance < Thing2.distance
	}
	var decreasingDistance func(*Thing, *Thing) bool
	decreasingDistance = func(p1, p2 *Thing) bool {
		return distance(p2, p1)
	}

	// Sort the Things by the various criteria.
	ByFactor(name).Sort(Things)
	fmt.Println("By name:", Things)

	ByFactor(mass).Sort(Things)
	fmt.Println("By mass:", Things)

	ByFactor(distance).Sort(Things)
	fmt.Println("By distance:", Things)

	ByFactor(decreasingDistance).Sort(Things)
	fmt.Println("By decreasing distance:", Things)

}
