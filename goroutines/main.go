package main

import (
	"fmt"
	"time"
)

func longOperation(duration int) {
	time.Sleep(time.Duration(duration) * time.Second)
	fmt.Printf("Operation finished! Took %d s\n", duration)
}

func main() {
	fmt.Println("Starting first operation")
	go longOperation(1)

	fmt.Println("Starting second operation")
	longOperation(1)

	time.Sleep(time.Duration(3) * time.Second)

}
