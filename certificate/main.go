package main

import (
	"fmt"

	"os"

	"training.go/certificate/cert"
	"training.go/certificate/cert/html"
)

func main() {
	c, err := cert.New("Golang programming", "Bob Dylan", "2023-01-11")
	if err != nil {
		fmt.Printf("Error during certificate creation: %v", err)
		os.Exit(1)
	}

	var saver cert.Saver
	saver, err = html.New("output")
	if err != nil {
		fmt.Printf("Error during pdf creation: %v", err)
		os.Exit(1)
	}
	saver.Save(*c)

}
