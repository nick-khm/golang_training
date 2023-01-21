package main

import (
	"fmt"
	"time"
)

func reader(c chan string) {
	fmt.Println("Start read")
	defer fmt.Println("Stop read")
	for n := range c {
		fmt.Println(n)
	}
}

func main() {
	c := make(chan string)
	go reader(c)

	c <- "Bob"
	c <- "Alice"
	close(c)

	time.Sleep(5 * time.Second)

}
