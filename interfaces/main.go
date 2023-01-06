package main

import "fmt"

type Instrument interface {
	Play()
}

type Amplifier interface {
	Connect(amp string)
}

type Guitar struct {
	Strings int
}

type Piano struct {
	Keys int
}

func (g Guitar) Play() {
	fmt.Printf("Guitar plays with %d strings\n", g.Strings)
}

func (g Guitar) Connect(amp string) {
	fmt.Printf("Connected to %v\n", amp)
}

func (p Piano) Play() {
	fmt.Printf("Piano plays with %d strings\n", p.Keys)
}

func main() {
	var instr Instrument
	instr = &Guitar{5}
	instr.Play()

	instr = &Piano{90}
	instr.Play()

	var amp Amplifier = &Guitar{7}
	amp.Connect("Ibanez")

	g := Guitar{12}
	g.Play()
	g.Connect("Marshall")
}
