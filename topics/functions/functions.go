package main

import "fmt"

func greet() {
	fmt.Println("Hello from greet function")
}

func double(x int) int {
	return x + x
}

func add(lhs, rhs int) int {
	return lhs + rhs
}

func main() {
	greet()

	dozen := double(6)
	fmt.Println("The dozen is ", dozen)

	bakerDozen := add(dozen, 1)
	fmt.Println("the bakersDozen is ", bakerDozen)

	anotherBakerDozens := add(double(6), 1)
	fmt.Println("Another baker dozens is", anotherBakerDozens)
}
