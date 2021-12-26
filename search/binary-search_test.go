package search

import (
	"testing"

	"github.com/kamaal111/go-algorithms/utils"
)

func TestBinarySearch(t *testing.T) {
	input := utils.NewSlice(0, 1_000, 1)

	var got int
	utils.TimeIt("binary search", func() { got = BinarySearch(input, 77) })

	want := 77

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestBinarySearchTargetNotFound(t *testing.T) {
	input := utils.NewSlice(0, 500, 2)

	var got int
	utils.TimeIt("binary search", func() { got = BinarySearch(input, 77) })

	want := -1

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
