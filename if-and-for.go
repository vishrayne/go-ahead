package main

import "fmt"

func main() {
	// a simple form
	i := 1
	for i <= 3 {
		i += 1
		// concatenating an int and a string.
		fmt.Println(i, "plus a string")
	}

	// classic for loop
	for j := i; j >= 0; j-- {
		fmt.Println(j)
	}

	// ask for `j` here and you get an undefined var error.
	// so for is the only loop in this lang. Also testing if loop.
	for {
		// another concat, and yup order doesn't matter.
		fmt.Println("Infinity and beyond...", i)

		i--
		if i == 0 {
			fmt.Println("Just joking!")
			break;
		}
	}
}