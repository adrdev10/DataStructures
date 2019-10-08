package linkedList

import "fmt"

type Node struct {
	property int
	nextNode *Node
}

type LinkedList struct {
	headNode *Node
}

func (listNode *LinkedList) Add(property int) {
	var node = &Node{}
	node.property = property
	if listNode.headNode == nil {
		listNode.headNode = node
	} else if listNode.headNode != nil {
		node.nextNode = listNode.headNode
		listNode.headNode = node
	}
}

func (listNode *LinkedList) PrintLinkedList() {
	node := &Node{}
	if listNode.headNode == nil {
		fmt.Println("Nothing to print")
	}

	for node = listNode.headNode; node != nil; node = node.nextNode {
		fmt.Println("Node with value", node.property)
	}
}
