///main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt, encoding/json, io/ioutil and os packages
import (
	"fmt"
  "encoding/json"
	"io/ioutil"
)
//Persons class
type Persons []Person
//Person class
type Person struct {
	Name   string
	Type   string
	Age    int
}
// main method
func main() {

	var bytes []byte
	var err error
	bytes, err = ioutil.ReadFile("test.json")

	var persons Persons
	err = json.Unmarshal(bytes, &persons)
	if err != nil {
		panic(err)
	}

	fmt.Println(persons)

}
