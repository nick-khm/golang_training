package hangman

import "fmt"

func DrawWelcome() {
	fmt.Println("Game HANGMAN")
}

func Draw(g *Game, guess string) {
	drawTurns(g.TurnsLeft)
	drawState(g, guess)
}

func drawTurns(l int) {
	switch l {
	case 6:
		fmt.Println(`
		  /-\
		 [o_o]
		/|___|\
		 // \\
		`)

	case 5:
		fmt.Println(`
		  /-\
		 [o_o]
		/|¤__|\
		 // \\
		`)

	case 4:
		fmt.Println(`
		  /-\
		 [o_o]
		/|¤¤_|\
		 // \\
		`)

	case 3:
		fmt.Println(`
		  /-\
		 [o_o]
		/|¤¤¤|\
		 // \\
		`)

	case 2:
		fmt.Println(`
		  /-\
		 [-_o]
		/|¤¤¤|\
		 // \\
		`)

	case 1:
		fmt.Println(`
		  /-\
		 [-_-]
		/|¤¤¤|\
		 // \\
		`)

	case 0:
		fmt.Println(`
		  /¤\
		 [x_x]
		/|¤¤¤|\
		 // \\
		`)

	}
}

func drawState(g *Game, guess string) {
	fmt.Print("Guessed: ")
	drawLetters(g.FoundLetters)

	fmt.Printf("Used: ")
	drawLetters(g.UsedLetters)

	switch g.State {
	case GoodGuess:
		fmt.Print("Good guess!\n")
	case AlreadyGuessed:
		fmt.Printf("Letter '%s' was already used\n", guess)
	case BadGuess:
		fmt.Printf("Letter '%s' is not in the word\n", guess)
	case Lost:
		fmt.Print("LOST! The word was:")
		drawLetters(g.Letters)
	case Won:
		fmt.Print("WON! The word was:")
		drawLetters(g.Letters)

	}
	fmt.Println()
	fmt.Println()
}

func drawLetters(l []string) {
	for _, c := range l {
		fmt.Printf("%v ", c)
	}
	fmt.Println()
}
