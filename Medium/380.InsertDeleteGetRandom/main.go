package main

import (
	"math/rand"
)

type RandomizedSet struct {
	m   map[int]int
	arr []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		m:   map[int]int{},
		arr: []int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, found := this.m[val]; found {
		return false
	}
	this.arr = append(this.arr, val)
	this.m[val] = len(this.arr) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if index, found := this.m[val]; found {
		this.arr[index] = this.arr[len(this.arr)-1]
		this.arr = this.arr[:len(this.arr)-1]
		delete(this.m, val)
		if index < len(this.arr) {
			this.m[this.arr[index]] = index
		}
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	return this.arr[rand.Intn(len(this.arr))]
}
