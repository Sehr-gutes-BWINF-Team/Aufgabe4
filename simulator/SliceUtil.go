package simulator

// CalculateAverage of a slice full of integers.
func CalculateAverage(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total = total + number
	}
	average := total / len(numbers) // len  function return array size
	return average
}

// Remove an order at a specific index.
func Remove(slice []Order, s int) []Order {
	return append(slice[:s], slice[s+1:]...)
}
