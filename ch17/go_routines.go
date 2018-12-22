//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt,sort, errors, math/rand and strings packages

import (
	"fmt"
	"time"
)

// Finder1 method
func Finder1(arr []string) {
	var str string
	for _, str = range arr {

		if str == "book" {

			fmt.Println("Inside Finder1 ", str)
			break
		}
	}

}

// Finder2 method
func Finder2(arr []string) {
	var str string

	for _, str = range arr {

		if str == "book" {

			fmt.Println("Inside Finder2 ", str)
			break
		}
	}

}

// main method
func main() {
	var theLibrary []string
	theLibrary = []string{"cup", "book", "book", "cup", "book"}
	go Finder1(theLibrary)
	go Finder2(theLibrary)
	<-time.After(time.Second * 1)
}
