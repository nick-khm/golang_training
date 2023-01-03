package main

import "fmt"

func main() {
	m := make(map[string]int)
	fmt.Printf("Map content %v (len=%v)\n", m, len(m))

	m["hello"] = 5
	m["goodbye"] = 10
	fmt.Printf("Map content %v (len=%v)\n", m, len(m))

	fmt.Printf("key=value, value=%v\n", m["hello"])

	i := m["goodbye"]
	fmt.Println(i)

	// error on key => init à default value and test key presence
	j, ok := m["hel"]
	fmt.Printf("j=%v and existance of key OK=%v\n", j, ok)

	// test online
	m["beatles"] = 2
	if _, testOk := m["beatles"]; testOk {
		fmt.Println("Beatles is OK")
	}

	delete(m, "beatles")
	fmt.Printf("Map content %v (len=%v)\n", m, len(m))

	m2 := m
	fmt.Printf("m2 content %v (len=%v)\n", m2, len(m2))
	m2["help"] = 44
	fmt.Printf("Map content %v (len=%v)\n", m, len(m))
	fmt.Printf("m2 content %v (len=%v)\n", m2, len(m2))

}
