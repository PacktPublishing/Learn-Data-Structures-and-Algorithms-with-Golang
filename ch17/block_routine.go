//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt,sort, errors, math/rand and strings packages

import (
	//"fmt"
	"time"
)

// main method
func main() {
 doneChan := make(chan string)
 go func() {
  doneChan <- "completed!"
 }()
 <-time.After(time.Second * 1)
 <-doneChan
 <-time.After(time.Second * 1)
}
