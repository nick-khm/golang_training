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

	d.Add("golang", "A wonderful language!")
	entry, err := d.Get("golang")
	handleError(err)
	fmt.Println(entry)

}

func handleError(e error) {
	if e != nil {
		fmt.Printf("Dictionary error: %v\n", e)
		os.Exit(1)
	}
}
