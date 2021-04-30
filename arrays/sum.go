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
