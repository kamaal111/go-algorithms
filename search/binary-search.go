package search

func BinarySearch(input []int, target int) int {
	lowerBound := 0
	upperBound := len(input) - 1

	for lowerBound <= upperBound {
		middleIndex := (lowerBound + upperBound) / 2
		middleValue := input[middleIndex]

		if middleValue == target {
			return middleIndex
		}

		if middleValue < target {
			lowerBound = middleIndex + 1
		} else {
			upperBound = middleIndex - 1
		}
	}

	return -1
}
