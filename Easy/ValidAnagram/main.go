package main

func isAnagram(s string, t string) bool {
	m := map[rune]int{}
	for _, v := range s {
		m[v]++
	}

	for _, v := range t {
		if vf, found := m[v]; !found || vf == 0 {
			return false
		}
		m[v]--
	}

	return len(s) == len(t)
}
