package merge

import "fmt"

func Merge(left, right []int) []int {
	sortedArray := []int{}
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		leftVal := left[i]
		rightVal := right[j]
		fmt.Println(i, leftVal, j, rightVal)
		if leftVal < rightVal {
			sortedArray = append(sortedArray, leftVal)
			i++
		} else {
			sortedArray = append(sortedArray, rightVal)
			j++
		}
	}

	for ; i < len(left); i++ {
		sortedArray = append(sortedArray, left[i])
	}

	for ; j < len(right); j++ {
		sortedArray = append(sortedArray, right[j])
	}

	return sortedArray
}

func MergeSort(dataArray []int) []int {
	if len(dataArray) == 1 {
		return dataArray
	}

	mid := len(dataArray) / 2

	left := dataArray[:mid]
	right := dataArray[mid:]

	sortedLeft := MergeSort(left)
	sortedRight := MergeSort(right)

	// fmt.Println(sortedLeft, sortedRight)
	return Merge(sortedLeft, sortedRight)
}
