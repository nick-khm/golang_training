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

	result, err := cert.ParseCSV("students.csv")
	if err != nil {
		fmt.Printf("Error %v", err)
		os.Exit(1)
	}
	for _, item := range result {

		err = saver.Save(*item)
		if err != nil {
			fmt.Printf("Error %v", err)
		}
	}

}
