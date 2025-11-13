package main

import (
	"fmt"
)

func main() {
	// Variables in Go
	var name string = "DevOps Engineer"
	var age int = 30
	var salary float64 = 120000.50

	// Short declaration (inside functions only)
	country := "United States"
	isLearning := true

	fmt.Println("=== Go Variables ===")
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Salary:", salary)
	fmt.Println("Country:", country)
	fmt.Println("Currently Learning:", isLearning)

	// Constants
	const maxServers = 100
	const appVersion = "1.0.0"

	fmt.Println("\n=== Constants ===")
	fmt.Println("Max Servers:", maxServers)
	fmt.Println("App Version:", appVersion)

	// Type inference
	role := "DevOps"
	containers := 42
	uptime := 99.95

	fmt.Println("\n=== Type Inference ===")
	fmt.Printf("Role type: %T, value: %v\n", role, role)
	fmt.Printf("Containers type: %T, value: %v\n", containers, containers)
	fmt.Printf("Uptime type: %T, value: %v\n", uptime, uptime)
}
