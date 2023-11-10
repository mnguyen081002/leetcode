package main

import (
	floodfill "leetcode/Easy/FloodFill"
	"math/rand"
	"testing"
	"time"
)

var num = 1000

func BenchmarkLoopNormal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LoopNormal(num)
	}
}

func BenchmarkLoopChannel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LoopChannel(num)
	}
}

var sr, sc, color = 1, 0, 2

var image = [][]int{}

func seedImage(m, n int) {
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 0, 1, and 2
	randomNum := rand.Intn(3)
	for i := 0; i < m; i++ {
		image = append(image, []int{})
		for j := 0; j < n; j++ {
			image[i] = append(image[i], randomNum)
		}
	}
}
func BenchmarkFloodFill(b *testing.B) {
	seedImage(10, 10)

	b.Run("FloodFillLoop", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			floodfill.FloodFillLoop(image, sr, sc, color)
		}
	})

	b.Run("FloodFillRecursive", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			floodfill.FloodFillRecursive(image, sr, sc, color)
		}
	})
}
