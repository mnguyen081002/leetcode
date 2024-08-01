package main

import (
	"math"
	"sort"
	"strconv"
)

func compress(chars []byte) int {
	writePointer, readPointer := 0, 0

	for readPointer < len(chars) {
		char := chars[readPointer]
		count := count(char, chars, readPointer)

		chars[writePointer] = char
		writePointer++
		writePointer = writeCount(chars, count, writePointer)

		readPointer += count
	}

	return writePointer
}

func count(char byte, chars []byte, readPointer int) int {
	c := 0
	for i := readPointer; i < len(chars); i++ {
		if char == chars[i] {
			c++
		} else {
			return c
		}
	}
	return c
}

func findMaxAverage(nums []int, k int) float64 {
	var m float64 = math.MinInt
	sum := 0
	for i := range nums {
		sum += nums[i]
		if i-k+1 >= 0 {
			m = max(m, float64(sum)/float64(k))
			sum -= nums[i-k+1]
		}
	}

	return m
}

func writeCount(chars []byte, count int, writePointer int) int {

	if count == 1 {
		return writePointer
	}

	s := strconv.Itoa(count)

	for i := 0; i < len(s); i++ {
		chars[writePointer] = byte(s[i])
		writePointer++
	}

	sort.Search(len(s), func(i int) bool {
		return s[i] == '0'
	})

	return writePointer
}
