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
	myChan := make(chan string)

	go func() {
		myChan <- "Message!"
	}()
  <-time.After(time.Second * 1)
	select {
	case myChan <- "message":
		fmt.Println("sent the message")
	default:
		fmt.Println("no message sent")
	}
	<-time.After(time.Second * 1)
	select {
	case myChan <- "message":
		fmt.Println("sent the message")
	default:
		fmt.Println("no message sent")
	}
	<-time.After(time.Second * 1)
}
