//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt package
import (
	"fmt"
  "sync"
)

//Reference Counter Class
type ReferenceCounter struct {
	num     *uint32
	pool    *sync.Pool
	removed *uint32
}

//new Reference Counter method
func newReferenceCounter() *ReferenceCounter {
	return &ReferenceCounter{
		num:     new(uint32),
		pool:    &sync.Pool{},
		removed: new(uint32),
	}
}
// New method of Stack class
func (stack *Stack) New()  {
    stack.references = make([]*ReferenceCounter,0)
}

// Stack class
type Stack struct {
    references []*ReferenceCounter
    Count int
}

// Push method
func (stack *Stack) Push(reference *ReferenceCounter) {
    stack.references = append(stack.references[:stack.Count], reference)
    stack.Count = stack.Count + 1
}

// Pop method
func (stack *Stack) Pop() *ReferenceCounter {
    if stack.Count == 0 {
        return nil
    }

		var length int = len(stack.references)
		var reference *ReferenceCounter = stack.references[length -1]
		if length > 1 {
		   stack.references = stack.references[:length-1]

	  } else {
		   stack.references = stack.references[0:]

		}
    stack.Count = len(stack.references)
    return reference
}

// main method
func main() {
 var stack *Stack = &Stack{}
 stack.New()
 var reference1 *ReferenceCounter = newReferenceCounter()
 var reference2 *ReferenceCounter = newReferenceCounter()
 var reference3 *ReferenceCounter = newReferenceCounter()
 var reference4 *ReferenceCounter = newReferenceCounter()

 stack.Push(reference1)
 stack.Push(reference2)
 stack.Push(reference3)
 stack.Push(reference4)
 fmt.Println(stack.Pop(), stack.Pop(), stack.Pop(), stack.Pop())
}
