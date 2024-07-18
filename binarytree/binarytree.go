package binarytree

import (
	"fmt"
)

type Node struct {
	data      int
	leftNode  *Node
	rightNode *Node
}

type Tree struct {
	root *Node
}

func (tree *Tree) Insert(val int) {

	if tree.root == nil {
		tree.root = &Node{data: val, leftNode: nil, rightNode: nil}
	} else {
		currentNode := tree.root

		currentNode.insertNode(&Node{data: val, leftNode: nil, rightNode: nil})
	}

}

func (node *Node) insertNode(newNode *Node) {
	if newNode.data < node.data {
		if node.leftNode == nil {
			node.leftNode = newNode
		} else {
			node.leftNode.insertNode(newNode)
		}
	} else {
		if node.rightNode == nil {
			node.rightNode = newNode
		} else {
			node.rightNode.insertNode(newNode)
		}
	}
}

func (tree *Tree) Delete(val int) {
	if tree.IsEmpty() {
		fmt.Println("Tree is empty.")
	} else {
		currentNode := tree.root

		currentNode.deleteNode(nil, val)
	}

}

func (node *Node) deleteNode(parentNode *Node, val int) {
	if node == nil {
		fmt.Println("Node does not exist")
	} else {
		switch {
		case val < node.data:
			node.leftNode.deleteNode(node, val)
		case val > node.data:
			node.rightNode.deleteNode(node, val)
		default:
			if node.leftNode == nil && node.rightNode == nil {
				fmt.Printf("deleting leaf node: %d\n", val)
				if parentNode.leftNode == node {
					parentNode.leftNode = nil
				} else if parentNode.rightNode == node {
					parentNode.rightNode = nil
				}
			} else if node.rightNode != nil && node.leftNode == nil {
				node.data = node.rightNode.data
				node.rightNode = nil
			} else if node.leftNode != nil && node.rightNode == nil {
				node.data = node.leftNode.data
				node.leftNode = nil
			} else {
				minNode := node.rightNode
				for minNode.leftNode != nil {
					minNode = minNode.leftNode
				}
				fmt.Printf("deleting node: %d\n", node.data)
				node.data = minNode.data
				node.rightNode.deleteNode(node, minNode.data)
			}
		}
	}
}

func (tree *Tree) Search(val int) int {

	if tree.IsEmpty() {
		return -1
	}

	root := tree.root

	foundVal := root.searchNode(val)

	return foundVal
}

func (node *Node) searchNode(val int) int {
	var returnVal int
	fmt.Printf("Searching for node...\n")
	if node.data == val {
		returnVal = node.data
	} else if val < node.data && node.leftNode != nil {
		returnVal = node.leftNode.searchNode(val)
	} else if val > node.data && node.rightNode != nil {
		returnVal = node.rightNode.searchNode(val)
	} else {
		return -1
	}

	return returnVal
}

func (tree *Tree) PrintTree() {
	root := tree.root

	root.preOrderPrint()
}

func (node *Node) inOrderPrint() {
	if node == nil {
		return
	}

	if node.leftNode != nil {
		node.leftNode.inOrderPrint()
	}
	fmt.Println(node.data)
	if node.rightNode != nil {
		node.rightNode.inOrderPrint()
	}
}

func (node *Node) preOrderPrint() {
	if node == nil {
		return
	}

	fmt.Println(node.data)
	if node.leftNode != nil {
		node.leftNode.preOrderPrint()
	}
	if node.rightNode != nil {
		node.rightNode.preOrderPrint()
	}
}

func (tree *Tree) IsEmpty() bool {
	return tree.root == nil
}
