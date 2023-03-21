package main

import (
	"fmt"
	"math/rand"
)

const (
	MaxRandom  = 100
	N          = 100
	Iterations = 100
)

func Rand() int {
	return rand.Intn(MaxRandom)
}

func uniqRandN(n int) []int {
	uniqMap := make(map[int]bool)
	for len(uniqMap) < n {
		uniqMap[Rand()] = true
	}
	uniqSlice := make([]int, 0, len(uniqMap))
	for k := range uniqMap {
		uniqSlice = append(uniqSlice, k)
	}
	return uniqSlice
}

func main() {
	fmt.Print(uniqRandN(N))
	for i := 0; i < Iterations; i++ { // just to be sure if it's generating right amount of numbers in each slice
		fmt.Println(len(uniqRandN(N)))
	}
}
