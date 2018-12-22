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

	go func() {
		fmt.Println("running in a go routine")
	}()
	<-time.After(time.Second * 1)
}
