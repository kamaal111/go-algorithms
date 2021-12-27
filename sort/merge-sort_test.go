package sort

import (
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	. "github.com/kamaal111/go-algorithms/utils"
)

var _ = DescribeTable("MergeSort", func(input []int, expectedResult []int) {
	var result []int
	TimeIt("merge sort", func() { result = MergeSort(input) })

	Expect(len(expectedResult)).To(Equal(len(result)))
	Expect(result).To(Equal(expectedResult))
},
	Entry("oneven", ShuffleSlice(NewSlice(0, 9, 1)), NewSlice(0, 9, 1)),
	Entry("even", ShuffleSlice(NewSlice(0, 10, 1)), NewSlice(0, 10, 1)),
)
