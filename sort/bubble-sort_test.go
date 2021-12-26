package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestBubbleSort(t *testing.T) {
	input := NewSlice(0, 10, 1)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(input), func(i, j int) { input[i], input[j] = input[j], input[i] })

	start := time.Now()
	got := BubbleSort(input)
	elapsed := time.Since(start)
	fmt.Printf("done sorting in %s\n", elapsed)

	want := NewSlice(0, 10, 1)

	for wantIndex, wantItem := range want {
		if got[wantIndex] != wantItem {
			t.Errorf("got %d, wanted %d", got[wantIndex], wantItem)
		}
	}

}

// Returns a slice of elements with exact count.
// step will be used as the increment between elements in the sequence.
// step should be given as a positive, negative or zero number.
func NewSlice(start, count, step int) []int {
	s := make([]int, count)
	for i := range s {
		s[i] = start
		start += step
	}
	return s
}
