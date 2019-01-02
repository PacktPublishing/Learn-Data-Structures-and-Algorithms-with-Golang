//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt and os packages
import (
	"fmt"
	"os"
)

// main method
func main() {

	var filenamestr string
	filenamestr = "config.jsn"
	var file *os.File
	var err error
	file, err = os.Open(filenamestr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("opened file", &file)
}
