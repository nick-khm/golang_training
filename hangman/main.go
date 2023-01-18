package main

import (
	"fmt"
	"os"

	"training.go/hangman/hangman"
	"training.go/hangman/hangman/dictionary"
)

func main() {
	err := dictionary.Load("./hangman/words.txt")
	if err != nil {
		fmt.Printf("Some dict error %v\n", err)
		os.Exit(1)
	}
	word := dictionary.PickWord()

	g := hangman.New(6, word)
	hangman.DrawWelcome()

	guess := ""
	for {
		hangman.Draw(g, guess)
		switch g.State {
		case hangman.Won, hangman.Lost:
			os.Exit(0)
		}
		l, err := hangman.ReadGuess()
		if err != nil {
			fmt.Printf("Could not read from terminal: %v", err)
			os.Exit(1)
		}
		guess = l
		g.MakeAGuess(guess)
	}

}
