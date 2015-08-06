package main

import "fmt"

func main() {
	var a [2]int
	a[0] = 1 // a[1] will be auto-initialized to 0.

	for i := 0; i < len(a); i++ { 
		fmt.Println("Value of array at pos",  i,  "=>", a[i])
	}

	// Arrays have their place, but they're a bit inflexible. 
	// See http://blog.golang.org/go-slices-usage-and-internals
	// So we trying out his brother, slice -- an abstraction built on top of go's array type.

	// Defining a slice is just like an array literal, but leave out the element count like,
	slice := []string {"Sunday", "Monday", "Tuesday", "Wednessday", "Thursday", "Friday", "Saturday"} 
	for i := 0; i < len(slice); i++ {
		fmt.Println("Value of slice at pos", i, "=>", slice[i])
	}

	// Finally! the much awaited switch statement...
	day := 10
	switch day {
		case 0: fmt.Println("Sunday")
		case 1: fmt.Println("Monday")
		case 2: fmt.Println("Tuesday")
		case 3: fmt.Println("Wednessday")
		case 4: fmt.Println("Thursday")
		case 5: fmt.Println("Friday")
		case 6: fmt.Println("Saturday")

		default: fmt.Println("Whaaawt? I don't know what day it is =>", day)
	}

	// Without switch. Meh!
	day = 3
	fmt.Println(day, "is", slice[day])
}