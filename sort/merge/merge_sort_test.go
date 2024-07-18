package merge_test

import (
	"fmt"
	"interview/practice/sort/merge"
	"log"
	"testing"
)

func TestSort(t *testing.T) {
	arrayToSort := []int{2, 7, 11, 4, 1}

	sortedArray := merge.MergeSort(arrayToSort)
	fmt.Println(sortedArray)

	if sortedArray == nil || sortedArray[0] != 1 || sortedArray[len(sortedArray)-1] != 11 {
		log.Fatalf("Array was not sorted")
	}
}
