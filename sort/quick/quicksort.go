package quick

import "fmt"

func Partition(array []int, low, high int) ([]int, int) {

	pivot := array[high]

	i := low

	for j := low; j < high; j++ {
		if array[j] < pivot {
			array[i], array[j] = array[j], array[i]
			i++
		}
	}

	fmt.Println(i)
	array[i], array[high] = array[high], array[i]

	return array, i
}

func Quicksort(array []int, low, high int) []int {

	if low < high {
		var pivot int
		array, pivot = Partition(array, low, high)
		array = Quicksort(array, low, pivot-1)
		array = Quicksort(array, pivot+1, high)
	}

	return array
}
