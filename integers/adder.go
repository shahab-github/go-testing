package integers

// Add takes two integers and returns the sum of them
func Add(x, y int) int {
	return x + y
}

// Sum takes integers and return thier sum
func Sum(numbers [5]int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}
