///main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt and os packages
import (
	"fmt"
	"os"
)

//main method
func main() {
	var jsonFile *os.File
	var err error
	jsonFile, err = os.Open("test.json")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Opened test.json")
	defer jsonFile.Close()

}
