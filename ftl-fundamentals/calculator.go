// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"fmt"
	"math"
)


// Add takes two numbers and returns the result of adding them together.
func AddOriginal(a, b float64) float64 {
	return a + b
}

// Add takes any number of numbers and returns the result of adding them together.
func Add(inputs ...float64) float64 {
	var total float64 = 0
	for _, input := range inputs {
		total += input
	}
	return total
}

	
func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a, b float64) float64 {
	return b - a
}

// Multiply takes two numbers and returns the result of multipluing them
func Multiply(a, b float64) float64 {
	return b * a
}

// Divide takes two numbers and returns the result of dividing them
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("bad inputs %f, %f: division by zero is not allowed", a, b)
	}
	return a / b, nil
}

func SqrRoot(a float64) (float64, error) {
	if a < 0 {
		return 0, fmt.Errorf("bad input %f: Sqaure Root of negative numbers is not allowed", a)
	}
	return math.Sqrt(a), nil
}
