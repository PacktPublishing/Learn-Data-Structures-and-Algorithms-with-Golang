//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing os and json packages
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	//"os"
)

// Configuration class
type Configuration []Record

//Record class
type Record struct {
	Name  string
	Ssn   string
	Place string
}

// Load Config method
func LoadConfig(file string) Configuration {
	var config Configuration
	var bytes []byte
	var err error
	bytes, err = ioutil.ReadFile(file)

	err = json.Unmarshal(bytes, &config)
	if err != nil {
		panic(err)
	}
	return config
}

// main method
func main() {

	var config Configuration

	config = LoadConfig("config.json")

	fmt.Println(config)
}
