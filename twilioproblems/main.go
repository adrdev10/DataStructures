package main

import (
	"fmt"
	"strings"
)

type Node struct {
	info     string
	next     *Node
	previous *Node
}

type LinkedList struct {
	firstNode *Node
	lastNode  *Node
}

func main() {
	var s string
	var t string
	s = "Adrian who is a student loves programming"
	t = "who student loves"
	findMissingWords(s, t)
}

//adrian -> who -> is -> a -> |student| -> loves -> programming

//Given two strings s, t; find the subsequent words in s that are missing in t
//They do not need to be contiguis
//Order matters

func (linkList *LinkedList) push(nodeInfo string) {
	node := &Node{info: nodeInfo}
	if linkList.firstNode != nil {
		node.next = linkList.firstNode
		linkList.firstNode.previous = node
		linkList.firstNode = node
	} else {
		linkList.firstNode = node
		linkList.firstNode.previous = node
	}
}

func (listNode *LinkedList) searchAndDelete(info string) {
	for node := listNode.firstNode; node != nil; node = node.next {
		if node.info == info {
			node.previous = node.next
			node.next = node.previous
			node.info = " "
		}
	}
}

func (listNode *LinkedList) printList() {
	for node := listNode.firstNode; node != nil; node = node.next {
		fmt.Println(node.info)
	}
}

func (listNode *LinkedList) collectNonEmpty() (arr []string) {
	for inTake := listNode.firstNode; inTake != nil; inTake = inTake.next {
		if inTake.next == nil {
			pNode := inTake.previous
			listNode.lastNode = inTake
			pNode.next = listNode.lastNode
			listNode.lastNode.previous = pNode
			break
		}
	}
	for lNode := listNode.lastNode; lNode != nil; lNode = lNode.previous {
		if lNode.info != " " {
			arr = append(arr, lNode.info)
		}
	}
	return
}

func findMissingWords(s, t string) {
	//put strings into arrays
	var sIntoArr []string
	var tIntoArr []string

	sIntoArr = strings.Split(s, " ")
	tIntoArr = strings.Split(t, " ")
	fmt.Println("Array for string s", sIntoArr)
	fmt.Println("Array for string t", tIntoArr)

	listString := &LinkedList{}
	for _, val := range sIntoArr {
		listString.push(val)
	}

	for _, val := range tIntoArr {
		listString.searchAndDelete(val)
	}

	listString.printList()
	arr := listString.collectNonEmpty()
	joinedString := strings.Join(arr, " ")
	fmt.Println(joinedString)
}
