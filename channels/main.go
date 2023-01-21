package main

import (
	"fmt"
	"time"
)

func longOperation(done chan<- bool) {
	fmt.Println("longOperation: Started long operation...")
	time.Sleep(2 * time.Second)
	fmt.Println("longOperation: Done!")
	done <- true
}

func main() {
	done := make(chan bool)
	go longOperation(done)
	<-done
	fmt.Println("Back to main")

}
