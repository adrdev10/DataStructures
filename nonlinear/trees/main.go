package main

import (
	"fmt"

	"github.com/xdragon1015/DataStructures/nonlinear/trees/bst"
)

func main() {
	fmt.Println("Test cases")
	tree := bst.BinarySearchTree{
		RootNode: &bst.TreeNode{
			Value: 8,
		},
	}
	//          8
	//       2     10
	//     1  5   9  100
	tree.InsertElement(1, 2)
	tree.InsertElement(1, 1)
	tree.InsertElement(4, 5)
	tree.InsertElement(4, 10)
	tree.InsertElement(4, 100)
	tree.InsertElement(4, 9)
	tree.InOrderTraversal(print)
	fmt.Println("Post order traversal")
	tree.PostOrderTraversal(print)

	lw := tree.MinNode()
	fmt.Println("Lowest Value in BST is:", *lw)
}

func print(val int) {
	fmt.Println(val)
}
