package main

func uniqueOccurrences(arr []int) bool {
	m := map[int]int{}

	for _, v := range arr {
		m[v]++
	}
	mr := map[int]struct{}{}
	for _, v := range m {
		if _, found := mr[v]; found {
			return false
		}
		mr[v] = struct{}{}
	}
	return true
}
