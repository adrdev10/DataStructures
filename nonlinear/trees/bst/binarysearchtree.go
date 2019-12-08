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

func (tree *BinarySearchTree) SearchNode(value int) bool {
	tree.Lock.Lock()
	defer tree.Lock.Unlock()
	return searchNode(tree.RootNode, value)
}

func searchNode(node *TreeNode, value int) bool {
	if node == nil {
		return false
	}
	if value < node.Value {
		return searchNode(node.LeftNode, value)
	}
	if value > node.Value {
		return searchNode(node.RightNode, value)
	}
	return true
}

func (tree *BinarySearchTree) RemoveNode(value int) {
	tree.Lock.Lock()
	defer tree.Lock.Unlock()
	removeNode(tree.RootNode, value)
}

func removeNode(nodeTree *TreeNode, value int) *TreeNode {
	if nodeTree == nil {
		return nil
	}
	if value < nodeTree.Value {
		nodeTree.LeftNode = removeNode(nodeTree.LeftNode, value)
		return nodeTree
	}
	if value > nodeTree.Value {
		nodeTree.RightNode = removeNode(nodeTree.RightNode, value)
		return nodeTree
	}
	//nodeValue == value
	//check that the node has no children
	if nodeTree.LeftNode == nil && nodeTree.RightNode == nil {
		nodeTree = nil
		return nil
	}
	if nodeTree.LeftNode == nil {
		nodeTree = nodeTree.RightNode
		return nodeTree
	}
	if nodeTree.RightNode == nil {
		nodeTree = nodeTree.LeftNode
		return nodeTree
	}
	var leftMostRightTree *TreeNode
	leftMostRightTree = nodeTree.RightNode
	for {
		if leftMostRightTree != nil && leftMostRightTree.LeftNode != nil {
			leftMostRightTree = leftMostRightTree.LeftNode
		} else {
			break
		}
	}
	nodeTree.Value = leftMostRightTree.Value
	nodeTree.RightNode = removeNode(nodeTree.RightNode, nodeTree.Value)
	return nodeTree
}

func (tree *BinarySearchTree) MaxNode() *int {
	tree.Lock.Lock()
	defer tree.Lock.Unlock()
	var treeNode *TreeNode
	treeNode = tree.RootNode
	if treeNode == nil {
		return (*int)(nil)
	}
	for {
		if treeNode.RightNode == nil {
			return &treeNode.Value
		}
		treeNode = treeNode.RightNode
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
