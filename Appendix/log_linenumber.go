//main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing path,runtime,fmt, log and time packages

import(
	"path"
	"runtime"
	"fmt"
	"log"
	"time"
)

//checkPoint method
func checkPoint() string {
    pc, file, line, _ := runtime.Caller(1)
    return fmt.Sprintf("\033[31m%v %s %s %d\x1b[0m", time.Now(),
    	runtime.FuncForPC(pc).Name(), path.Base(file), line)
}

//method1
func method1(){
	fmt.Println(checkPoint())
}

//main method
func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Println("logging the time and flags")

	method1()

}
