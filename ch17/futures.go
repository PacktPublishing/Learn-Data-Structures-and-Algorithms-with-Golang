//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt, io/ioutil and net/http packages

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// ResponseData class
type ResponseData struct {
	Body  []byte
	Error error
}

// GetResponse method
func GetResponse(url string) <-chan ResponseData {
	var channel chan ResponseData
	channel = make(chan ResponseData, 1)

	go func() {
		var body []byte
		var err error
		var resp *http.Response
		resp, err = http.Get(url)
		defer resp.Body.Close()

		body, err = ioutil.ReadAll(resp.Body)

		channel <- ResponseData{Body: body, Error: err}
	}()

	return channel
}

// main method
func main() {
	var future <-chan ResponseData
	future = GetResponse("http://www.google.com")

	var body ResponseData
	body = <-future
	fmt.Println(" URL Response: %#v", string(body.Body))
	fmt.Println("URL error: %#v", body.Error)
}
