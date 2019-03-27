//main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
)

//IProces interface
type IProcess interface {
	process()
}

//Adapter struct
type Adapter struct {
	adaptee Adaptee
}

//Adapter class method process
func (adapter Adapter) process() {
	fmt.Println("Adapter process")
	adapter.adaptee.convert()
}

//Adaptee Struct
type Adaptee struct {
	adapterType int
}

// Adaptee class method convert
func (adaptee Adaptee) convert() {
	fmt.Println("Adaptee convert method")
}

// main method
func main() {

	var processor IProcess = Adapter{}

	processor.process()

}
