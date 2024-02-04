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
}

func sum(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
