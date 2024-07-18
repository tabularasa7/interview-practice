package queue_test

import (
	"interview/practice/queue"
	"log"
	"testing"
)

func TestPush(t *testing.T) {
	var queue queue.Queue

	values := []int{2, 3, 4}

	for _, val := range values {
		queue.Push(val)
	}

	if queue.IsEmpty() || queue[len(queue)-1] != 4 {
		log.Fatalf("Pushing to queue failed")
	}
}

func TestPop(t *testing.T) {
	var queue queue.Queue

	values := []int{2, 3, 4}

	for _, val := range values {
		queue.Push(val)
	}

	lengthBefore := len(queue)
	returnVal := queue.Pop()
	lengthAfter := len(queue)

	if returnVal != 2 {
		log.Fatalf("Pop function did not return the first value")
	}

	if lengthBefore != lengthAfter+1 {
		log.Fatalf("Pop did not properly reduce the size of the queue after returning the first value")
	}
}
