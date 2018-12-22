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
	var myChan chan string
	myChan = make(chan string)

	go func() {
		myChan <- "Message!"
	}()

	var msg string
	select {
	case msg = <-myChan:
		fmt.Println(msg)
	default:
		fmt.Println("No Message")
	}
	<-time.After(time.Second * 1)
	select {
	case msg = <-myChan:
		fmt.Println(msg)
	default:
		fmt.Println("No Message")
	}
	<-time.After(time.Second * 1)
}
