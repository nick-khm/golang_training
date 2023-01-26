package main

import (
	"fmt"
	"time"
)

func write(c chan string) {
	for i := 1; i < 5; i++ {
		c <- fmt.Sprintf("write %v", i)
	}
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go write(c1)
	go write(c2)

	for {
		time.Sleep(1 * time.Second)
		select {
		case v := <-c1:
			fmt.Println("main c1 val", v)
		case v := <-c2:
			fmt.Println("main c2 val", v)
		default:
			fmt.Println("Nothing")
		}
	}
}
