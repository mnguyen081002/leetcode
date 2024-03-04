package main

func numFriendRequests(ages []int) int {
	ans := 0
	count := make([]int, 121)
	for _, age := range ages {
		count[age]++
	}

	for a := 0; a < 121; a++ {
		countA := count[a]
		if countA == 0 {
			continue
		}
		for b := 0; b < 121; b++ {
			countB := count[b]
			if countB == 0 {
				continue
			}
			if b <= (a/2 + 7) {
				continue
			}
			if b > a {
				continue
			}
			if b > 100 && a < 100 {
				continue
			}

			ans += (countA * countB)
			if a == b {
				ans -= countA
			}
		}
	}

	return ans
}

func resultArray(nums []int) []int {
	a1, a2 := []int{nums[0]}, []int{nums[1]}

	for i := 2; i < len(nums); i++ {
		if a1[len(a1)-1] > a2[len(a2)-1] {
			a1 = append(a1, nums[i])
		} else {
			a2 = append(a2, nums[i])
		}
	}
	return append(a1, a2...)
}

func countSubmatrices(grid [][]int, k int) int {

}

func subarraySum(nums []int, k int) int {
	count := map[int]int{0: 1}
	pre := 0
	ans := 0
	for _, num := range nums {
		pre += num
		if _, ok := count[pre-k]; ok {
			ans += count[pre-k]
		}
		count[pre]++
	}
	return ans
}
