//main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt and bytes package
import (
	"bytes"
	"fmt"
)

// ThueMorseSequence method
func ThueMorseSequence(buffer *bytes.Buffer) {

	var b int
	var currLength int
	var currBytes []byte
	for b, currLength, currBytes = 0, buffer.Len(), buffer.Bytes(); b < currLength; b++ {
		if currBytes[b] == '1' {
			buffer.WriteByte('0')
		} else {
			buffer.WriteByte('1')
		}
	}
}

// main method
func main() {
	var buffer bytes.Buffer
	buffer.WriteByte('0')
	fmt.Println(buffer.String())
	var i int
	for i = 2; i <= 7; i++ {
		ThueMorseSequence(&buffer)
		fmt.Println(buffer.String())
	}
}
