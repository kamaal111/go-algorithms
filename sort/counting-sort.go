package sort

func CountingSort(input []int) []int {
	length := len(input)
	lookup := make(map[int]int)
	count := 0
	for i := 0; i < length; i += 1 {
		lookup[input[i]] += 1
		count += 1
	}

	arr := []int{}
	for i := 0; i < len(lookup); i += 1 {
		currentLookup := lookup[i]
		for j := 0; j < currentLookup; j += 1 {
			arr = append(arr, i)
		}
	}
	return arr
}
