package main

// TODO: Write a function called SumArray that takes a slice of integers and returns their sum
func SumArray(numbers []int) int {
	// Your code here
	if len(numbers) == 0 {
		return 0
	}

	return numbers[0] + SumArray(numbers[1:])
}
