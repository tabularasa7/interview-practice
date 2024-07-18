package linkedlist_test

import (
	"interview/practice/linkedlist"
	"log"
	"testing"
)

func TestInsert(t *testing.T) {
	var linkedList linkedlist.LinkedList

	nodes := []int{4, 3, 6, 8, 5, 5}

	for _, val := range nodes {
		linkedList.Insert(val)
	}

	if linkedList.IsEmpty() {
		log.Fatalf("List is empty, inserting failed")
	}
}

func TestSearch(t *testing.T) {
	var linkedList linkedlist.LinkedList

	nodes := []int{4, 3, 6, 8, 5, 5}

	for _, val := range nodes {
		linkedList.Insert(val)
	}

	found := linkedList.Search(8)

	if !found {
		log.Fatalf("Node was not successfully found in list")
	}
}

func TestDelete(t *testing.T) {
	var linkedList linkedlist.LinkedList

	nodes := []int{4, 3, 6, 8, 5, 5}

	for _, val := range nodes {
		linkedList.Insert(val)
	}

	linkedList.Delete(8)

	found := linkedList.Search(8)

	if found {
		log.Fatalf("Node was not successfully deleted from list")
	}
}
