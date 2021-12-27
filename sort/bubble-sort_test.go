package sort

import (
	"math/rand"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kamaal111/go-algorithms/utils"
)

var _ = Describe("BubbleSort", func() {
	It("successfully sorts slice", func() {
		input := NewSlice(0, 10, 1)
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(input), func(i, j int) { input[i], input[j] = input[j], input[i] })

		var result []int
		TimeIt("bubble sort", func() { result = BubbleSort(input) })

		expectedResult := NewSlice(0, 10, 1)

		Expect(len(expectedResult)).To(Equal(len(result)))
		Expect(expectedResult).To(Equal(result))
	})
})
