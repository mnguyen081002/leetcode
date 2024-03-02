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
