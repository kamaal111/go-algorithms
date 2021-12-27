package sort

func SelectionSort(input []int) []int {
	arr := input
	arrLength := len(arr)

	for i := 0; i < arrLength; i += 1 {
		lowestIndex := i
		lowestValue := arr[i]

		for j := i + 1; j < arrLength; j += 1 {
			comparisonValue := arr[j]
			if lowestValue > comparisonValue {
				lowestValue = comparisonValue
				lowestIndex = j
			}
		}

		arr[i], arr[lowestIndex] = arr[lowestIndex], arr[i]
	}

	return arr
}
