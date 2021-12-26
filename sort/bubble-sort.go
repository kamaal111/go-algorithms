package sort

func BubbleSort(input []int) []int {
	arr := input
	arrLength := len(arr)

	for i := 0; i < arrLength; i += 1 {
		swapped := false

		for j := 0; j < arrLength-i-1; j += 1 {
			first, second := arr[j], arr[j+1]
			if first > second {
				arr[j], arr[j+1] = second, first
				swapped = true
			}
		}

		if !swapped {
			return arr
		}
	}
	return arr
}
