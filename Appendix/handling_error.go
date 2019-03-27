//main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt and errors packages

import(
	"fmt"
	"errors"

)

//First Func method
func FirstFunc(v interface{}) (interface{}, error) {
	var ok bool

	if !ok {
		return nil, errors.New("false error")
	}
	return v, nil
}

//SecondFunc method
func SecondFunc() {
	defer func() {
		var err interface{}
		if err = recover(); err != nil {
			fmt.Println("recovering error ", err)
		}
	}()
	var v interface{}
	v = struct{}{}
	var err error
	if _, err = FirstFunc(v); err != nil {
		panic(err)
	}

	fmt.Println("The error never happen")
}

//main method
func main() {
	SecondFunc()
	fmt.Println("The execution ended")
}
