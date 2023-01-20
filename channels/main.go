package main

import (
	"fmt"
	"time"
)

func ping(c chan string) {
	for i := 1; ; i++ {
		c <- fmt.Sprintf("ping %v", i)
	}
}

func pong(c chan string) {
	for i := 100; ; i += 100 {
		c <- fmt.Sprintf("pong %v", i)
	}
}

func print(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	c := make(chan string)
	go ping(c)
	go pong(c)
	go print(c)

	time.Sleep(10 * time.Second)
}
