package main

import (
	"fmt"
	"os"

	"training.go/dictionary/dictionary"
)

func main() {
	d, err := dictionary.New("./badger")
	handleError(err)
	defer d.Close()

	d.Remove("toto")
	words, entries, err := d.List()
	handleError(err)
	for _, word := range words {

		fmt.Println(entries[word])
	}
}

func handleError(e error) {
	if e != nil {
		fmt.Printf("Dictionary error: %v\n", e)
		os.Exit(1)
	}
}
