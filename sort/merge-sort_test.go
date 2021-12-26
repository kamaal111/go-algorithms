package sort

import (
	"math/rand"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"github.com/kamaal111/go-algorithms/utils"
)

var _ = Describe("MergeSort", func() {
	DescribeTable("Oneven slice", func(input []int, expectedResult []int) {
		var result []int
		utils.TimeIt("merge sort", func() { result = MergeSort(input) })

		Expect(len(expectedResult)).To(Equal(len(result)))
		Expect(result).To(Equal(expectedResult))
	},
		Entry("oneven", func() []int {
			input := utils.NewSlice(0, 9, 1)
			rand.Seed(time.Now().UnixNano())
			rand.Shuffle(len(input), func(i, j int) { input[i], input[j] = input[j], input[i] })
			return input
		}(), utils.NewSlice(0, 9, 1)),
		Entry("even", func() []int {
			input := utils.NewSlice(0, 10, 1)
			rand.Seed(time.Now().UnixNano())
			rand.Shuffle(len(input), func(i, j int) { input[i], input[j] = input[j], input[i] })
			return input
		}(), utils.NewSlice(0, 10, 1)),
	)
})
