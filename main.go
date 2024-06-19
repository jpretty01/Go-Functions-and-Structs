package main

import "fmt"

// Function
func add(x int, y int) int {
	return x + y
}

// Struct
type Person struct {
	Name string
	Age  int
}

func main() {
	result := add(10, 20)
	fmt.Println("Result:", result)

	// Struct Initialization
	person := Person{Name: "John", Age: 30}
	fmt.Printf("Name: %s, Age: %d\n", person.Name, person.Age)
}
