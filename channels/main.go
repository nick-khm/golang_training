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

func one() {
	fmt.Println("one")
}

func hello(c chan string) {
	c <- "Hello there" // blocked until red by someone else
	fmt.Println("hello: finised")
	fmt.Printf("hello: %v\n", <-c)
}

func main() {
	c := make(chan string)

	go hello(c)
	msg := <-c
	fmt.Printf("main %v\n", msg)
	c <- "Main are writing"

	time.Sleep(10 * time.Second)
}
