def countingSort(arr, maxValue):
    b_len = maxValue + 1
    bucket = [0] * b_len
    sortedIndex = 0
    a_len = len(arr)
    
    for i in range(a_len):
        if not bucket[arr[i]]:
            bucket[arr[i]] = 0
        bucket[arr[i]] += 1

    for j  in range(b_len):
        while bucket[j] > 0:
            arr[sortedIndex] = j
            sortedIndex += 1
            bucket[j] -= 1

    return arr