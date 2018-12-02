package main

import (
	"log"
	"os"
)

func main() {

	file, err := os.OpenFile("output.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
  	defer file.Close()
	log.SetOutput(file)
	log.Print("Writing to the output log")


}
