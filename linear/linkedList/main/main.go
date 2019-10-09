package main

import (
	"fmt"

	"github.com/xdragon1015/DataStructures/linear/linkedList"
)

func main() {
	fmt.Println("Hello World")
	list := &linkedList.LinkedList{}

	list.Add(100)
	list.Add(1000)
	list.Add(10000)

	list.PrintLinkedList()

	node := *list.LastNode()
	fmt.Println()

}
