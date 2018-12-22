//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt and time packages

import "fmt"
import "time"

// method function
func Method(completed chan bool) {
	fmt.Print("starting...")
	time.Sleep(time.Second)
	fmt.Println("completed")

	completed <- true
}

// main method
func main() {
  var completed chan bool
	completed = make(chan bool, 1)
	go Method(completed)

	<-completed
}
