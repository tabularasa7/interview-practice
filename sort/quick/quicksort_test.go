package quick_test

import (
	"fmt"
	"interview/practice/sort/quick"
	"testing"
)

func TestSort(t *testing.T) {
	arrayToSort := []int{5, 6, 2, 3, 9}

	quick.Quicksort(arrayToSort, 0, len(arrayToSort)-1)

	fmt.Println(arrayToSort)
}
