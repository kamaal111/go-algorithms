package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// Returns a slice of elements with exact count.
// step will be used as the increment between elements in the sequence.
// step should be given as a positive, negative or zero number.
func NewSlice(start, count, step int) []int {
	s := make([]int, count/step)
	for i := range s {
		s[i] = start
		start += step
	}
	return s
}

func ShuffleSlice(input []int) []int {
	shuffledInput := input
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(shuffledInput), func(i, j int) { shuffledInput[i], shuffledInput[j] = shuffledInput[j], shuffledInput[i] })
	return shuffledInput
}

func TimeIt(context string, f func()) {
	start := time.Now()
	f()
	elapsed := time.Since(start)
	fmt.Printf("%s in %s\n", context, elapsed)
}
