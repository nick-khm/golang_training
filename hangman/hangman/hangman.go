package hangman

import (
	"fmt"
	"strings"
)

type GameStateType string

const (
	GoodGuess      GameStateType = "goodGuess"
	AlreadyGuessed GameStateType = "alreadyGuessed"
	BadGuess       GameStateType = "badGuess"
	Lost           GameStateType = "lost"
	Won            GameStateType = "won"
)

type Game struct {
	State        GameStateType // Game state
	Letters      []string      // Letters in the word to find
	FoundLetters []string      // Good guesses
	UsedLetters  []string      // Used letters
	TurnsLeft    int           // Remaining attempts
}

func New(turns int, word string) (*Game, error) {
	if len(word) < 3 {
		return nil, fmt.Errorf("word '%s' must be at least 3 chars, got=%v", word, len(word))
	}

	letters := strings.Split(strings.ToUpper(word), "")
	found := make([]string, len(letters))
	for i := 0; i < len(letters); i++ {
		found[i] = "_"
	}

	g := &Game{
		State:        "",
		Letters:      letters,
		FoundLetters: found,
		UsedLetters:  []string{},
		TurnsLeft:    turns,
	}
	return g, nil
}

func (g *Game) MakeAGuess(guess string) {

	guess = strings.ToUpper(guess)
	if letterInWord(guess, g.UsedLetters) {
		g.State = AlreadyGuessed
	} else if letterInWord(guess, g.Letters) {
		g.State = GoodGuess
		g.RevealLetter(guess)
		if hasWon(g.Letters, g.FoundLetters) {
			g.State = Won
		}
	} else {
		g.State = BadGuess
		g.LoseTurn(guess)

		if g.TurnsLeft <= 0 {
			g.State = Lost
		}
	}
}

func (g *Game) RevealLetter(guess string) {
	g.UsedLetters = append(g.UsedLetters, guess)
	for i, l := range g.Letters {
		if l == guess {
			g.FoundLetters[i] = guess
		}
	}
}

func hasWon(letters []string, foundLetters []string) bool {
	for i := range letters {
		if letters[i] != foundLetters[i] {
			return false
		}
	}
	return true
}

func (g *Game) LoseTurn(guess string) {
	g.TurnsLeft--
	g.UsedLetters = append(g.UsedLetters, guess)
}

func letterInWord(guess string, letters []string) bool {
	for _, l := range letters {
		if l == guess {
			return true
		}
	}
	return false
}
