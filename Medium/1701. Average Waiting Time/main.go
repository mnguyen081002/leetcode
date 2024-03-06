package main

func averageWaitingTime(customers [][]int) float64 {
	currentTime, wait := 0, 0
	for _, customer := range customers {
		if currentTime > customer[0] {
			wait += currentTime - customer[0] + customer[1]
		} else {
			currentTime = customer[0]
			wait += customer[1]
		}
		currentTime += customer[1]
	}
	return float64(wait) / float64(len(customers))
}
