//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt package
import (
	"fmt"
	"math/rand"
	"time"
)

// constants
const (
	messagePassStart = iota
	messageTicketStart
	messagePassEnd
	messageTicketEnd
)

//Queue class
type Queue struct {
	waitPass    int
	waitTicket  int
	playPass    bool
	playTicket  bool
	queuePass   chan int
	queueTicket chan int
	message     chan int
}

// New method initialises queue
func (queue *Queue) New() {

	queue.message = make(chan int)
	queue.queuePass = make(chan int)
	queue.queueTicket = make(chan int)

	go func() {
		var message int
		for {
			select {
			case message = <-queue.message:
				switch message {
				case messagePassStart:
					queue.waitPass++
				case messagePassEnd:
					queue.playPass = false
				case messageTicketStart:
					queue.waitTicket++
				case messageTicketEnd:
					queue.playTicket = false
				}
				if queue.waitPass > 0 && queue.waitTicket > 0 && !queue.playPass && !queue.playTicket {
					queue.playPass = true
					queue.playTicket = true
					queue.waitTicket--
					queue.waitPass--
					queue.queuePass <- 1
					queue.queueTicket <- 1
				}
			}
		}
	}()
}

// StartTicketIssue starts the ticket issue
func (queue *Queue) StartTicketIssue() {
	queue.message <- messageTicketStart
	<-queue.queueTicket
}

// EndTicketIssue ends the ticket issue
func (queue *Queue) EndTicketIssue() {
	queue.message <- messageTicketEnd
}

//StartPass ends the Pass queue
func (queue *Queue) StartPass() {
	queue.message <- messagePassStart
	<-queue.queuePass
}

//EndPass ends the Pass queue
func (queue *Queue) EndPass() {
	queue.message <- messagePassEnd
}

//ticketIssue starts and ends the ticket issue
func ticketIssue(queue *Queue) {
	for {
		// Sleep up to 10 seconds.
		time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond)
		queue.StartTicketIssue()
		fmt.Println("Ticket Issue starts")
		// Sleep up to 2 seconds.
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
		fmt.Println("Ticket Issue ends")
		queue.EndTicketIssue()
	}
}

//passenger method starts and ends the pass queue
func passenger(queue *Queue) {
	for {
		// Sleep up to 10 seconds.
		time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond)
		queue.StartPass()
		fmt.Println("  Passenger starts")
		// Sleep up to 2 seconds.
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
		fmt.Println(" Passenger ends")

		queue.EndPass()
	}
}

// main method
func main() {
	var queue *Queue = &Queue{}
	queue.New()
	fmt.Println(queue)
	var i int
	for i = 0; i < 10; i++ {
		go passenger(queue)
	}
	var j int
	for j = 0; j < 5; j++ {
		go ticketIssue(queue)
	}
	select {}
}
