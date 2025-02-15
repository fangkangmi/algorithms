package main

import "fmt"

// Factorial function
// factorial(n) = n * factorial(n-1) with base case factorial(0) = 1
func factorial(n int) int {
	// TODO: Implement the factorial function here
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

// Fibonacci function
// fib(n) = fib(n-1) + fib(n-2) with base cases fib(0) = 0 and fib(1) = 1
func fibonacci(n int) int {
	// TODO: Implement the fibonacci function here
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	// Test the factorial function
	fmt.Println("Testing Factorial:")
	testCasesFactorial := []int{0, 1, 5, 10}
	for _, n := range testCasesFactorial {
		fmt.Printf("factorial(%d) = %d\n", n, factorial(n))
	}

	fmt.Println("\nTesting Fibonacci:")
	// Test the fibonacci function
	testCasesFibonacci := []int{0, 1, 5, 10}
	for _, n := range testCasesFibonacci {
		fmt.Printf("fibonacci(%d) = %d\n", n, fibonacci(n))
	}
}
