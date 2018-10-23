///main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
)

//Circular Queue
type CircularQueue struct {
	size  int
	nodes []interface{}
	head  int
	last  int
}

// NewCircularQueue method
func NewQueue(num int) *CircularQueue {
	var circularQueue CircularQueue
	circularQueue = CircularQueue{size: num + 1, head: 0, last: 0}
	circularQueue.nodes = make([]interface{}, circularQueue.size)
	return &circularQueue
}

//  IsUnUsed method
func (circularQueue CircularQueue) IsUnUsed() bool {
	return circularQueue.head == circularQueue.last
}

// IsComplete method
func (circularQueue CircularQueue) IsComplete() bool {
	return circularQueue.head == (circularQueue.last+1)%circularQueue.size
}

// Add method
func (circularQueue *CircularQueue) Add(element interface{}) {
	if circularQueue.IsComplete() {
		panic("Queue is Completely Utilized")
	}
	circularQueue.nodes[circularQueue.last] = element
	circularQueue.last = (circularQueue.last + 1) % circularQueue.size
}

//MoveOneStep method
func (circularQueue *CircularQueue) MoveOneStep() (element interface{}) {
	if circularQueue.IsUnUsed() {
		return nil
	}
	element = circularQueue.nodes[circularQueue.head]
	circularQueue.head = (circularQueue.head + 1) % circularQueue.size
	return
}

// main method
func main() {
	var circularQueue *CircularQueue

	circularQueue = NewQueue(5)

	circularQueue.Add(1)
	circularQueue.Add(2)
	circularQueue.Add(3)
	circularQueue.Add(4)
	circularQueue.Add(5)

	fmt.Println(circularQueue.nodes)

}
