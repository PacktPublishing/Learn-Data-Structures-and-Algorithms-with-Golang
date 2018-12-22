//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt packages

import (
	"fmt"
)

// GetFibonacciNumbersLesserThan method
func GetFibonacciNumbersLesserThan(n int) chan int {
	var channel chan int
	channel = make(chan int)
	go func() {
		var i int
		var j int
		for i, j = 0, 1; i < n; i, j = i+j, i {
			channel <- i
		}
		close(channel)
	}()
	return channel
}

// main method
func main() {
	var fibonacci int
	for fibonacci = range GetFibonacciNumbersLesserThan(1000) {
		fmt.Println(fibonacci)
	}
}
