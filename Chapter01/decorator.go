//main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
)

// IProcess Interface
type IProcess interface {
	process()
}

//ProcessClass struct
type ProcessClass struct{}

//ProcessClass method process
func (process *ProcessClass) process() {
	fmt.Println("ProcessClass process")
}

//ProcessDecorator struct
type ProcessDecorator struct {
	processInstance *ProcessClass
}

//ProcessDecorator class method process
func (decorator *ProcessDecorator) process() {
	if decorator.processInstance == nil {
		fmt.Println("ProcessDecorator  process")
	} else {
		fmt.Printf("ProcessDecorator  process  and ")
		decorator.processInstance.process()

	}
}

//main method
func main() {

	var process = &ProcessClass{}

	var decorator = &ProcessDecorator{}

	decorator.process()

	decorator.processInstance = process

	decorator.process()

}
