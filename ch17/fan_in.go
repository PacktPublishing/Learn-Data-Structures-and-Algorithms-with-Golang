//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt and time packages

import (
    "fmt"
    "time"
)


func producer(ch chan int, d time.Duration) {
    var i int
    for i= 0; i < 10; i++{
        ch <- i
        i++
        time.Sleep(d)
    }
}

func reader(out chan int) {
    var x int
    for x = range out {
        fmt.Println(x)
    }
}
// main method
func main() {
    var ch chan int
    ch = make(chan int)
    var out chan int
    out = make(chan int)
    go producer(ch, 100*time.Millisecond)
    go producer(ch, 250*time.Millisecond)
    go reader(out)
    var i int
    for i = range ch {
        out <- i
    }
}
