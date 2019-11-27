package bst

//pupose of preventing race conditions
import "sync"

type BinarySearchTree struct {
	RootNode *TreeNode
	Lock     sync.RWMutex
}

func (tree *BinarySearchTree) InsertElement(key int, value int) {
	tree.Lock.Lock()
	defer tree.Lock.Unlock()
	var treeNode *TreeNode
	treeNode = &TreeNode{key, value, nil, nil}
	if tree.RootNode == nil {
		tree.RootNode = treeNode
	} else {
		insertTreeNode(tree.RootNode, treeNode)
	}
}

// func (tree *BinarySearchTree) MaxNode() *int {

//}

func (tree *BinarySearchTree) MinNode() *int {
	tree.Lock.Lock()
	defer tree.Lock.Unlock()
	var treeNode *TreeNode
	treeNode = tree.RootNode
	if treeNode == nil {
		return (*int)(nil)
	}
	for {
		if treeNode.LeftNode == nil {
			return &treeNode.Value
		}
		treeNode = treeNode.LeftNode
	}
}

func (tree *BinarySearchTree) InOrderTraversal(function func(int)) {
	tree.Lock.Lock()
	defer tree.Lock.Unlock()
	inOrderTraversal(tree.RootNode, function)
}

func inOrderTraversal(tree *TreeNode, function func(int)) {
	if tree != nil {
		inOrderTraversal(tree.LeftNode, function)
		function(tree.Value)
		inOrderTraversal(tree.RightNode, function)
	}
}

func (tree *BinarySearchTree) PostOrderTraversal(function func(int)) {
	tree.Lock.Lock()
	defer tree.Lock.Unlock()
	postOrderTraversal(tree.RootNode, function)
}

func postOrderTraversal(tree *TreeNode, function func(int)) {
	if tree != nil {
		postOrderTraversal(tree.LeftNode, function)
		postOrderTraversal(tree.RightNode, function)
		function(tree.Value)
	}
}

func insertTreeNode(rootNode *TreeNode, newTreeNode *TreeNode) {
	if newTreeNode.Value < rootNode.Value {
		if rootNode.LeftNode == nil {
			rootNode.LeftNode = newTreeNode
		} else {
			insertTreeNode(rootNode.LeftNode, newTreeNode)
		}
	} else {
		if rootNode.RightNode == nil {
			rootNode.RightNode = newTreeNode
		} else {
			insertTreeNode(rootNode.RightNode, newTreeNode)
		}
	}
}
