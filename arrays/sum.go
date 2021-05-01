package arrays

func Sum(numbers []int) int {
	sum := 0

	// range let me iterate over an array, returns an index and the value
	// Im ignoring the index by using the _ <- blank identifier

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	// lengthOfNumbers := len(numbersToSum)
	// sums := make([]int, lengthOfNumbers)
	//
	// for i, numbers := range numbersToSum {
	//   sums[i] = Sum(numbers)
	// }

	// This implementation allow us to not to worry about sizes
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
