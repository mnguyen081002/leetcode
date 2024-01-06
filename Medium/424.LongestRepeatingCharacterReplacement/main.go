package main

func characterReplacement(s string, k int) int {
	charCounts := make([]int, 26) // only uppercase letters

	lo := 0
	longestCnt := 0
	maxLen := 0

	// the sub-string under consideration is bound between hi and lo indices/'pointers'
	//   hi advances the window  by one char and lo shrinks it by one char
	for hi := 0; hi < len(s); hi++ {
		c := s[hi] - 'A'
		charCounts[c]++
		// keep track of the max length of contiguous chars in a separate variable.
		longestCnt = max(longestCnt, charCounts[c])

		// see if you can 'flip' characters on either side of this contiguous segment
		//   Ofc, u are allowed to flip at most k chars.
		//   That means, we have to omit/ignore anything greater than longestCnt + k.
		//   So we need to 'stop' processing the current sub-string and move on to the next sub-string
		//          we do so by advancing lo and decrementing char count @ lo
		if longestCnt+k < hi-lo+1 {
			charCounts[s[lo]-'A']--
			lo++
		}

		// keep track of the max. by the time we are here, the sub-string has at most longestCnt + k length
		maxLen = max(maxLen, hi-lo+1)
	}
	return maxLen
}
