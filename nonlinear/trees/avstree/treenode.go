package avstree

import "fmt"

type TreeNode struct {
	KeyValue     KeyValue
	BalanceValue int
	LinkedNodes  [2]*TreeNode
}

func opposite(nodeValue int) int {
	return 1 - nodeValue
}

func InsertNode(treenode **TreeNode, key KeyValue) {
	*treenode, _ = insertRNode(*treenode, key)
}

//single rotation function
func singleRotation(rootNode *TreeNode, nodeValue int) *TreeNode {
	var saveNode *TreeNode
	saveNode = rootNode.LinkedNodes[opposite(nodeValue)]
	rootNode.LinkedNodes[opposite(nodeValue)] = saveNode.LinkedNodes[nodeValue]
	saveNode.LinkedNodes[nodeValue] = rootNode
	return saveNode
}

func doubleRotation(rootNode *TreeNode, nodeValue int) *TreeNode {
	var saveNode *TreeNode
	saveNode = rootNode.LinkedNodes[opposite(nodeValue)].LinkedNodes[nodeValue]
	rootNode.LinkedNodes[opposite(nodeValue)].LinkedNodes[nodeValue] = saveNode.LinkedNodes[opposite(nodeValue)]
	saveNode.LinkedNodes[opposite(nodeValue)] = rootNode.LinkedNodes[opposite(nodeValue)]
	rootNode.LinkedNodes[opposite(nodeValue)] = saveNode
	saveNode = rootNode.LinkedNodes[opposite(nodeValue)]
	rootNode.LinkedNodes[opposite(nodeValue)] = saveNode.LinkedNodes[nodeValue]
	saveNode.LinkedNodes[nodeValue] = rootNode
	return saveNode
}

//adjust balance of the tree
func adjustBalance(rootNode *TreeNode, nodeValue int, balanceValue int) {
	var node *TreeNode
	node = rootNode.LinkedNodes[nodeValue]
	var oppNode *TreeNode
	oppNode = node.LinkedNodes[opposite(balanceValue)]
	switch oppNode.BalanceValue {
	case 0:
		rootNode.BalanceValue = 0
		node.BalanceValue = 0
	case balanceValue:
		rootNode.BalanceValue = rootNode.BalanceValue - balanceValue
		node.BalanceValue = 0
	default:
		rootNode.BalanceValue = 0
		node.BalanceValue = balanceValue
	}
	oppNode.BalanceValue = 0
}

//Changes balance factor by a single or double rotation
func BalanceTree(rootNode *TreeNode, nodeValue int) *TreeNode {
	var node *TreeNode
	node = rootNode.LinkedNodes[nodeValue]
	var balance int
	balance = nodeValue - 1
	if node.BalanceValue == balance {
		rootNode.BalanceValue = 0
		node.BalanceValue = 0
		return singleRotation(rootNode, opposite(nodeValue))
	}
	adjustBalance(rootNode, nodeValue, balance)
	return doubleRotation(rootNode, opposite(nodeValue))
}

func insertRNode(rootNode *TreeNode, key KeyValue) (*TreeNode, bool) {
	if rootNode == nil {
		return &TreeNode{KeyValue: key}, false
	}
	var dir int
	dir = 0
	if rootNode.KeyValue.LessThan(key) {
		dir = 1
	}
	var done bool
	rootNode.LinkedNodes[dir], done = insertRNode(rootNode.LinkedNodes[dir], key)
	if done {
		return rootNode, true
	}
	rootNode.BalanceValue = rootNode.BalanceValue + (dir - 1)
	switch rootNode.BalanceValue {
	case 0, 1, -1:
		return rootNode, true
	case -2, 2:
		return rootNode, false
	}
	fmt.Println(dir)
	return BalanceTree(rootNode, dir), true
}
