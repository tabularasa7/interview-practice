package binarytree_test

import (
	"interview/practice/binarytree"
	"log"
	"testing"
)

func TestInsert(t *testing.T) {
	var tree binarytree.Tree

	values := []int{4, 9, 0, 2, -3}
	for _, val := range values {
		tree.Insert(val)
	}

	if tree.IsEmpty() {
		log.Fatal("tree is empty")
	}
}

func TestSearch(t *testing.T) {
	var tree binarytree.Tree

	values := []int{4, 9, 0, 2, -3}

	for _, val := range values {
		tree.Insert(val)
	}

	searchVal := tree.Search(9)

	if searchVal == -1 {
		log.Fatalf("Did not find value in tree")
	}

}

func TestDelete(t *testing.T) {
	var tree binarytree.Tree

	values := []int{4, 9, 0, 2, -3}

	for _, val := range values {
		tree.Insert(val)
	}

	tree.Delete(-3)

	searchVal := tree.Search(-3)

	if searchVal != -1 {
		log.Fatalf("Did not properly delete node")
	}

}
