package main

import (
	"flag"
	"fmt"

	"os"

	"training.go/certificate/cert"
	"training.go/certificate/cert/html"
	"training.go/certificate/pdf"
)

func main() {
	outputType := flag.String("type", "pdf", "Output type of the certificate.")
	flag.Parse()

	var saver cert.Saver
	var err error
	switch *outputType {
	case "html":
		saver, err = html.New("output")
	case "pdf":
		saver, err = pdf.New("output")
	default:
		fmt.Printf("Unkown output type. got=%v\n", *outputType)
		os.Exit(1)
	}
	if err != nil {
		fmt.Printf("Error during certificate creation: %v", err)
		os.Exit(1)
	}

	c, err := cert.New("Golang programming", "Bob Dylan", "2023-01-11")
	if err != nil {
		fmt.Printf("Error during certificate creation: %v", err)
		os.Exit(1)
	}
	saver.Save(*c)

}
