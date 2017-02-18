package heapSort

func heapSort(arr []int) []int {
	arrlen := len(arr)
	buildMaxHeap(arr, arrlen)
	for i := arrlen - 1; i >= 0; i-- {
		swap(arr, 0, i)
		arrlen -= 1
		heapify(arr, 0, arrlen)
	}
	return arr
}

func buildMaxHeap(arr []int, arrlen int) {
	for i := arrlen / 2; i >= 0; i-- {
		heapify(arr, i, arrlen)
	}
}

func heapify(arr []int, i, arrlen int) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i

	if left < arrlen && arr[left] > arr[largest] {
		largest = left
	}

	if right < arrlen && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		swap(arr, i, largest)
		heapify(arr, largest, arrlen)
	}
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
