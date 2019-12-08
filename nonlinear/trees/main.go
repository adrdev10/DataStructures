package main

import (
	"fmt"

	"github.com/xdragon1015/DataStructures/nonlinear/trees/avstree"
)

type integerKey int

func (k integerKey) LessThan(k1 avstree.KeyValue) bool {
	return k < k1.(integerKey)
}

func (k integerKey) EqualTo(k1 avstree.KeyValue) bool {
	return k == k1.(integerKey)
}

func main() {
	// fmt.Println("Test cases")
	// tree := bst.BinarySearchTree{
	// 	RootNode: &bst.TreeNode{
	// 		Value: 8,
	// 	},
	// }
	// //          8
	// //       2     10
	// //     1  5   9  100
	// tree.InsertElement(1, 2)
	// tree.InsertElement(1, 1)
	// tree.InsertElement(4, 5)
	// tree.InsertElement(4, 10)
	// tree.InsertElement(4, 100)
	// tree.InsertElement(4, 9)
	// tree.InOrderTraversal(print)
	// fmt.Println("Post order traversal")
	// tree.PostOrderTraversal(print)

	// lw := tree.MinNode()
	// fmt.Println("Lowest Value in BST is:", *lw)

	// fmt.Println(*tree.MaxNode())

	// fmt.Println(tree.SearchNode(10))
	// tree.RemoveNode(10)
	// tree.InOrderTraversal(print)

	//            5
	// 			3   8
	//                10
	//
	var avst *avstree.TreeNode
	avstree.InsertNode(&avst, integerKey(5))
	avstree.InsertNode(&avst, integerKey(3))
	avstree.InsertNode(&avst, integerKey(8))
	avstree.InsertNode(&avst, integerKey(10))
	avstree.InsertNode(&avst, integerKey(20))
	fmt.Println(avst)
}

func print(val int) {
	fmt.Println(val)
}
