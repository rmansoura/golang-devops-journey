package main

import (
	"fmt"
)

func add(a int, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func greet(name string) string {
	return "Hello, " + name + "!"
}

// Multiple return values
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

// Named return values
func computeStats(nums ...int) (sum, count int) {
	for _, n := range nums {
		sum += n
		count++
	}
	return
}

func main() {
	fmt.Println("=== Go Functions ===")

	// Simple functions
	result := add(10, 20)
	fmt.Printf("add(10, 20) = %d\n", result)

	result = subtract(50, 30)
	fmt.Printf("subtract(50, 30) = %d\n", result)

	greeting := greet("DevOps")
	fmt.Println(greeting)

	// Multiple return values
	fmt.Println("\n=== Multiple Returns ===")
	quotient, err := divide(100, 4)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("100 / 4 = %.2f\n", quotient)
	}

	_, err = divide(100, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Named returns
	fmt.Println("\n=== Named Returns ===")
	sum, count := computeStats(10, 20, 30, 40, 50)
	fmt.Printf("Sum: %d, Count: %d, Average: %.1f\n", sum, count, float64(sum)/float64(count))

	// Anonymous function
	fmt.Println("\n=== Anonymous Functions ===")
	square := func(x int) int {
		return x * x
	}
	fmt.Printf("Square of 7: %d\n", square(7))
}
