package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct {
	color string
}

func (d Dog) Speak() string {
	return "Dog saying"
}

type Cat struct {
	jumpHeight int
}

func (c Cat) Speak() string {
	return "Cat saying"
}

func main() {
	animals := []Animal{
		Dog{"white"},
		Cat{2},
	}

	for _, animal := range animals {
		fmt.Println(animal.Speak())
		fmt.Printf("Type of animal %T\n", animal)

		// type assection
		// instType, ok := animal.(Dog)
		// fmt.Printf("Type assection value=%v, ok=%v\n", instType, ok)
		// == THE SAME THING BELOW
		if instType, ok := animal.(Dog); ok {
			fmt.Printf("We have a dog! color=%v\n", instType.color)
		} else {
			fmt.Println("It's not a dog here...")
		}

	}

	fmt.Println("------------------")

	for _, animal := range animals {
		describeAnimal(animal)
	}
}

func describeAnimal(animal Animal) {
	switch v := animal.(type) {
	case Dog:
		fmt.Printf("This is a dog; color=%v\n", v.color)
	case Cat:
		fmt.Printf("This is a cat; jumpHeight=%v\n", v.jumpHeight)
	}
}
