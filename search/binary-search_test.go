package search

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	input := NewSlice(0, 1_000_000, 1)
	got := BinarySearch(input, 7777)
	want := 7777

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestBinarySearchTargetNotFound(t *testing.T) {
	input := NewSlice(0, 1_000_000, 2)
	got := BinarySearch(input, 7777)
	want := -1

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
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
