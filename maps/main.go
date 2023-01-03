package main

import "fmt"

func main() {
	m := map[string]int{
		"Bob":     11,
		"Alice":   22,
		"Bobette": 33,
	}
	fmt.Printf("m=%v\n", m)

	for name, age := range m {
		fmt.Printf("name=%s age=%v\n", name, age)
	}

	i := 1
	for name := range m {
		m[name] = i
		i++
	}
	fmt.Printf("m=%v\n", m)

}
