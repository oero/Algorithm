package countingSort

func countingSort(arr []int, maxValue int) []int {
	b_len := maxValue + 1
	bucket := make([]int, b_len)

	sortedIndex := 0
	length := len(arr)

	for i := 0; i < length; i++ {
		bucket[arr[i]] += 1
	}

	for j := 0; j < b_len; j++ {
		for bucket[j] > 0 {
			arr[sortedIndex] = j
			sortedIndex += 1
			bucket[j] -= 1
		}
	}

	return arr
}
