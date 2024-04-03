package main

import (
	"fmt"
)

func minLevelsToGainMorePoints(possible []int) int {
	n := len(possible)
	danielPoints := 0
	bobPoints := 0
	minLevels := n + 1

	for i := 0; i < n; i++ {
		if possible[i] == 1 {
			danielPoints++
		}
		for j := i + 1; j < n; j++ {
			if possible[j] == 1 {
				bobPoints++
			}
		}
		if danielPoints > bobPoints {
			minLevels = min(minLevels, i+1)
		}
	}

	if minLevels == n+1 {
		return -1
	}
	return minLevels
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	possible := []int{1, 0, 1, 0}
	fmt.Println(minLevelsToGainMorePoints(possible)) // Output: 1
}
