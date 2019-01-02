//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing flag and fmt packages
import (
	"flag"
	"fmt"
)

// main method
func main() {
	var (
		count    = flag.String("count", "count records", "number of records")
		source   = flag.String("source", "data source", "service name")
		filename = flag.String("filename", "File Name", "Data file")
	)
	flag.Parse()

	fmt.Println(*count, *source, *filename)
}
