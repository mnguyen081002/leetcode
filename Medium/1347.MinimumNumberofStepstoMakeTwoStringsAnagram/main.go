package main

func minSteps(s string, t string) (res int) {
	dict := [26]int{}
	for k, v := range s {
		dict[v-'a']++
		dict[t[k]-'a']--
	}

	for _, v := range dict {
		if v > 0 {
			res += v
		}
	}
	return res
}
