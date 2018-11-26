///main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
	"math"
)

// Place class
type Place struct {
	Name      string
	Latitude  float64
	Longitude float64
}

//RoutePlan class
type RoutePlan struct {
	legs []RouteLeg
	Total_Time float64
}

//RouteLeg class
type RouteLeg struct {
	Origin       Place
	Destination  Place
	MaximumSpeed float64
	Distance     float64
}

// RouteMap class
type RouteMap struct {
	GraphNodes map[Place]struct{}
	Links      map[Place]map[Place]struct{}
}

// NewRouteMap method
func NewRouteMap() *RouteMap {
	return &RouteMap{
		GraphNodes: make(map[Place]struct{}),
		Links:      make(map[Place]map[Place]struct{}),
	}
}

// AddPlace method
func (routeMap *RouteMap) AddPlace(place Place) bool {

	var exists bool
	if _, exists = routeMap.GraphNodes[place]; exists {
		return true
	}
	routeMap.GraphNodes[place] = struct{}{}
	return true
}

// Add Link
func (routeMap *RouteMap) AddLink(place1 Place, place2 Place) {
	var exists bool
	if _, exists = routeMap.GraphNodes[place1]; !exists {
		routeMap.AddPlace(place1)
	}
	if _, exists = routeMap.GraphNodes[place2]; !exists {
		routeMap.AddPlace(place2)
	}

	if _, exists = routeMap.Links[place1]; !exists {
		routeMap.Links[place1] = make(map[Place]struct{})
	}
	routeMap.Links[place1][place2] = struct{}{}

}

// PrintLinks method
func (routeMap *RouteMap) PrintLinks() {
	var root Place
	root = Place{"London", 3, 28}

	fmt.Printf("Printing all links adjacent to %s\n", root.Name)

	var node Place
	for node = range routeMap.Links[root] {
		// Edge exists from u to v.
		fmt.Printf("Link: %s -> %s\n", root.Name, node.Name)
	}

	var m map[Place]struct{}
	fmt.Println("Printing all links.")
	for root, m = range routeMap.Links {
		var vertex Place
		for vertex = range m {
			// Edge exists from u to v.
			fmt.Printf("Link: %s -> %s\n", root.Name, vertex.Name)
		}
	}
}

//CreateRoutePlan method
func CreateRoutePlan(routeMap *RouteMap) RoutePlan {
	var routePlan RoutePlan
	routePlan = RoutePlan{}
	var node Place
	var m map[Place]struct{}
  var total_time float64
	total_time = 0
	for node, m = range routeMap.Links {
		var node2 Place
		for node2 = range m {
			var leg RouteLeg
			leg = RouteLeg{Origin: node, Destination: node2, MaximumSpeed: 80, Distance: GetDistance(node, node2)}
      routePlan.legs = append(routePlan.legs, leg)
			total_time = total_time + (leg.Distance/leg.MaximumSpeed)

		}
	}
	routePlan.Total_Time = total_time
	return routePlan
}

//PrintRoutePlan method
func PrintRoutePlan(routePlan RoutePlan) {

	var leg RouteLeg
	for _,leg = range routePlan.legs {
		fmt.Println("Origin", leg.Origin.Name, "Destination", leg.Destination.Name, "MaximumSpeed", leg.MaximumSpeed, "Distance", leg.Distance)
    var time float64
		time = (leg.Distance/leg.MaximumSpeed)
    fmt.Println("Leg time",time )
	}

	fmt.Println("Total_Time",routePlan.Total_Time)
}

const radius_earth = 6371

//ConvertDegreeToRadian method
func ConvertDegreeToRadian(degree float64) float64 {
	return degree * (math.Pi / 180)
}

// GetDistance method
func GetDistance(origin Place, destination Place) float64 {
	var lat1 float64
	var long1 float64
	var lat2 float64
	var long2 float64

	lat1 = origin.Latitude
	long1 = origin.Longitude

	lat2 = destination.Latitude
	long2 = destination.Longitude

	var diffLat float64
	diffLat = ConvertDegreeToRadian(lat2 - lat1)
	var diffLong float64
	diffLong = ConvertDegreeToRadian(long2 - long1)
	var a float64
	a = math.Sin(diffLat/2)*math.Sin(diffLat/2) +
		math.Cos(ConvertDegreeToRadian(lat1))*math.Cos(ConvertDegreeToRadian(lat2))*
			math.Sin(diffLong/2)*math.Sin(diffLong/2)
	var c float64
	c = 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	var distance float64
	distance = radius_earth * c
	return distance
}

// main method
func main() {

	var routeMap *RouteMap

	routeMap = NewRouteMap()

	var root Place = Place{"London", 3, 28}
	var paris Place = Place{"Paris", 5.75, 52.5}

	var berlin Place = Place{"Berlin", 124.1, -8.36}
	var amsterdam Place = Place{"Amsterdam", 9, 34}

	routeMap.AddPlace(root)
	routeMap.AddPlace(paris)
	routeMap.AddPlace(berlin)
	routeMap.AddPlace(amsterdam)

	routeMap.AddLink(root, paris)
	routeMap.AddLink(root, berlin)
	routeMap.AddLink(root, amsterdam)

	var barcelona Place = Place{"Barcelona", 103.8, 1.36}
	var vienna Place = Place{"Vienna", 54, 24}
	var prague Place = Place{"Prague", 139.75, 35.68}

	routeMap.AddLink(paris, barcelona)
	routeMap.AddLink(paris, vienna)
	routeMap.AddLink(berlin, prague)

	routeMap.PrintLinks()

	var routePlan RoutePlan
	routePlan = CreateRoutePlan(routeMap)

	PrintRoutePlan(routePlan)
}
