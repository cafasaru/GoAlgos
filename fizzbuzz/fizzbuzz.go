package fizzbuzz

import (
	"fmt"
)

func FizzBuzz(n int) {
	for i := 1; i < n; i++ {
		printFizzBuzz(i)
		fmt.Print(", ")
	}
	printFizzBuzz(n)
	fmt.Println()
}

func printFizzBuzz(i int) {
	switch {
	case i%3 == 0 && i%5 == 0:
		fmt.Print("FizzBuzz")
	case i%3 == 0:
		fmt.Print("Fizz")
	case i%5 == 0:
		fmt.Print("Buzz")
	default:
		fmt.Print(i)
	}
}
