//main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing errors,log and time packages
import (
	"errors"
	"log"
	"time"
)

// delayTimeOut method
func delayTimeOut(channel chan interface{}, timeOut time.Duration) (interface{}, error) {
	log.Printf("delayTimeOut enter")
	defer log.Printf("delayTimeOut exit")
	var data interface{}
	select {
	case <-time.After(timeOut):
		return nil, errors.New("delayTimeOut time out")
	case data = <-channel:
		return data, nil
	}
}

//main method
func main() {

	channel := make(chan interface{})
	go func() {
		var err error
		var data interface{}
		data, err = delayTimeOut(channel, time.Second)
		if err != nil {
			log.Printf("error %v", err)
			return
		}
		log.Printf("data %v", data)
	}()

	channel <- struct{}{}
	time.Sleep(time.Second * 2)

	go func() {
		var err error
		var data interface{}
		data, err = delayTimeOut(channel, time.Second)
		if err != nil {
			log.Printf("error %v", err)
			return
		}
		log.Printf("data %v", data)
	}()
	time.Sleep(time.Second * 2)

}
