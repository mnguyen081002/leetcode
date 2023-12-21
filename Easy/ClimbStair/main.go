package main

import (
	"fmt"
	"math"
)

func fibonacci(n int) int {
	phi := (1 + math.Sqrt(5)) / 2
	// psi := (1 - math.Sqrt(5)) / 2

	result := (math.Pow(phi, float64(n)) - math.Pow(1-phi, float64(n))) / math.Sqrt(5)
	return int(math.Round(result))
}

func main() {
	fmt.Println(fibonacci(9 + 1))
}
