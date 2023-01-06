package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Cooker interface {
	Cook()
}

type Chef struct {
	Person
	Stars int
}

func (c Chef) Cook() {
	fmt.Printf("I'm cooking with %v start\n", c.Stars)
}

func main() {
	var x interface{} = Person{"Bob", 10}
	fmt.Printf("x type=%T, data=%v\n", x, x)

	// always need to test with "ok" while casting variable
	s, ok := x.(string)
	fmt.Printf("Person as string ok? %v. value='%v'\n", ok, s)

	i, ok := x.(int)
	fmt.Printf("Person as int ok? %v. value='%v'\n", ok, i)

	processPerson(x)

	x = Chef{
		Person: Person{"Toto", 33},
		Stars:  5,
	}

	processPerson(x)
	processPerson(5)
	processPerson("John")
}

func processPerson(i interface{}) {
	switch p := i.(type) {
	case Person:
		fmt.Printf("We have a person=%v\n", p)
	case Chef:
		fmt.Printf("We have a chef=%v, let him cook\n", p)
		p.Cook()
	default:
		fmt.Printf("Unknown type=%T, value=%v\n", p, p)
	}
}
