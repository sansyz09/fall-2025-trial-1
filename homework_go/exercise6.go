package main

// TODO: Write a function called Factorial that calculates the factorial of a positive integer
// Factorial of n is n! = n × (n-1) × (n-2) × ... × 1
// Factorial of 0 is 1
func Factorial(n int) int {
	// Your code here
	if n == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}
