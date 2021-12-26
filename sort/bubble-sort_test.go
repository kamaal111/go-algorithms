package sort

import (
	"math/rand"
	"testing"
	"time"

	"github.com/kamaal111/go-algorithms/utils"
)

func TestBubbleSort(t *testing.T) {
	input := utils.NewSlice(0, 10, 1)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(input), func(i, j int) { input[i], input[j] = input[j], input[i] })

	var got []int
	utils.TimeIt("bubble sort", func() { got = BubbleSort(input) })

	want := utils.NewSlice(0, 10, 1)

	for wantIndex, wantItem := range want {
		if got[wantIndex] != wantItem {
			t.Errorf("got %d, wanted %d", got[wantIndex], wantItem)
		}
	}

}