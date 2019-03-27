//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt package
import (
	"fmt"
	"strconv"
)

//Element class
type Element struct {
	elementValue int
}

// String method on Element class
func (element *Element) String() string {
	//fmt.Println(element.elementValue)

	return strconv.Itoa(element.elementValue)
}

// NewStack returns a new stack.
func (stack *Stack) New() {
	stack.elements = make([]*Element, 0)
}

// Stack is a basic LIFO stack that resizes as needed.
type Stack struct {
	elements     []*Element
	elementCount int
}

// Push adds a node to the stack.
func (stack *Stack) Push(element *Element) {
	stack.elements = append(stack.elements[:stack.elementCount], element)
	stack.elementCount = stack.elementCount + 1
}

// Pop removes and returns a node from the stack in last to first order.
func (stack *Stack) Pop() *Element {
	if stack.elementCount == 0 {
		return nil
	}

	var length int = len(stack.elements)
	var element *Element = stack.elements[length-1]
	//stack.elementCount = stack.elementCount - 1
	if length > 1 {
		stack.elements = stack.elements[:length-1]

	} else {
		stack.elements = stack.elements[0:]

	}
	stack.elementCount = len(stack.elements)
	return element
}

// main method
func main() {
	var stack *Stack = &Stack{}
	stack.New()
	var element1 *Element = &Element{3}
	var element2 *Element = &Element{5}
	var element3 *Element = &Element{7}
	var element4 *Element = &Element{9}
	stack.Push(element1)
	stack.Push(element2)
	stack.Push(element3)
	stack.Push(element4)
	fmt.Println(stack.Pop(), stack.Pop(), stack.Pop(), stack.Pop())
}
