//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt and os packages
import (
	"fmt"
	"os"
)

// ErrLoadFailure class
type ErrLoadFailure struct {
	filename string
}

// Error method
func (err *ErrLoadFailure) Error() string {
	return err.filename
}

// New method
func New(filename string) error {
	return &ErrLoadFailure{filename}
}

// main method
func main() {
	var filenamestr string
	filenamestr = "config.jsn"
	var file *os.File
	var err error
	file, err = os.Open(filenamestr)
	if err != nil {
		err = New(filenamestr)
		fmt.Println("File ", err.Error(), "is not there")
		return
	}

	fmt.Println("opened file", &file)

}
