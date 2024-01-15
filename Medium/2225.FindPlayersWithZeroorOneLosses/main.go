package main

func findWinners(matches [][]int) [][]int {
	scores := [100001]int{}

	for i := 0; i < len(matches); i++ {
		if scores[matches[i][0]] == 0 {
			scores[matches[i][0]] = 1
		}

		if scores[matches[i][1]] == 0 {
			scores[matches[i][1]] = 1
		}
		scores[matches[i][1]]++

	}

	zeros := make([]int, 0)
	ones := make([]int, 0)
	for i := 0; i < len(scores); i++ {
		switch scores[i] {
		case 1:
			zeros = append(zeros, i)
		case 2:
			ones = append(ones, i)
		}
	}
	return [][]int{zeros, ones}
}
