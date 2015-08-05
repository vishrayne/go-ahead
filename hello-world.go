package main

import "fmt"

func main() {
	fmt.Println("hello world")

	var a, b string = "gostring", "anothergostring"
	fmt.Println("this is a " + a)
	fmt.Println("this also is a " + b)

	// short-cut syntax
	d := "finalstring"
	fmt.Println("Last one, and I'm " + d)

	// A constant, and weird its there in place of var :|
	const s string = "CONSTANT STRING"
	fmt.Println("One more time, but this time a constant " + s)
}