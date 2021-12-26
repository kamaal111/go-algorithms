package sort

func BubbleSort(input []int) []int {
	arr := input
	for i := 0; i < len(arr); i += 1 {
		swapped := false

		for j := 0; j < len(arr)-i-1; j += 1 {
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
