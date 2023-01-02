package hangman

import "testing"

func TestGameAlreadyGuessed(t *testing.T) {
	g, _ := New(6, "alice")
	g.MakeAGuess("a")
	validState(t, GoodGuess, g.State)
	g.MakeAGuess("a")
	validState(t, AlreadyGuessed, g.State)
}

func TestGameBadGuess(t *testing.T) {
	g, _ := New(6, "alice")
	g.MakeAGuess("d")
	validState(t, BadGuess, g.State)
}

func TestGameWon(t *testing.T) {
	g, _ := New(6, "bob")
	g.MakeAGuess("b")
	g.MakeAGuess("o")
	validState(t, Won, g.State)
}

func TestGameLost(t *testing.T) {
	g, _ := New(3, "bob")
	g.MakeAGuess("a")
	g.MakeAGuess("l")
	g.MakeAGuess("i")
	validState(t, Lost, g.State)
}

func TestGameLost2(t *testing.T) {
	g, _ := New(3, "bob")
	g.MakeAGuess("b")
	g.MakeAGuess("l")
	g.MakeAGuess("i")
	g.MakeAGuess("c")
	validState(t, Lost, g.State)
}

func TestGameLost3(t *testing.T) {
	g, _ := New(3, "bob")
	g.MakeAGuess("b")
	g.MakeAGuess("l")
	g.MakeAGuess("i")
	g.MakeAGuess("i")
	g.MakeAGuess("i")
	g.MakeAGuess("c")
	validState(t, Lost, g.State)
}
