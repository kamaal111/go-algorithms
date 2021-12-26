package sort

func MergeSort(input []int) []int {
	inputLength := len(input)
	if inputLength < 2 {
		return input
	}

	middleIndex := inputLength / 2
	left := MergeSort(input[:middleIndex])
	right := MergeSort(input[middleIndex:])
	return merge(left, right)
}

func merge(left []int, right []int) []int {
	arr := []int{}
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			arr = append(arr, left[i])
			i += 1
		} else {
			arr = append(arr, right[j])
			j += 1
		}
	}

	for ; i < len(left); i += 1 {
		arr = append(arr, left[i])
	}
	for ; j < len(right); j += 1 {
		arr = append(arr, right[j])
	}

	return arr
}
