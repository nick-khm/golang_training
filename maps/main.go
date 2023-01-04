package main

import "fmt"

type User struct {
	name string
}

type Key struct {
	ID   int
	Name string
}

func main() {
	m := map[string]*User{
		"HR":  {"Bob"},
		"CEO": {"Alice"},
	}

	fmt.Printf("m=%v\n", m)
	for key, value := range m {
		fmt.Printf("key=%v val=%v\n", key, value)
	}
	hr := m["HR"]
	hr.name = "TOTO"
	fmt.Printf("hr=%v\n", hr)
	fmt.Printf("m=%v\n", m)
	for key, value := range m {
		fmt.Printf("key=%v val=%v\n", key, value)
	}

	m2 := m
	m2["HR"] = &User{"Yolo"}
	fmt.Printf("m=%v\n", m)
	for key, value := range m {
		fmt.Printf("key=%v val=%v\n", key, value)
	}
	fmt.Printf("m2=%v\n", m2)
	for key, value := range m2 {
		fmt.Printf("key=%v val=%v\n", key, value)
	}

	resourceWeb := make(map[Key]string)
	resourceWeb[Key{1, "azerty"}] = "file1"

	k := Key{2, "toto"}
	resourceWeb[k] = "file2"

	delete(resourceWeb, Key{1, "azerty"})

	fmt.Println(resourceWeb)

}
