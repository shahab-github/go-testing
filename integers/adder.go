package integers

// Add takes two integers and returns the sum of them
func Add(x, y int) int {
	return x + y
}

// Sum takes integers and return thier sum
func Sum(numbers []int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}

// func SumAll(numbersToSum ...[]int) []int {
// 	lengthOfNumbers := len(numbersToSum)
// 	sums := make([]int, lengthOfNumbers)

// 	for i, numbers := range numbersToSum {
// 		sums[i] = Sum(numbers)
// 	}
// 	return sums
// }

func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}
