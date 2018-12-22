//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt,sort, errors, math/rand and strings packages

import (
	"fmt"
	"time"
)

// CallMethod method
func CallMethod() {
	fmt.Println("Calling method")
}

// main method
func main() {
	go CallMethod()
	time.Sleep(1 * time.Second)
	fmt.Println("From Main Method")
}
