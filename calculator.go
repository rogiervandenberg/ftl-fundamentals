// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"fmt"
	"math"
)

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply takes 2 numbers and returns the result of the multiplication
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide takes 2 numbers and returns the division of those numbers
func Divide(a, b float64) (float64, error) {
	if b == 0 {

		return 0, fmt.Errorf("Bad input: %f, %f (division by zero)", a, b)
	}

	return a / b, nil
}

// Squared takes a number and tries to get the square root from it
func Squared(a float64) (float64, error) {

	if a < 0 {

		return 0, fmt.Errorf("Bad input: %f, input cannot be nagative", a)
	}

	return math.Sqrt(a), nil
}
