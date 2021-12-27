package sort

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kamaal111/go-algorithms/utils"
)

var _ = Describe("BubbleSort", func() {
	It("successfully sorts slice", func() {
		input := ShuffleSlice(NewSlice(0, 10, 1))

		var result []int
		TimeIt("bubble sort", func() { result = BubbleSort(input) })

		expectedResult := NewSlice(0, 10, 1)

		Expect(len(expectedResult)).To(Equal(len(result)))
		Expect(expectedResult).To(Equal(result))
	})
})
