package main

import "sort"

func isNStraightHand(hand []int, gs int) bool {
	// has to be evenly divisible by group size
	if len(hand)%gs != 0 {
		return false
	}
	// just each individual card
	if gs == 1 {
		return true
	}

	deck, cards := make(map[int]int), []int{}

	for _, card := range hand {
		// store cards with no duplicates
		if deck[card] == 0 {
			cards = append(cards, card)
		}
		deck[card]++
	}

	sort.Ints(cards)

	for _, card := range cards {
		// while the card exists check for straight
		for deck[card] > 0 {
			for i := 0; i < gs; i++ {
				if deck[card+i] > 0 {
					deck[card+i]--
				} else {
					return false
				}
			}
		}
	}

	return true
}

func main() {
	isNStraightHand([]int{1, 2, 3, 6, 2, 3, 4, 7, 8}, 3)
}
