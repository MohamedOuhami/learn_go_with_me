package main

import (
	"fmt"
	"strconv"

	"rsc.io/quote/v4"
)

func main() {
	fmt.Println("Hello, world")
	fmt.Println(quote.Go())

	// In Go, we can instantly interfer the type of a variable
	var a, b int = 4, 5
	println(a, b)

	// := is a shorthand for declaring and initialzing a variable
	c := 9
	println("c is shorthanded declared : " + strconv.Itoa(c))

	// To declare consts
	const EMBLEM = "Dima Raja"
	println(EMBLEM)

	// For loops in Go
	for i := 1; i < 3; i++ {
		println(i)
	}

	// We have range in Go
	for i := range 6 {
		println(i)
	}

	// Conditionals
	if 7%2 == 0 {
		println("7 is even")
	} else {
		println("7 is odd")
	}

	// We can even take a value and get It inside the conditionals
	if num := 9; num < 0 {
		println("The num is negative")
	} else if num > 0 {
		println("The num is positive")
	} else {
		println("The num is nil")
	}

	// We can also use switch
	whatAmI := func(i interface{}) {
		switch i.(type) {
		case bool:
			println("THe var is a bool")
		case int:
			println("THis var is an int")
		default:
			println("Other types")
		}
	}

	whatAmI(true)
	whatAmI(2)
	whatAmI("Dima Raja")

	// Handling arrays

	Learningarrays()
}
