package main

import "fmt"

func main() {
	fmt.Println("starting the debugging process")

	name := "debug"
	age := 18

	fmt.Println(name, age)

	if age > 18 {
		fmt.Println("eligible")
	} else if age > 10 {
		fmt.Println("can enter with parents")
	} else {
		fmt.Println("not eligible")
	}

	calculate := sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(calculate)

	fac := factorial(10)
	fmt.Println(fac)
}

func sum(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// factorial function calculates the factorial of a given number n
func factorial(n int) int {
	// Base case: factorial of 0 is 1
	if n == 0 {
		return 1
	}
	// Recursive case: n! = n * (n-1)!
	return n * factorial(n-1)
}
