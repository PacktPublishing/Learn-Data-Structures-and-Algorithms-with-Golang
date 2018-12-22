//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt package

import (
	"fmt"
)

//GetPipeline method
func GetPipeline(numbers []int) <-chan int {
	var outChannel chan int
	outChannel = make(chan int)
	go func() {
		var n int
		for _, n = range numbers {
			outChannel <- n
		}
		close(outChannel)
	}()
	return outChannel
}

// GetSquare method
func GetSquare(inChannel <-chan int) <-chan int {
	var outChannel chan int
	outChannel = make(chan int)
	go func() {
		var n int
		for n = range inChannel {
			outChannel <- n * n
		}
		close(outChannel)
	}()
	return outChannel
}

//GetFanIn method
func GetFanIn(input1 <-chan int, input2 <-chan int) <-chan int {
	var channel chan int
	channel = make(chan int)
	go func() {
		for {
			select {
			case s := <-input1:
				channel <- s
			case s := <-input2:
				channel <- s
			}
		}
	}()
	return channel
}

// main method
func main() {
	var randomNums []int
	randomNums = []int{13, 44, 56, 99, 9, 45, 67, 90, 78, 23}
	var inputChannel <-chan int
	inputChannel = GetPipeline(randomNums)

	var c1 <-chan int
	c1 = GetSquare(inputChannel)
	var c2 <-chan int
	c2 = GetSquare(inputChannel)

	var channel <-chan int
	channel = GetFanIn(c1, c2)
	var sum int
	sum = 0

	var i int
	for i = 0; i < len(randomNums); i++ {
		sum += <-channel
	}
	fmt.Println(" Sum of Squares:", sum)
}
