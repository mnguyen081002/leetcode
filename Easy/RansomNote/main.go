package main

func canConstruct(ransomNote string, magazine string) bool {
	m := map[rune]int{}

	for _, v := range magazine {
		m[v]++
	}

	for _, v := range ransomNote {
		if cv, found := m[v]; !found || cv == 0 {
			return false
		}
		m[v]--
	}

	return true
}
