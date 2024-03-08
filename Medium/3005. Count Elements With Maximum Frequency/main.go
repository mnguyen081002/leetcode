package main

func maxFrequencyElements(nums []int) int {
	countMax, countSame, m := 0, 0, map[int]int{}

	for _, v := range nums {
		m[v]++
		if m[v] > countMax {
			countMax = m[v]
			countSame = 0
		}
		if m[v] == countMax {
			countSame++
		}
	}

	return countMax * countSame
}
