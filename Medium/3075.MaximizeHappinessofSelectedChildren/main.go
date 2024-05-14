package main

import "slices"

func maximumHappinessSum(happiness []int, k int) int64 {
	slices.SortFunc(happiness, func(i, j int) int {
		return j - i
	})
	des := 0
	var ans int64
	for i := 0; i < k; i++ {
		happiness[i] -= des
		happiness[i] = max(happiness[i], 0)
		ans += int64(happiness[i])
		des++
	}
	return ans
}

func main() {
	happiness := []int{1, 2, 3, 4, 5}
	k := 2
	println(maximumHappinessSum(happiness, k))
}
