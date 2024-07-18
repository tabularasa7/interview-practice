package main

import (
	"fmt"
	// "interview/practice/binarytree"
	"interview/practice/heap"
	// httpclient "interview/practice/http_client"
	// "interview/practice/linkedlist"
	// "interview/practice/queue"
	// "interview/practice/search/binary"
	breadthfirst "interview/practice/search/breadth_first"
	depthfirst "interview/practice/search/depth_first"
	// "interview/practice/stack"
	// "log"
)

func main() {
	// var queue queue.Queue
	// var stack stack.Stack
	// var linkedList linkedlist.LinkedList
	// var binarytree binarytree.Tree
	maxHeapInit := []int{5, 10, 3, 100, 28, 77}
	minHeapInit := []int{5, 10, 3, 100, 28, 77}
	maxHeap := heap.NewMaxHeap(maxHeapInit)
	minHeap := heap.NewMinHeap(minHeapInit)

	maxHeap.Insert(101)
	for i := maxHeap.Length() / 2; i >= 0; i-- {
		maxHeap.MaxHeapify(i)
	}

	minHeap.Insert(1)
	for i := minHeap.Length() / 2; i >= 0; i-- {
		minHeap.MinHeapify(i)
	}

	// queue.Push(4)
	// queue.Push(2)
	// stack.Push(2)
	// stack.Push(4)
	// linkedList.Insert(1)
	// linkedList.Insert(2)
	// linkedList.Insert(3)
	// linkedList.Insert(4)
	// binarytree.Insert(4)
	// binarytree.Insert(10)
	// binarytree.Insert(3)
	// binarytree.Insert(1)
	// binarytree.Insert(8)
	// binarytree.Insert(7)
	// binarytree.Insert(9)

	// // linkedList.PrintList()
	// binarytree.PrintTree()

	// fmt.Println(binarytree.Search(3))
	// fmt.Println(binarytree.Search(2))

	// binarytree.Delete(4)

	// binarytree.PrintTree()

	// linkedList.Delete(3)

	// fmt.Println(linkedList.Search(2))

	// linkedList.PrintList()

	// for len(queue) > 0 {
	// 	val := queue.Pop()
	// 	if val >= 0 {
	// 		fmt.Printf("queue: %d\n", val)
	// 	}
	// }

	// for len(stack) > 0 {
	// 	val := stack.Pop()
	// 	if val >= 0 {
	// 		fmt.Printf("stack: %d\n", val)
	// 	}
	// }

	// err := httpclient.MakeHTTPRequest("https://www.google.com/")

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	graph := breadthfirst.CreateGraph()
	// distance, err := breadthfirst.BFSearch(graph, 2, 4)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("Distance to given node is: %d\n", distance)

	depthfirst.DFSearch(graph, 5)

	// val := binary.Search([]int{1, 2, 3, 4, 5, 6, 7}, 6)

	// fmt.Printf("Found value using binary search: %d\n", val)

	fmt.Printf("Initial max heap: %v\n\n", maxHeap)

	fmt.Printf("Initial min heap: %v\n\n", minHeap)
}
