package main

import "fmt"

// This is the main function
func main() {

	// Declare a variable 'i' of type int and print its value
	var i int
	fmt.Println("Value of i:", i)

	// Declare and initialize a variable 'j' with the value 10
	var j = 10
	fmt.Println("Value of j:", j)

	// Declare and initialize a variable 'k' with the value true
	var k = true
	fmt.Println("Value of k:", k)

	// Use short variable declaration to declare and initialize 'm' with the value "Hello, Go"
	m := "Hello, Go"
	fmt.Println("Value of m:", m)

	// Declare and initialize multiple variables 'x' and 'y'
	var x, y = 1, 2
	fmt.Println("Values of x and y:", x, y)

	// Use short variable declaration to declare and initialize multiple variables 'p', 'q' and 'r'
	p, q, r := 7, false, "hi"
	fmt.Println("Values of p, q and r:", p, q, r)
}
