package stack_test

import (
	"interview/practice/stack"
	"log"
	"testing"
)

func TestPush(t *testing.T) {
	var stack stack.Stack

	values := []int{2, 3, 4}

	for _, val := range values {
		stack.Push(val)
	}

	if stack.IsEmpty() {
		log.Fatalf("Stack should not be empty")
	}

	if stack[len(stack)-1] != values[len(values)-1] {
		log.Fatalf("Last value in stack is %d, not %d\n", stack[len(stack)-1], values[len(values)-1])
	}
}

func TestPop(t *testing.T) {
	var stack stack.Stack

	values := []int{2, 3, 4}

	for _, val := range values {
		stack.Push(val)
	}

	lengthBefore := len(stack)
	returnVal := stack.Pop()
	lengthAfter := len(stack)

	if returnVal != 4 {
		log.Fatalf("Pop function did not return the first value")
	}

	if lengthBefore != lengthAfter+1 {
		log.Fatalf("Pop did not properly reduce the size of the queue after returning the first value")
	}
}
