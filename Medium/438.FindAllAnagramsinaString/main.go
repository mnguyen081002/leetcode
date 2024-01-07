package main

func findAnagrams(s string, p string) []int {
	l, h, result := 0, 0, []int{}
	k := [26]int{}
	for _, v := range p {
		k[v-'a']++
	}

	kf := [26]int{}
	for h < len(s) {
		kf[s[h]-'a']++
		if h-l == len(p) {
			kf[s[l]-'a']--
			l++
		}
		if k == kf {
			result = append(result, l)
		}
		h++
	}

	return result
}

func main() {
	s := "cbaebabacd"
	findAnagrams(s, "abc")
}
