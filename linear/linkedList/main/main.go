package main

import (
	"fmt"

	linkedlist "github.com/xdragon1015/DataStructures/linear/linkedList"
)

func main() {
	list := &linkedlist.LinkedList{}

	list.Add(100)
	list.Add(1000)
	list.Add(10000)

	list.PrintLinkedList()

	fmt.Println(*list.LastNode())

	list.AddToTheEnd(1000000)
	list.PrintLinkedList()

	// fmt.Println(list.Nod(1000000))

}
