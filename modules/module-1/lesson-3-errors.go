package main

import (
	"errors"
	"fmt"
)

// Custom error
var ErrInvalidInput = errors.New("invalid input provided")

// Function that returns error
func readConfig(filename string) (string, error) {
	if filename == "" {
		return "", ErrInvalidInput
	}
	return "config data", nil
}

// Wrapping errors
func processDeployment(env string) error {
	if env != "production" && env != "staging" && env != "development" {
		return fmt.Errorf("invalid environment: %s", env)
	}
	fmt.Printf("Processing deployment to %s\n", env)
	return nil
}

// Panic recovery
func safeDevide(a, b int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	if b == 0 {
		panic("Division by zero!")
	}
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
}

func main() {
	fmt.Println("=== Error Handling in Go ===")

	// Checking errors
	fmt.Println("\n=== Check Errors ===")
	config, err := readConfig("app.json")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Config:", config)
	}

	// Error with empty filename
	_, err = readConfig("")
	if err != nil {
		fmt.Println("Error reading config:", err)
	}

	// Formatted errors
	fmt.Println("\n=== Formatted Errors ===")
	err = processDeployment("testing")
	if err != nil {
		fmt.Println("Deployment error:", err)
	}

	err = processDeployment("production")
	if err != nil {
		fmt.Println("Deployment error:", err)
	}

	// Panic and recover
	fmt.Println("\n=== Panic and Recover ===")
	safeDevide(100, 5)
	safeDevide(100, 0)
	fmt.Println("Program continues after panic recovery")

	// Error type assertion
	fmt.Println("\n=== Error Type Assertion ===")
	err = fmt.Errorf("wrapped: %w", ErrInvalidInput)
	if errors.Is(err, ErrInvalidInput) {
		fmt.Println("Error is ErrInvalidInput")
	}
}
