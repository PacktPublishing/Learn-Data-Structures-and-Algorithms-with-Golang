//main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing errors,context,log and time packages

import (
	"errors"
	"golang.org/x/net/context"
	"log"
	"time"
)

// main method
func main() {

  var delay time.Duration

	delay = time.Millisecond

	var cancel context.CancelFunc

  var contex context.Context

	contex, cancel = context.WithTimeout(context.Background(), delay)

	go func(context.Context) {
		<-contex.Done()
		log.Printf("contex done")
	}(contex)

	_ = cancel

	time.Sleep(delay * 2)

	log.Printf("contex end %v", contex.Err())

	channel := make(chan struct{})

	var err error
	go func(chan struct{}) {
		select {
		case <-time.After(delay):
			err = errors.New("ch delay")
		case <-channel:
		}
		log.Printf("channel done")
	}(channel)


	time.Sleep(delay * 2)

	log.Printf("channel end %v", err)
}
