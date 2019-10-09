package linkedList

import "fmt"

//Node represents the node in the linkedlist data structure
type Node struct {
	property int
	nextNode *Node
}

//LinkedList represents the actual DT
type LinkedList struct {
	headNode *Node
}

//Add adds a new node to the linkedlist
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

// //Remove removes a node from a linkedlist
// func (listNode *LinkedList) Remove(property int) {

// }

//LastNode returns the last node of the linkedlist. Deletion and insertion takes 0(1) constant time
func (listNode *LinkedList) LastNode() *Node {
	node := &Node{}
	outNode := &Node{}
	for node = listNode.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			outNode = node.nextNode
		}
	}
	return outNode
}

//PrintLinkedList prints the content of the linkedList
func (listNode *LinkedList) PrintLinkedList() {
	node := &Node{}
	if listNode.headNode == nil {
		fmt.Println("Nothing to print")
	}

	for node = listNode.headNode; node != nil; node = node.nextNode {
		fmt.Println("Node with value", node.property)
	}
}
