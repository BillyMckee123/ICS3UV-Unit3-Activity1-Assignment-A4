package main
// Assignment A4 - Average of Three Numbers
// Calculates the average of 56.9, 89.7, 90.2

import "fmt"

func main() {
	num1 := 56.9
	num2 := 89.7
	num3 := 90.2

	sum := num1 + num2 + num3
	average := sum / 3

	fmt.Printf("The average of %.1f, %.1f, and %.1f is %.14f.\n", num1, num2, num3, average)
}
