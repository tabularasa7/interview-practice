package linkedlist

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

// inserts at tail
func (list *LinkedList) Insert(val int) {
	newNode := &Node{data: val}

	if list.head == nil {
		list.head = newNode
	} else {
		currentNode := list.head

		for currentNode.next != nil {
			currentNode = currentNode.next
		}

		currentNode.next = newNode
	}

}

func (list *LinkedList) Delete(val int) {
	deleteNode := list.head
	var prevNode *Node

	for deleteNode.data != val && deleteNode.next != nil {
		prevNode = deleteNode
		deleteNode = deleteNode.next
	}

	prevNode.next = deleteNode.next
}

func (list *LinkedList) Search(val int) bool {
	nodeExist := false
	currentNode := list.head

	for currentNode.data != val && currentNode.next != nil {
		currentNode = currentNode.next
	}

	if currentNode.data == val {
		nodeExist = true
	}

	return nodeExist
}

func (list *LinkedList) PrintList() {
	current := list.head

	if current == nil {
		fmt.Println("Linked list is empty")
	}
	for current != nil {
		fmt.Printf("list node: %d\n", current.data)
		current = current.next
	}
}

func (list *LinkedList) IsEmpty() bool {
	return list.head == nil
}
