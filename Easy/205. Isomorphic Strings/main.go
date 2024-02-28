package main

func isIsomorphic(s string, t string) bool {
	from := make([]byte, 256)
	to := make([]byte, 256)

	for i := range s {
		if from[s[i]] == 0 {

			if to[t[i]] != 0 {
				return false
			}
			from[s[i]] = t[i]
			to[t[i]] = s[i]
			continue
		}

		if from[s[i]] != t[i] {
			return false
		}
	}

	return true
}
