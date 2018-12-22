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
	var theLibrary []string
	theLibrary = []string{"book1", "book2", "book3"}

	bookChan := make(chan string)

	go func(mine []string) {
		var item string
		for _, item = range mine {
			bookChan <- item //send
		}
	}(theLibrary)

	go func() {
		var i int
		for i = 0; i < 3; i++ {
			foundBook := <-bookChan
			fmt.Println("BookReader: Received " + foundBook + " from Finder")
		}
	}()
	<-time.After(time.Second * 5)
}
