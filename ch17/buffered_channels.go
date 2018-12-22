//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt,sort, errors, math/rand and strings packages

import (
	"fmt"
	"time"
)

// main method
func main() {
	var bufferedChan chan string
	bufferedChan = make(chan string, 3)
	go func() {
		bufferedChan <- "First"
		fmt.Println("Sent First")
		bufferedChan <- "Second"
		fmt.Println("Sent Second")
		bufferedChan <- "Third"
		fmt.Println("Sent Third")
	}()
	<-time.After(time.Second * 1)
	go func() {
		var firstMessage string
		firstMessage = <-bufferedChan
		fmt.Println("Receiving..")
		fmt.Println(firstMessage)
		var secondMessage string
		secondMessage = <-bufferedChan
		fmt.Println(secondMessage)
		var thirdMessage string
		thirdMessage = <-bufferedChan
		fmt.Println(thirdMessage)
	}()
	<-time.After(time.Second * 1)
}
