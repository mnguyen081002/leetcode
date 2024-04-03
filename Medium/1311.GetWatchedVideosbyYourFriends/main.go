package main

import (
	"bytes"
	"math"
	"slices"
)

func watchedVideosByFriends(watchedVideos [][]string, friends [][]int, id int, level int) []string {
	visit := make([]bool, len(friends))
	q := []int{id}
	lev := math.MaxInt
	freq := map[string]int{}
	for len(q) > 0 {
		for _, f := range q {
			q = q[1:]
			if visit[f] {
				continue
			}
			visit[f] = true

			if id == f {
				lev = level
			}

			if lev == 0 {
				for _, v := range watchedVideos[f] {
					freq[v]++
				}
			}

			q = append(q, friends[f]...)
		}
		lev--
	}
	res := []string{}
	for k := range freq {
		res = append(res, k)
	}

	slices.SortFunc(res, func(x, y string) int {
		if freq[x] == freq[y] {
			return bytes.Compare([]byte(x), []byte(y))
		}
		return freq[x] - freq[y]
	})
	return res
}

func main() {
	watchedVideos := [][]string{
		{"m"}, {"xaqhyop", "lhvh"},
	}

	friends := [][]int{
		{1}, {0},
	}

	id := 1
	level := 1

	watchedVideosByFriends(watchedVideos, friends, id, level)
}
