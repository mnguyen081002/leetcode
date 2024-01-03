package main

func findMaxLength(nums []int) int {
	m := map[int]int{
		0: -1,
	}
	length, s := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			s++
		} else {
			s--
		}
		if v, found := m[s]; found {
			length = max(length, i-v)
		} else {
			m[s] = i
		}
	}

	return length
}
