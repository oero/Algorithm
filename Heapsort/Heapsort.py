def buildMaxHeap(arr):
    import math
    for i in range(math.floor(len(arr)/2), -1, -1):
        heapify(arr, i)

def heapify(arr, i):
    left = 2 * i + 1
    right = 2 * i + 2
    largest = i

    if left < arrlen and arr[left] > arr[largest]:
        largest = left 

    if right < arrlen and arr[right] > arr[largest]:
        largest = right

    if largest != i:
        swap(arr, i, largest)
        heapify(arr, largest)

def swap(arr, i, j):
    arr[i], arr[j] = arr[j], arr[i]

def heapSort(arr):
    global arrlen
    arrlen = len(arr)
    buildMaxHeap(arr)
    for i in range(len(arr)-1, 0, -1):
        swap(arr, 0, i)
        arrlen -= 1
        heapify(arr, 0)
    return arr