//Package linkedList represents the linkedlist data structures
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
	if listNode.headNode != nil {
		node.nextNode = listNode.headNode
		listNode.headNode = node
	} else {
		listNode.headNode = node

	}
}

//LastNode returns the last node of the linkedlist. Deletion and insertion takes 0(1) constant time
func (listNode *LinkedList) LastNode() *Node {
	node := &Node{}
	outNode := &Node{}
	for node = listNode.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			outNode = node
			return outNode
		}
	}
	return nil
}

//AddToTheEnd adds a new node to the end of the list
func (listNode *LinkedList) AddToTheEnd(property int) {
	node := &Node{property: property}
	if n := listNode.LastNode(); n != nil {
		n.nextNode = node
	}
}

//NodeWithValue returns the node with property value
func (listNode LinkedList) NodeWithValue(property int) {

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
