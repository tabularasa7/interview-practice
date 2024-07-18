package binary

import "fmt"

func Search(numArray []int, val int) int {

	if len(numArray) == 0 {
		return -1
	}
	fmt.Printf("NumArray: %v\n", numArray)
	if len(numArray) == 1 {
		return numArray[0]
	}

	mid := len(numArray) / 2

	if val == numArray[mid] {
		return numArray[mid]
	} else if val < numArray[mid] {
		return Search(numArray[0:mid], val)
	} else {
		return Search(numArray[mid+1:], val)
	}
}
