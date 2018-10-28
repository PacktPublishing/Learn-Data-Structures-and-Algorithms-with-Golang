///main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt, sync and atomic package
import (
	"fmt"
	"sync"
	"sync/atomic"
)

//Reference Counter
type ReferenceCounter struct {
	num     *uint32
	pool    *sync.Pool
	removed *uint32
}

//new Reference Counter method
func newReferenceCounter() ReferenceCounter {
	return ReferenceCounter{
		num:     new(uint32),
		pool:    &sync.Pool{},
		removed: new(uint32),
	}
}

// Add method
func (referenceCounter ReferenceCounter) Add() {
	atomic.AddUint32(referenceCounter.num, 1)
}

// Subtract method
func (referenceCounter ReferenceCounter) Subtract() {
	if atomic.AddUint32(referenceCounter.num, ^uint32(0)) == 0 {
		atomic.AddUint32(referenceCounter.removed, 1)

	}
}

// main method
func main() {

	var referenceCounter ReferenceCounter

	referenceCounter = newReferenceCounter()

	referenceCounter.Add()
	fmt.Println(*referenceCounter.num)

}
