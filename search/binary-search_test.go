package search

import (
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"github.com/kamaal111/go-algorithms/utils"
)

var _ = DescribeTable("BinarySearch", func(input []int, expectedResult int) {
	var result int
	utils.TimeIt("binary search", func() { result = BinarySearch(input, 77) })

	Expect(result).To(Equal(expectedResult))
},
	Entry("target found", utils.NewSlice(0, 1_000, 1), 77),
	Entry("target not found", utils.NewSlice(0, 500, 2), -1),
)
